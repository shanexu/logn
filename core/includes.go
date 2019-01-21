package core

import (
	_ "github.com/shanexu/logn/appender"

	_ "github.com/shanexu/logn/appender/writer/console"
	_ "github.com/shanexu/logn/appender/writer/file"
	_ "github.com/shanexu/logn/appender/writer/rollingfile"

	_ "github.com/shanexu/logn/appender/encoder"
	_ "github.com/shanexu/logn/appender/encoder/json"
)
