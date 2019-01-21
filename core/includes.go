package core

import (
	_ "github.com/shanexu/logn/appender"
	_ "github.com/shanexu/logn/appender/console"
	_ "github.com/shanexu/logn/appender/file"
	_ "github.com/shanexu/logn/appender/rollingfile"

	_ "github.com/shanexu/logn/appender/encoder"
	_ "github.com/shanexu/logn/appender/encoder/json"
)
