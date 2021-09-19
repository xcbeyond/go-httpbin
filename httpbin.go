package main

import (
	"go-httpbin/router/method"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// create gin object
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	methodsAPI := router.Group("/methods")
	{
		methodsAPI.GET("/get", method.Get)

		methodsAPI.POST("/post", method.Post)

		methodsAPI.PUT("/put", method.Put)

		methodsAPI.PATCH("/patch", method.Patch)

		methodsAPI.DELETE("/delete", method.Delete)
	}

	router.Run(":8080")
}
