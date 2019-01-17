package configuration

type Logger struct {
	Name         string   `json:"name"`
	Level        string   `json:"level"`
	AppenderRefs []string `json:"appender_refs"`
}

type Loggers struct {
	Logger []Logger `json:"logger"`
	Root   Logger   `json:"root"`
}

type Appender map[string]interface{}

func (ap Appender) Name() string {
	return ap["name"].(string)
}

type Configuration struct {
	Appenders map[string][]Appender `json:"appenders"`
	Loggers   Loggers               `json:"loggers"`
}
