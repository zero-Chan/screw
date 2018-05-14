package log

import (
	"encoding/json"
	"io/ioutil"
	"screw/log/appender"
	"screw/log/layout"
)

// Conf
type Conf struct {
	Entrys []EntryConf
}

type EntryConf struct {
	Name     string
	Appender json.RawMessage
	Layout   json.RawMessage
}

// ConfParser
type ConfParser struct {
}

func NewConfParser() *ConfParser {
	parser := &ConfParser{}

	return parser
}

func (this *ConfParser) ParseJSON(data []byte) (logger *Logger, err error) {
	conf := Conf{}

	err = json.Unmarshal(data, &conf)
	if err != nil {
		return
	}

	logger = NewLogger()
	for _, entry := range conf.Entrys {
		appender, nerr := appender.ParseJSON([]byte(entry.Appender))
		if nerr != nil {
			return
		}

		layout, nerr := layout.ParseJSON([]byte(entry.Layout))
		if nerr != nil {
			return
		}

		logger.AddEntry(NewEntry(entry.Name, appender, layout))
	}

	return
}

func (this *ConfParser) LoadJSONFile(filepath string) (logger *Logger, err error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}

	return this.ParseJSON(data)
}
