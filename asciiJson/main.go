package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ascii_json", func(c *gin.Context) {
		data := map[string]interface{}{
			"name":       "gin web framework",
			"categories": "golang",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.Run(":8000")
}
