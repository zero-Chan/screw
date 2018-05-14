package level

import "fmt"

type Priority int

type Level interface {
	String() string
	Priority() Priority

	Equal(Level) bool
	GreaterEqueal(Level) bool
	Greater(Level) bool
}

const (
	LEVEL_DEBUG Priority = iota
	LEVEL_INFO
	LEVEL_NOTICE
	LEVEL_WARN
	LEVEL_ERROR
	LEVEL_CRITICAL
	LEVEL_ALERT
	LEVEL_FATAL
)

type DefaultLevel struct {
	priority Priority
}

func (this *DefaultLevel) String() string {
	switch this.priority {
	case LEVEL_DEBUG:
		return "DEBUG"
	case LEVEL_INFO:
		return "INFO"
	case LEVEL_NOTICE:
		return "NOTICE"
	case LEVEL_WARN:
		return "WARN"
	case LEVEL_ERROR:
		return "ERROR"
	case LEVEL_CRITICAL:
		return "CRITICAL"
	case LEVEL_ALERT:
		return "ALERT"
	case LEVEL_FATAL:
		return "FATAL"
	}

	panic("No such priority string")
}

func (this *DefaultLevel) Priority() Priority {
	return this.priority
}

func (this *DefaultLevel) Equal(lvl Level) bool {
	if this.priority == lvl.Priority() {
		return true
	}
	return false
}

func (this *DefaultLevel) GreaterEqueal(lvl Level) bool {
	if this.priority >= lvl.Priority() {
		return true
	}
	return false
}

func (this *DefaultLevel) Greater(lvl Level) bool {
	if this.priority > lvl.Priority() {
		return true
	}
	return false
}

func NewLevel(priority Priority) Level {
	return &DefaultLevel{
		priority: priority,
	}
}

func DebugLevel() Level {
	return NewLevel(LEVEL_DEBUG)
}

func InfoLevel() Level {
	return NewLevel(LEVEL_INFO)
}

func NoticeLevel() Level {
	return NewLevel(LEVEL_NOTICE)
}

func WarnLevel() Level {
	return NewLevel(LEVEL_WARN)
}

func ErrorLevel() Level {
	return NewLevel(LEVEL_ERROR)
}

func CriticalLevel() Level {
	return NewLevel(LEVEL_CRITICAL)
}

func AlertLevel() Level {
	return NewLevel(LEVEL_ALERT)
}

func FatalLevel() Level {
	return NewLevel(LEVEL_FATAL)
}

func ParseString(str string) (lvl Level, err error) {
	switch str {
	case "DEBUG":
		return DebugLevel(), nil
	case "INFO":
		return InfoLevel(), nil
	case "NOTICE":
		return NoticeLevel(), nil
	case "WARN":
		return WarnLevel(), nil
	case "ERROR":
		return ErrorLevel(), nil
	case "CRITICAL":
		return CriticalLevel(), nil
	case "ALERT":
		return AlertLevel(), nil
	case "FATAL":
		return FatalLevel(), nil
	}
	return nil, fmt.Errorf("No such level string")
}
