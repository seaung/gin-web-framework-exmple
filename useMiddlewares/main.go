package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "get")
	})
	r.Run(":9000")
}
