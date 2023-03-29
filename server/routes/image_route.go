package routes

import (
	"OverHere/server/controllers/image_controller"

	"github.com/gin-gonic/gin"
)

func ImageRoute(router *gin.Engine) {
	router.POST("/images/post", image_controller.PostImage())
	router.GET("/images/get/:imageid", image_controller.GetImage())
	router.PUT("/images/put/:imageid", image_controller.PutImage())
	router.DELETE("/images/delete/:imageid", image_controller.DeleteImage())
}
