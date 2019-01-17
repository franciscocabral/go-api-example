package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/franciscocabral/go-api-example/logging"
)

type RequestContext interface {
	NegotiateFormat(...string) string
	ShouldBindJSON(interface{}) error
	JSON(int, interface{})
	Header(string, string)
	Data(int, string, []byte)
}

func invokeMethod(c *gin.Context, h Handler) {
	switch c.Request.Method {
	case http.MethodGet:
		h.Get()
	case http.MethodPost:
		h.Post()
	case http.MethodPut:
		h.Put()
	case http.MethodDelete:
		h.Delete()
	default:
		h.Default()
	}
}

type baseHandler struct {
	context RequestContext
	Logger  logging.Logger
}

func (b *baseHandler) Get() {
	b.notAllowed()
}

func (b *baseHandler) Post() {
	b.notAllowed()
}

func (b *baseHandler) Put() {
	b.notAllowed()
}

func (b *baseHandler) Delete() {
	b.notAllowed()
}

func (b *baseHandler) Default() {
	b.notAllowed()
}

func (b *baseHandler) notAllowed() {
	b.Logger.Warning("Requested by a not allowed method")
	b.context.JSON(http.StatusMethodNotAllowed, ResponseError{"Method not allowed!"})
}
