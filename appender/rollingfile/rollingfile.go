package rollingfile

import (
	"github.com/shanexu/logp/appender"
	"github.com/shanexu/logp/common"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type RollingFile struct {
	zapcore.WriteSyncer
}

type Config struct {
	// FileName is the file to write logs to.  Backup log files will be retained
	// in the same directory.  It uses <processname>-lumberjack.log in
	// os.TempDir() if empty.
	FileName string `config:"file_name" validate:"required"`

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int `config:"max_size" validate:"omitempty,gte=1"`

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `config:"max_age"`

	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups int `config:"max_backups"`

	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	LocalTime bool `config:"local_time"`

	// Compress determines if the rotated log files should be compressed
	// using gzip. The default is not to perform compression.
	Compress bool `config:"compress"`
}

func NewRollingFile(v *common.Config) (appender.Writer, error) {
	cfg := Config{}
	if err := v.Unpack(&cfg); err != nil {
		return nil, err
	}
	if err := common.Validate().Struct(cfg); err != nil {
		return nil, err
	}
	if cfg.MaxSize == 0 {
		cfg.MaxSize = 500
	}
	if cfg.MaxAge == 0 {
		cfg.MaxAge = 7
	}
	w := &lumberjack.Logger{
		Filename:   cfg.FileName,
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
		LocalTime:  cfg.LocalTime,
		Compress:   cfg.Compress,
	}
	return &RollingFile{zapcore.AddSync(w)}, nil
}

func init() {
	appender.RegisterType("rolling_file", NewRollingFile)
}
