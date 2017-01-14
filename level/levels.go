package level

type Level uint8

const (
	DEBUG Level = iota
	TRACE
	INFO
	WARN
	ERROR
	FATAL
)

var labels = []string{"DEBUG", "TRACE", "INFO", "WARN", "ERROR", "FATAL"}

func (level Level) Lable() string {
	return labels[level]
}
