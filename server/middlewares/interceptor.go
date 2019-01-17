package middlewares

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

type interceptor struct {
	gin.ResponseWriter
	startTime   time.Time
	requestBuf  *bytes.Buffer
	responseBuf *bytes.Buffer
	context     *gin.Context
}

func newInterceptor(context *gin.Context) *interceptor {
	itc := &interceptor{
		ResponseWriter: context.Writer,
		context:        context,
	}
	return itc
}

func (itc *interceptor) Write(b []byte) (int, error) {
	if n, err := itc.responseBuf.Write(b); err != nil {
		return n, err
	}
	return itc.ResponseWriter.Write(b)
}

func (itc *interceptor) StartTrack() error {
	itc.startTime = time.Now()
	itc.requestBuf = bytes.NewBufferString("")
	itc.responseBuf = bytes.NewBufferString("")

	rb, err := ioutil.ReadAll(itc.context.Request.Body)
	if err != nil {
		return err
	}

	reader := ioutil.NopCloser(bytes.NewBuffer(rb))
	itc.context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rb))

	if _, err = itc.requestBuf.ReadFrom(reader); err != nil {
		return err
	}

	itc.context.Writer = itc
	return nil
}

func (itc *interceptor) Request() interface{} {
	var data interface{}
	if err := json.Unmarshal(itc.requestBuf.Bytes(), &data); err != nil {
		return itc.requestBuf.String()
	}

	return data
}

func (itc *interceptor) Response() interface{} {
	var data interface{}
	if err := json.Unmarshal(itc.responseBuf.Bytes(), &data); err != nil {
		return nil
	}

	return data
}

func (itc *interceptor) Duration() int64 {
	return time.Since(itc.startTime).Nanoseconds()
}

func (itc *interceptor) Status() int {
	return itc.ResponseWriter.Status()
}
