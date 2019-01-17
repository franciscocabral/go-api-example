package controllers

import (
	"net/http"

	"github.com/franciscocabral/go-api-example/logging"
	"github.com/gin-gonic/gin"
)

// Management is a base structure for the Management Controller
type Management struct{}

type component struct{}

type healthResponse struct {
	Application string      `json:"application"`
	Company     string      `json:"company"`
	Hostname    string      `json:"hostname"`
	Version     string      `json:"version"`
	Status      string      `json:"status"`
	Components  []component `json:"components"`
}

// SetupRoutes configure the controller routes
func (Management) SetupRoutes(rg RouterGroup, relativePath string) {
	rg.Group.GET(relativePath+"/health-check", checkHealth)
}

// checkHealth verify the health of the application
func checkHealth(c *gin.Context) {
	logger.Info("Health-check requested", logging.AdditionalData(map[string]interface{}{
		"url":     c.Request.URL,
		"request": c.Request,
	}))

	response := healthResponse{
		Application: config.Name,
		Company:     config.CompanyName,
		Hostname:    config.HostName,
		Version:     config.Version,
		Status:      "10",
	}

	c.JSON(http.StatusOK, response)
}
