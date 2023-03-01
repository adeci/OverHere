package routes

import "github.com/gin-gonic/gin"

func CreateRouter() *gin.Engine {
	return gin.Default()
}

func Route(router *gin.Engine) {
	UserRoute(router)
	ImageRoute(router)
}

func Run(router *gin.Engine) {
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Run("localhost:8000")
}
