package appender

import (
	"testing"
)

func Test_Factory(t *testing.T) {
	data := []byte(`
	{
		"Type": "ConsoleAppender",
		"ConsoleType": "STDOUT"
	}
	`)

	appender, err := globalFactory.ParseJSON(data)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}

	appender.Append("hello world!!!\n")
}
