package appender

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// e.g ConsoleAppenderConf
//{
//	"Type": "ConsoleAppender",
//	"ConsoleType": "STDOUT"
//}

type ConsoleType int

const (
	CONSOLE_STDOUT ConsoleType = 0
	CONSOLE_STDERR             = 1
)

// ConsoleAppenderConf
type ConsoleAppenderConf struct {
	// 必须为: ConsoleAppender
	Type string

	// 取值: STDOUT/STDERR
	ConsoleType string
}

func (this *ConsoleAppenderConf) TypeName() string {
	return "ConsoleAppender"
}

func (this *ConsoleAppenderConf) Marshal_JSON() ([]byte, error) {
	return json.Marshal(this)
}

func (this *ConsoleAppenderConf) Unmarshal_JSON(data []byte) error {
	err := json.Unmarshal(data, this)
	if err != nil {
		return err
	}

	if this.Type != this.TypeName() {
		return fmt.Errorf("ConsoleAppenderConf Type must equal[%s] not [%s].", this.TypeName(), this.Type)
	}

	return nil
}

// ConsoleAppender
type ConsoleAppender struct {
	typ ConsoleType
	out *bufio.Writer

	lock sync.Mutex
}

func NewConsoleAppender(typ ConsoleType) *ConsoleAppender {
	appender := &ConsoleAppender{
		typ: typ,
	}

	switch typ {
	case CONSOLE_STDOUT:
		appender.out = bufio.NewWriter(os.Stdout)
	case CONSOLE_STDERR:
		appender.out = bufio.NewWriter(os.Stderr)
	default:
		appender.out = bufio.NewWriter(os.Stdout)
	}

	return appender
}

func (this *ConsoleAppender) Close() (err error) {
	return nil
}

func (this *ConsoleAppender) Append(message string) {
	this.lock.Lock()
	defer this.lock.Unlock()

	this.out.WriteString(message)
	this.out.Flush()
}

func (this *ConsoleAppender) LoadConf(conf AppenderConf) error {
	confE, ok := conf.(*ConsoleAppenderConf)
	if !ok {
		return fmt.Errorf("LoadConf error: conf can not convert type [%s].", "ConsoleAppenderConf")
	}

	this.lock.Lock()
	defer this.lock.Unlock()

	switch confE.ConsoleType {
	case "STDOUT":
		this.typ = CONSOLE_STDOUT
		this.out = bufio.NewWriter(os.Stdout)
	case "STDERR":
		this.typ = CONSOLE_STDERR
		this.out = bufio.NewWriter(os.Stderr)
	default:
		this.typ = CONSOLE_STDOUT
		this.out = bufio.NewWriter(os.Stdout)
	}

	return nil
}
