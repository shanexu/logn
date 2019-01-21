package appender

import (
	"github.com/shanexu/logp/appender/encoder"
	"github.com/shanexu/logp/appender/writer"
	"github.com/shanexu/logp/common"
)

type Appender struct {
	Writer  writer.Writer
	Encoder encoder.Encoder
}

func CreateAppender(writerType string, config *common.Config) (*Appender, error) {
	w, err := writer.NewWriter(writerType, config)
	if err != nil {
		return nil, err
	}
	encoderConfig, err := config.Child("encoder", -1)
	if err != nil {
		return nil, err
	}
	ec := encoder.Config{}
	if err := encoderConfig.Unpack(&ec); err != nil {
		return nil, err
	}
	e, err := encoder.CreateEncoder(ec)
	if err != nil {
		return nil, err
	}
	return &Appender{w, e}, nil
}
