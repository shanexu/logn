package gelfudp

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/shanexu/logn/appender/writer"
	"github.com/shanexu/logn/common"
	"go.uber.org/zap/zapcore"
	"io"
	"io/ioutil"
	"net"
)

type Config struct {
	Host             string `logn-config:"host"`
	Port             int    `logn-config:"port"`
	CompressionType  string `logn-config:"compression_type" logn-validate:"logn.oneof=none gzip zlib"`
	CompressionLevel int    `logn-config:"compression_level"`
}

var defaultConfig = Config{
	Host:             "127.0.0.1",
	Port:             12201,
	CompressionType:  "gzip",
	CompressionLevel: gzip.DefaultCompression,
}

const (
	MaxDatagramSize = 1420
	HeadSize        = 12
	MaxChunkSize    = MaxDatagramSize - HeadSize
	MaxChunks       = 128
	MaxMessageSize  = MaxChunkSize * MaxChunks

	CompressionNone = 0
	CompressionGzip = 1
	CompressionZlib = 2
)

var Magic = []byte{0x1e, 0x0f}
var ErrTooLargeMessageSize = errors.New("too large message size")

type UDPSender struct {
	raddr *net.UDPAddr
	conn  *net.UDPConn
	id    IdGenerator
}

func NewUDPSender(address string) (*UDPSender, error) {
	ip, err := GuessIP()
	if err != nil {
		return nil, err
	}
	id := NewDefaultIdGenerator(ip)
	raddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenUDP("udp", nil)
	if err != nil {
		return nil, err
	}
	return &UDPSender{
		raddr: raddr,
		conn:  conn,
		id:    id,
	}, nil
}

func (s *UDPSender) Send(message []byte) error {
	if len(message) > MaxMessageSize {
		return ErrTooLargeMessageSize
	}

	if len(message) <= MaxDatagramSize {
		_, err := s.conn.WriteToUDP(message, s.raddr)
		return err
	}

	chunks := len(message) / MaxChunkSize
	if chunks*MaxChunkSize < len(message) {
		chunks = chunks + 1
	}

	messageID := s.id.NextId()
	chunk := make([]byte, MaxDatagramSize)
	for i := 0; i < chunks; i++ {
		copy(chunk[0:2], Magic)
		binary.BigEndian.PutUint64(chunk[2:10], messageID)
		chunk[10] = byte(i)
		chunk[11] = byte(chunks)
		begin, end := i*MaxChunkSize, (i+1)*MaxChunkSize
		if end > len(message) {
			end = len(message)
		}
		copy(chunk[12:12+end-begin], message[begin:end])
		_, err := s.conn.WriteToUDP(chunk[0:12+end-begin], s.raddr)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewCompressor(compressionType string, compressionLevel int) (*Compressor, error) {
	switch compressionType {
	case "none":
		return &Compressor{CompressionNone, compressionLevel}, nil
	case "gzip":
		if _, err := gzip.NewWriterLevel(ioutil.Discard, compressionLevel); err != nil {
			return nil, err
		}
		return &Compressor{CompressionGzip, compressionLevel}, nil
	case "zlib":
		if _, err := zlib.NewWriterLevel(ioutil.Discard, compressionLevel); err != nil {
			return nil, err
		}
		return &Compressor{CompressionZlib, compressionLevel}, nil
	default:
		return nil, fmt.Errorf("no compression type %q", compressionType)
	}
}

type Compressor struct {
	compressionType  int
	compressionLevel int
}

func (c *Compressor) Compress(buf []byte) (int, []byte, error) {
	var (
		cw   io.WriteCloser
		cBuf bytes.Buffer
		err  error
	)
	switch c.compressionType {
	case CompressionNone:
		return len(buf), buf, nil
	case CompressionGzip:
		cw, err = gzip.NewWriterLevel(&cBuf, c.compressionLevel)
	case CompressionZlib:
		cw, err = zlib.NewWriterLevel(&cBuf, c.compressionLevel)
	}

	if err != nil {
		return 0, nil, err
	}

	n, err := cw.Write(buf)
	if err != nil {
		return 0, nil, err
	}

	if err := cw.Close(); err != nil {
		return 0, nil, err
	}

	return n, cBuf.Bytes(), nil
}

type Writer struct {
	sender     *UDPSender
	compressor *Compressor
}

func (w *Writer) Write(p []byte) (n int, err error) {
	n, b, err := w.compressor.Compress(p)
	if err != nil {
		return 0, err
	}
	if err := w.sender.Send(b); err != nil {
		return 0, err
	}
	return n, nil
}

func init() {
	writer.RegisterType("gelf_udp", func(rawConfig *common.Config) (writer.Writer, error) {
		config := defaultConfig
		if err := rawConfig.Unpack(&config); err != nil {
			return nil, err
		}
		c, err := NewCompressor(config.CompressionType, config.CompressionLevel)
		if err != nil {
			return nil, err
		}
		s, err := NewUDPSender(fmt.Sprintf("%s:%d", config.Host, config.Port))
		if err != nil {
			return nil, err
		}
		return zapcore.AddSync(&Writer{s, c}), nil
	})
}
