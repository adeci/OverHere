package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handle() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"172.0.0.1"})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
