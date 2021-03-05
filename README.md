# logn

A wrap of [zap](https://github.com/uber-go/zap), maybe a simplified [log4j](https://logging.apache.org/log4j/2.x/)
for golang.

## Installation

`go get -u github.com/shanexu/logn`

## Quick Start

Logn loads configuration file from system environment virable `LOGN_CONFIG`. If the
variable is unset, then Logn will try to load the configuration file from current 
work directory, the file name is "logn.yaml" or "logn.yml".

```yaml
appenders:
  console:
    - name: CONSOLE
      target: stdout
      encoder:
        console:
          time_encoder: ISO8601
  file:
    - name: FILE
      file_name: /tmp/app.log
      encoder:
        json:
  gelf_udp:
    - name: GRAYLOG
      host: 127.0.0.1
      port: 12201
      compression_type: none
      encoder:
        gelf:
          key_value_pairs:
            - key: env
              value: ${ENV:dev}
            - key: app
              value: ${APPNAME:demo}
            - key: file
              value: app.log
  rolling_file:
    - name: GELF_FILE
      file_name: /tmp/app_gelf.log
      max_size: 100
      encoder:
        gelf:
          key_value_pairs:
            - key: env
              value: ${ENV:dev}
            - key: app
              value: ${APPNAME:demo}
            - key: file
              value: app.log
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
        - GRAYLOG
      level: debug      
```

sample code:

```go
package main

import "github.com/shanexu/logn"

func main() {
    defer logn.Sync()
	helloWorld := logn.GetLogger("helloworld")
	helloWorld.Info("hello, shane!")

	some := logn.GetLogger("some")
	some.Infof("hello, %s", "shane")
}
```
