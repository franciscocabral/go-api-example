package logging

type Logger interface {
	SetLevel(level LogLevel)
	Debug(message string, data ...AdditionalData)
	Info(message string, data ...AdditionalData)
	Warning(message string, data ...AdditionalData)
	Error(message string, data ...AdditionalData)
	Critical(message string, data ...AdditionalData)
}

type AdditionalData map[string]interface{}
