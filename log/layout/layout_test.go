package layout

import (
	"testing"

	"screw/log/event"
	"screw/log/level"
)

func Test_Factory(t *testing.T) {
	data := []byte(`
	{
		"Type": "SimpleLayout"
	}
	`)

	layout, err := globalFactory.ParseJSON(data)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}

	event := event.NewEvent(0, level.DebugLevel(), "hello world!!!")
	t.Log(layout.Format(event))
}

func Test_SimpleLayout(t *testing.T) {
	event := event.NewEvent(0, level.DebugLevel(), "hello world!!!")
	simlay := NewSimpleLayout()

	t.Log(simlay.Format(event))
}
