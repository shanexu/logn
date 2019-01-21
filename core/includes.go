package core

import (
	_ "github.com/shanexu/logp/appender"
	_ "github.com/shanexu/logp/appender/console"
	_ "github.com/shanexu/logp/appender/file"
	_ "github.com/shanexu/logp/appender/rollingfile"

	_ "github.com/shanexu/logp/appender/encoder"
	_ "github.com/shanexu/logp/appender/encoder/json"
)
