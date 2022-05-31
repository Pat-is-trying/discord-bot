package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	registerRoutes(r)
	r.Run()
}

func registerRoutes(r *gin.Engine) {
	r.GET("/hello", sayHello)
	r.GET("/goodbye", sayGoodbye)
}

func sayHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello!")
}

func sayGoodbye(c *gin.Context) {
	c.String(http.StatusOK, "Goodbye!")
}