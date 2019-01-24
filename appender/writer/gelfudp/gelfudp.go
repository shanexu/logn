package gelfudp

import (
	"encoding/binary"
	"errors"
	"net"
)

type Config struct {
}

const (
	MAX_DATAGRAM_SIZE = 16
	HEAD_SIZE         = 12
	MAX_CHUNK_SIZE    = MAX_DATAGRAM_SIZE - HEAD_SIZE
	MAX_CHUNKS        = 128
	MAX_MESSAGE_SIZE  = MAX_CHUNK_SIZE * 128
)

var MAGIC = []byte{0x1e, 0x0f}
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
	if len(message) > MAX_MESSAGE_SIZE {
		return ErrTooLargeMessageSize
	}

	if len(message) <= MAX_DATAGRAM_SIZE {
		_, err := s.conn.WriteToUDP(message, s.raddr)
		return err
	}

	chunks := len(message) / MAX_CHUNK_SIZE
	if chunks*MAX_CHUNK_SIZE < len(message) {
		chunks = chunks + 1
	}

	messageID := s.id.NextId()
	chunk := make([]byte, MAX_DATAGRAM_SIZE)
	for i := 0; i < chunks; i++ {
		copy(chunk[0:2], MAGIC)
		binary.BigEndian.PutUint64(chunk[2:10], messageID)
		chunk[10] = byte(i)
		chunk[11] = byte(chunks)
		begin, end := i*MAX_CHUNK_SIZE, (i+1)*MAX_CHUNK_SIZE
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
