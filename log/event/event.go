package event

import (
	"runtime"
	"screw/log/level"
	"time"
)

// Event 记录每条log产生的数据
type Event struct {
	message   string
	timeStamp time.Time
	level     level.Level
	caller    *Caller
}

func NewEvent(calldepth int, lvl level.Level, message string) *Event {
	var (
		caller *Caller
	)

	_, file, line, ok := runtime.Caller(calldepth + 1)
	if ok {
		caller = NewCaller(file, line)
	} else {
		caller = NewCaller("", 0)
	}

	event := &Event{
		message:   message,
		timeStamp: time.Now(),
		level:     lvl,
		caller:    caller,
	}

	return event
}

func (this *Event) Level() level.Level {
	return this.level
}

func (this *Event) TimeStamp() time.Time {
	return this.timeStamp
}

func (this *Event) Caller() *Caller {
	return this.caller
}

func (this *Event) Message() string {
	return this.message
}
