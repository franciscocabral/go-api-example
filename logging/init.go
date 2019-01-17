package logging

import (
	"github.com/franciscocabral/go-api-example/settings"
	"strings"
)

var myLogger Logger

func GetLogger(logging *settings.Logging) Logger {
	myLevel := LevelDict[strings.ToLower(logging.Level)]

	if myLogger == nil {
		myLogger = new(ConsoleLogger)
		myLogger.SetLevel(myLevel)
	}

	return myLogger
}
