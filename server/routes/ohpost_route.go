package routes

import (
	"OverHere/server/controllers/ohpost_controller"

	"github.com/gin-gonic/gin"
)

func OHPostRoute(router *gin.Engine) {
	router.POST("/ohpost/post", ohpost_controller.PostOHPost())
	router.GET("/ohpost/get/:ohpostid", ohpost_controller.GetOHPost())
	router.PUT("/ohpost/put/:ohpostid", ohpost_controller.PutOHPost())
	router.DELETE("/ohpost/delete/:ohpostid", ohpost_controller.DeleteOHPost())
}
