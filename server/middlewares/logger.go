package middlewares

import (
	"github.com/franciscocabral/go-api-example/logging"
	"github.com/gin-gonic/gin"
)

//Logger will log all requests
func Logger(context *gin.Context) {
	itc := newInterceptor(context)

	var err string
	if e := itc.StartTrack(); e != nil {
		err = e.Error()
	}
	request := itc.Request()

	logger.Info("Request started", logging.AdditionalData(map[string]interface{}{
		"url":     context.Request.URL,
		"error":   err,
		"request": request,
	}))

	context.Next()

	logger.Info("Request finished", logging.AdditionalData(map[string]interface{}{
		"url":          context.Request.URL,
		"error":        err,
		"status":       itc.Status(),
		"durationTime": itc.Duration(),
		"request":      request,
		"response":     itc.Response(),
	}))
}
