package routes

import (
	"OverHere/server/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/users/post", user_controller.PostUser())
	router.GET("/users/get/:userid", user_controller.GetUser())
	router.PUT("/users/put/:userid", user_controller.PutUser())
	router.GET("/users/delete/:userid", user_controller.DeleteUser())
}
