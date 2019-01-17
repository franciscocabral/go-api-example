package controllers

import "github.com/gin-gonic/gin"

// Controller defines what is needed to create a controller
type Controller interface {
	SetupRoutes(group RouterGroup, relativePath string)
}

// RouterGroup encapsulates the gin.RouterGroup to uncouple the controller
type RouterGroup struct {
	Group *gin.RouterGroup
}
