package log

import (
	"fmt"

	"screw/log/event"
	"screw/log/level"
	"sync"
)

type Logger struct {
	entrys map[string]*Entry
	lock   sync.Mutex
}

func NewLogger() *Logger {
	logger := &Logger{
		entrys: make(map[string]*Entry),
	}

	return logger
}

// AddEntry 追加实例
func (this *Logger) AddEntry(entry *Entry) {
	if entry == nil {
		return
	}

	this.lock.Lock()
	this.entrys[entry.name] = entry
	this.lock.Unlock()
}

func (this *Logger) ListEntry() []*Entry {
	this.lock.Lock()
	defer this.lock.Unlock()

	es := make([]*Entry, 0, len(this.entrys))
	for _, entry := range this.entrys {
		es = append(es, entry)
	}

	return es
}

func (this *Logger) Debugf(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	eve := event.NewEvent(1, level.NewLevel(level.LEVEL_DEBUG), message)

	entrys := this.ListEntry()
	for _, entry := range entrys {
		entry.Output(eve)
	}
}

func (this *Logger) Debug(a ...interface{}) {
	message := fmt.Sprint(a...)
	eve := event.NewEvent(1, level.NewLevel(level.LEVEL_DEBUG), message)

	entrys := this.ListEntry()
	for _, entry := range entrys {
		entry.Output(eve)
	}
}
