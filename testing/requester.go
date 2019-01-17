package testing

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

type Requester struct {
	URL     string
	Handler gin.HandlerFunc
}

func (r *Requester) Post(data []byte, headers map[string]string) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.POST(r.URL, r.Handler)

	recorder := httptest.NewRecorder()

	request, _ := http.NewRequest("POST", r.URL, bytes.NewReader(data))

	for key, value := range headers {
		request.Header.Add(key, value)
	}

	router.ServeHTTP(recorder, request)

	return recorder
}

func (r *Requester) Get(headers map[string]string) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET(r.URL, r.Handler)

	recorder := httptest.NewRecorder()

	request, _ := http.NewRequest("GET", r.URL, nil)

	for key, value := range headers {
		request.Header.Add(key, value)
	}

	router.ServeHTTP(recorder, request)

	return recorder
}
