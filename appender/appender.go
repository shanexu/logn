package appender

import "io"

type Appender interface {
	io.Writer
	Sync() error
}
