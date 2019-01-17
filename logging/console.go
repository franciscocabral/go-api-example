package logging

import "fmt"

type ConsoleLogger struct {
	level LogLevel
}

func (c *ConsoleLogger) Debug(message string, data ...AdditionalData) {
	c.log(message, DebugLevel, data...)
}

func (c *ConsoleLogger) Info(message string, data ...AdditionalData) {
	c.log(message, InfoLevel, data...)
}

func (c *ConsoleLogger) Warning(message string, data ...AdditionalData) {
	c.log(message, WarningLevel, data...)
}

func (c *ConsoleLogger) Error(message string, data ...AdditionalData) {
	c.log(message, ErrorLevel, data...)
}

func (c *ConsoleLogger) Critical(message string, data ...AdditionalData) {
	c.log(message, CriticalLevel, data...)
}

func (c *ConsoleLogger) log(message string, logLevel LogLevel, data ...AdditionalData) {
	if logLevel >= c.level {
		println("------------------------------------------------------------")
		fmt.Printf("Log Level: %d\n", c.level)
		fmt.Printf("Message: %s\n", message)
		fmt.Printf("Data: %#v\n", data)
	}
}

func (c *ConsoleLogger) SetLevel(l LogLevel) {
	c.level = l
}
