package server

import (
	"time"

	"github.com/franciscocabral/go-api-example/controllers"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"

	"github.com/franciscocabral/go-api-example/server/middlewares"
	//"github.com/stone-payments/receipt-generator/server/handlers"
)

// ApiServer - is the structure that defines the api server
type APIServer struct {
	Engine *gin.Engine
}

// Setup creates a new instance of APIServer and configure routes.
func Setup() *APIServer {

	server := APIServer{Engine: gin.Default()}
	server.addRoutes()

	return &server
}

// Run starts the server.
func (server *APIServer) Run() error {
	return server.Engine.Run()
}

// setMiddlewares configure all used middlewares.
func (server *APIServer) setMiddlewares(router *gin.RouterGroup) {

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE, PATCH",
		RequestHeaders:  "Origin, Authorization, Content-Type, X-Token",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.Use(middlewares.Logger)
}

// addRoutes configure routes.
func (server *APIServer) addRoutes() {
	root := server.Engine.Group("/")
	//v1 := server.Engine.Group("/v1")

	server.setMiddlewares(root)
	server.addController(root, "/management", new(controllers.Management))
}

func (server *APIServer) addController(group *gin.RouterGroup, relativePath string, controller controllers.Controller) {
	rg := controllers.RouterGroup{
		Group: group,
	}

	controller.SetupRoutes(rg, relativePath)
}
