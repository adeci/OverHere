package image_controller

import (
	"OverHere/server/services/database"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func DeleteImagesByOHPostId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if enough time passes.
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		ohpostID := c.Param("ohpostid")
		defer cancel()

		fmt.Print("Deleting image by UserID: " + ohpostID)

		err := database.DeleteImage_OHPostID(ohpostID)

		if err == nil {
			c.JSON(http.StatusOK, DeleteMultipleImagesResponse())
		} else {
			c.JSON(http.StatusBadRequest, BadRequestImageResponse(err.Error()))
		}
	}
}
