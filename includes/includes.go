package includes

import (
	_ "github.com/shanexu/logn/appender/writer/console"
	_ "github.com/shanexu/logn/appender/writer/file"
	_ "github.com/shanexu/logn/appender/writer/gelfudp"
	_ "github.com/shanexu/logn/appender/writer/rollingfile"

	_ "github.com/shanexu/logn/appender/encoder/console"
	_ "github.com/shanexu/logn/appender/encoder/gelf"
	_ "github.com/shanexu/logn/appender/encoder/json"

	_ "github.com/shanexu/logn/core/zap"
)
