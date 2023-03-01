package routes

import (
	"OverHere/server/controllers/image_controller"

	"github.com/gin-gonic/gin"
)

func ImageRoute(router *gin.Engine) {
	router.POST("/images/create", image_controller.CreateImage())
	router.GET("/images/get/:imageid", image_controller.GetImage())
}
