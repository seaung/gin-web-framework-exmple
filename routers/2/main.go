package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "seaung")
		hobby := c.Query("hobby")
		c.String(http.StatusOK, fmt.Sprintf("hello %s hobby is %s", name, hobby))
	})
	r.Run(":8099")
}
