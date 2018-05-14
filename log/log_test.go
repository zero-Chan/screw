package log

import (
	"screw/log/appender"
	"screw/log/layout"
	"testing"
)

func Test_ConfParser(t *testing.T) {
	data := []byte(`
	{
		"Entrys": [
			{
				"Name": "stdout_simple",
				"Appender": {
					"Type": "ConsoleAppender",
					"ConsoleType": "STDOUT"
				},
				"Layout": {
					"Type": "SimpleLayout"
				}
			}
		]
		
	}
	`)

	parser := NewConfParser()
	logger, err := parser.ParseJSON(data)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}

	logger.Debugf("conf's debugf log %s", "hello world!")
}

func Test_Log(t *testing.T) {
	log := NewLogger()

	entry := NewEntry(
		"log1",
		appender.NewConsoleAppender(appender.CONSOLE_STDOUT),
		layout.NewSimpleLayout(),
	)

	log.AddEntry(entry)

	log.Debugf("debugf log: %s", "hello world!")
	log.Debug("debug log: ", "hello world!")
}
