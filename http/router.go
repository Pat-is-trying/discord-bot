package http

import (
	"fmt"
	"net/http"
	"strings"

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
	r.POST("/authorize", auth)
}

func sayHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello!")
}

func sayGoodbye(c *gin.Context) {
	c.String(http.StatusOK, "Goodbye!")
}

func auth(c *gin.Context) {
	headers := c.Request.Header
	xSig := strings.Join(headers["X-Signature-Ed25519"], "")
	tstmp := strings.Join(headers["X-Signature-Timestamp"], "")
	fmt.Println(xSig)
	fmt.Printf("t5: %T\n", xSig)
	fmt.Println(tstmp)
	fmt.Printf("t5: %T\n", tstmp)
	a := authorize{Type: 1}
	c.Header("X-Signature-Ed25519", xSig)
	c.Header("X-Signature-Timestamp", tstmp)
	c.JSON(http.StatusOK, a)
}

type authorize struct {
	Type int `json:"type"`
}