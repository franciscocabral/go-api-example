package middlewares

import (
	"github.com/franciscocabral/go-api-example/logging"
	"github.com/franciscocabral/go-api-example/settings"
)

var (
	logger logging.Logger
	config *settings.Application
)

func init() {
	config = settings.Get()
	logger = logging.GetLogger(config.Logging)
}
