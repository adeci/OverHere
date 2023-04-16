package routes

import (
	"OverHere/server/controllers/ohpost_controller1"

	"github.com/gin-gonic/gin"
)

func OHPostRoute(router *gin.Engine) {
	router.POST("/ohpost/post", ohpost_controller1.PostOHPost())
	router.POST("/ohpost/post/withimageids", ohpost_controller1.PostOHPostWithImageIds())
	router.GET("/ohpost/get/:ohpostid", ohpost_controller1.GetOHPost())
	router.GET("/ohpost/get/bycoordboundary/:topleftXcoord/:topleftYcoord/:botrightXcoord/:botrightYcoord", ohpost_controller1.GetOHPostsByCoordBoundary())
	router.GET("/ohpost/get/byuserid/:userid", ohpost_controller1.GetOHPostsByUserId())
	router.PUT("/ohpost/put/:ohpostid", ohpost_controller1.PutOHPost())
	router.DELETE("/ohpost/delete/:ohpostid", ohpost_controller1.DeleteOHPost())
}
