package log

import (
	"screw/log/appender"
	"screw/log/event"
	"screw/log/layout"
	"screw/log/level"
	"sync"
)

type Entry struct {
	name   string
	append appender.Appender
	layout layout.Layout
	lvl    level.Level

	lock sync.Mutex
}

func NewEntry(name string, append appender.Appender, layout layout.Layout) *Entry {
	entry := &Entry{
		name:   name,
		append: append,
		layout: layout,
		lvl:    level.DebugLevel(),
	}

	return entry
}

func (this *Entry) SetLevel(lvl level.Level) {
	this.lock.Lock()
	defer this.lock.Unlock()

	this.lvl = lvl
}

func (this *Entry) Name() string {
	return this.name
}

func (this *Entry) Output(event *event.Event) {
	this.lock.Lock()
	defer this.lock.Unlock()

	// 等级过滤
	if this.lvl.Greater(event.Level()) {
		return
	}

	this.append.Append(this.layout.Format(event))
}
