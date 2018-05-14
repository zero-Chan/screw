package event

// Caller 记录运行时产生的数据
type Caller struct {
	file string
	line int
}

func NewCaller(file string, line int) *Caller {
	c := &Caller{
		file: file,
		line: line,
	}
	return c
}

func (this *Caller) File() string {
	return this.file
}

func (this *Caller) Line() int {
	return this.line
}
