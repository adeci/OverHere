package routes

import (
	"OverHere/server/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/users/post", user_controller.CreateUser())
	router.GET("/users/get/:userid", user_controller.GetUser())

	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
}
