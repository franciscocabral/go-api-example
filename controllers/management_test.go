package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/franciscocabral/go-api-example/logging/mocks"
	"github.com/franciscocabral/go-api-example/settings"

	. "github.com/franciscocabral/go-api-example/testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

//CancellationLetterHandler test methos from Management Controller
var _ = Describe("ManagementController", func() {
	Describe("[GET] When checking the api health", func() {

		var (
			req      = Requester{URL: "/management/health-check", Handler: checkHealth}
			recorder *httptest.ResponseRecorder
		)

		var (
			mockLogger *mocks.MockLogger
			mockConfig *settings.Application
		)

		BeforeEach(func() {
			mockLogger = new(mocks.MockLogger)
			mockConfig = new(settings.Application)

			mockConfig.Name = "MockApplication"
			mockConfig.CompanyName = "MockCompanyName"
			mockConfig.HostName = "MockHostName"
			mockConfig.Version = "MockVersion"

			logger = mockLogger
			config = mockConfig
		})

		Context("with an empty request data", func() {
			BeforeEach(func() {
				mockLogger.On("Info", mock.Anything, mock.Anything).Once()
				recorder = req.Get(nil)
			})
			It("should return api info with status 10", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))

				buf := new(bytes.Buffer)
				buf.ReadFrom(recorder.Body)

				var result healthResponse
				err := json.Unmarshal(buf.Bytes(), &result)

				Expect(err).To(BeNil())
				Expect(result.Application).To(Equal(config.Name))
				Expect(result.Company).To(Equal(config.CompanyName))
				Expect(result.Hostname).To(Equal(config.HostName))
				Expect(result.Version).To(Equal(config.Version))
				Expect(result.Status).To(Equal("10"))
			})
		})
	})
})
