# logn

A wrap of [zap](https://github.com/uber-go/zap), maybe a simplified [log4j](https://logging.apache.org/log4j/2.x/)
for golang.

## Installation

`go get -u github.com/shanexu/logn`

## Quick Start

Logn loads configuration file from system environment virable `LOGN_CONFIG`. If the
variable is unset, then Logn will try to load the configuration file from current 
work directory, the file name is "logn.yaml" or "logn.yml".

```yam
appenders:
  console:
    - name: CONSOLE
      target: stdout
      encoder:
        console:
  file:
    - name: FILE
      file_name: /tmp/app.log
      encoder:
        json:
    - name: GELF_FILE
      file_name: /tmp/app_gelf.log
      encoder:
        gelf:
          key_value_pairs:
            - key: env
              value: ${ENV:dev}
            - key: app
              value: ${APPNAME:demo}
            - key: file
              value: app.log
    - name: METRICS
      file_name: /tmp/metrics.log
      encoder:
        json:
loggers:
  root:
    level: info
    appender_refs:
      - CONSOLE
  logger:
    - name: helloworld
      appender_refs:
        - CONSOLE
        - FILE
        - GELF_FILE
      level: debug

```

sample code:

```go
package main

import "github.com/shanexu/logn"

func main() {
	helloWorld := logn.GetLogger("helloworld")
	helloWorld.Info("hello, shane!")

	some := logn.GetLogger("some")
	some.Infof("hello, %s", "shane")
}
```
