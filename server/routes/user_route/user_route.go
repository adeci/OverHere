package user_route

import (
	"OverHere/server/controllers/user_controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	return gin.Default()
}

func UserRoute(router *gin.Engine) {
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.POST("/user", user_controller.CreateUser())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func Run(router *gin.Engine) {
	router.Run("localhost:8000")
}
