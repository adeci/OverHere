package routes

import (
	"OverHere/server/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/users/post", user_controller.PostUser())
	router.GET("/users/get/:userid", user_controller.GetUser())
	router.GET("/users/get/byusername/:username", user_controller.GetUserByUsername())
	router.PUT("/users/put/:userid", user_controller.PutUser())
	router.DELETE("/users/delete/:userid", user_controller.DeleteUser())
	router.DELETE("/users/delete/byusername/:username", user_controller.DeleteUserByUsername())
}
