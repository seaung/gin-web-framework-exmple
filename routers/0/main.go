package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin")
	})
	r.Run(":8099")
}
