package routes

import (
	"OverHere/server/controllers/user_controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.POST("/users", user_controller.CreateUser())
	router.GET("/users?userid=:userid", user_controller.GetUser())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
