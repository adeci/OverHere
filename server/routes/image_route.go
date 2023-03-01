package routes

import (
	"OverHere/server/controllers/image_controller"

	"github.com/gin-gonic/gin"
)

func ImageRoute(router *gin.Engine) {
	router.POST("/images", image_controller.CreateImage())
	router.GET("/images?imageid=:userid", image_controller.GetImage())
}
