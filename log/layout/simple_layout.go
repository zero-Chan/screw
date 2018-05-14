package layout

import (
	"encoding/json"
	"fmt"
	"screw/log/event"
)

// e.g SimpleLayoutConf
//{
//	"Type": "SimpleLayout"
//}

// SimpleLayoutConf
type SimpleLayoutConf struct {
	Type string
}

func (this *SimpleLayoutConf) TypeName() string {
	return "SimpleLayout"
}

func (this *SimpleLayoutConf) Marshal_JSON() (data []byte, error error) {
	return json.Marshal(this)
}

func (this *SimpleLayoutConf) Unmarshal_JSON(data []byte) (err error) {
	err = json.Unmarshal(data, this)
	if err != nil {
		return
	}

	if this.Type != this.TypeName() {
		return fmt.Errorf("SimpleLayoutConf Type must equal[%s] not [%s].", this.TypeName(), this.Type)
	}

	return
}

// SimpleLayout
type SimpleLayout struct {
}

func NewSimpleLayout() *SimpleLayout {
	layout := &SimpleLayout{}
	return layout
}

func (this *SimpleLayout) Format(event *event.Event) string {
	// [time] [level] file:line:message\n
	message := fmt.Sprintf(
		"[%s] [%s] %s:%d:%s\n",
		event.TimeStamp().Format("2006-01-02 15:04:05.000"),
		event.Level().String(),
		event.Caller().File(),
		event.Caller().Line(),
		event.Message(),
	)

	return message
}

func (this *SimpleLayout) LoadConf(LayoutConf) error {
	return nil
}
