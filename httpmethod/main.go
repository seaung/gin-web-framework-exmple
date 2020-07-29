package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func get(c *gin.Context) {
	c.String(http.StatusOK, "get")
}

func post(c *gin.Context) {
	c.String(http.StatusOK, "post")
}

func put(c *gin.Context) {
	c.String(http.StatusOK, "put")
}

func main() {
	r := gin.Default()

	r.GET("/get", get)
	r.POST("/post", post)
	r.PUT("/put", put)

	r.Run(":9000")
}
