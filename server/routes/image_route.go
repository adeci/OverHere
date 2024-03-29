package routes

import (
	"OverHere/server/controllers/image_controller"

	"github.com/gin-gonic/gin"
)

func ImageRoute(router *gin.Engine) {
	router.POST("/images/post", image_controller.PostImage())
	router.GET("/images/get/:imageid", image_controller.GetImage())
	router.GET("/images/get/byuserid/:userid", image_controller.GetImagesByUserId())
	router.GET("/images/get/byohpostid/:ohpostid", image_controller.GetImagesByOHPostId())
	router.GET("/images/get/bytag/:tag", image_controller.GetImagesByTag())
	router.PUT("/images/put/:imageid", image_controller.PutImage())
	router.PUT("/images/put/addtoohpost/:imageid/:ohpostid", image_controller.PutAddImageToOHPost())
	router.DELETE("/images/delete/:imageid", image_controller.DeleteImage())
	router.DELETE("/images/delete/byuserid/:userid", image_controller.DeleteImagesByUserId())
	router.DELETE("/images/delete/byohpostid/:ohpostid", image_controller.DeleteImagesByOHPostId())
}
