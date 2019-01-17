package logging

type LogLevel int

const (
	DebugLevel    LogLevel = 1
	InfoLevel     LogLevel = 2
	WarningLevel  LogLevel = 3
	ErrorLevel    LogLevel = 4
	CriticalLevel LogLevel = 5
)

var LevelDict = map[string]LogLevel{
	"debug":    DebugLevel,
	"info":     InfoLevel,
	"warning":  WarningLevel,
	"error":    ErrorLevel,
	"critical": CriticalLevel,
}
