package gelf

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"testing"
	"time"
)

type dummyEncoder struct {
}

func (*dummyEncoder) AddArray(key string, marshaler zapcore.ArrayMarshaler) error {
	panic("implement me")
}

func (*dummyEncoder) AddBinary(key string, value []byte) {
	panic("implement me")
}

func (*dummyEncoder) AddBool(key string, value bool) {
	panic("implement me")
}

func (*dummyEncoder) AddByteString(key string, value []byte) {
	panic("implement me")
}

func (*dummyEncoder) AddComplex128(key string, value complex128) {
	panic("implement me")
}

func (*dummyEncoder) AddComplex64(key string, value complex64) {
	panic("implement me")
}

func (*dummyEncoder) AddDuration(key string, value time.Duration) {
	panic("implement me")
}

func (*dummyEncoder) AddFloat32(key string, value float32) {
	panic("implement me")
}

func (*dummyEncoder) AddFloat64(key string, value float64) {
	panic("implement me")
}

func (*dummyEncoder) AddInt(key string, value int) {
	panic("implement me")
}

func (*dummyEncoder) AddInt16(key string, value int16) {
	panic("implement me")
}

func (*dummyEncoder) AddInt32(key string, value int32) {
	panic("implement me")
}

func (*dummyEncoder) AddInt64(key string, value int64) {
	panic("implement me")
}

func (*dummyEncoder) AddInt8(key string, value int8) {
	panic("implement me")
}

func (*dummyEncoder) AddObject(key string, marshaler zapcore.ObjectMarshaler) error {
	panic("implement me")
}

func (*dummyEncoder) AddReflected(key string, value interface{}) error {
	panic("implement me")
}

func (*dummyEncoder) AddString(key, value string) {
	panic("implement me")
}

func (*dummyEncoder) AddTime(key string, value time.Time) {
	panic("implement me")
}

func (*dummyEncoder) AddUint(key string, value uint) {
	panic("implement me")
}

func (*dummyEncoder) AddUint16(key string, value uint16) {
	panic("implement me")
}

func (*dummyEncoder) AddUint32(key string, value uint32) {
	panic("implement me")
}

func (*dummyEncoder) AddUint64(key string, value uint64) {
	panic("implement me")
}

func (*dummyEncoder) AddUint8(key string, value uint8) {
	panic("implement me")
}

func (*dummyEncoder) AddUintptr(key string, value uintptr) {
	panic("implement me")
}

func (*dummyEncoder) Clone() zapcore.Encoder {
	panic("implement me")
}

func (*dummyEncoder) EncodeEntry(_ zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	b := &buffer.Buffer{}
	b.Write([]byte(fmt.Sprintf("%+v", fields)))
	return b, nil
}

func (*dummyEncoder) OpenNamespace(key string) {
	panic("implement me")
}

func TestEncoder_EncodeEntry(t *testing.T) {
	e := Encoder{
		Fields:  []zapcore.Field{zap.String("version", "1.1")},
		Encoder: &dummyEncoder{},
	}
	b, _ := e.EncodeEntry(zapcore.Entry{}, []zapcore.Field{zap.String("seq_id", "123")})
	assert.Equal(t, `[{Key:version Type:15 Integer:0 String:1.1 Interface:<nil>} {Key:_seq_id Type:15 Integer:0 String:123 Interface:<nil>}]`, b.String())
}
