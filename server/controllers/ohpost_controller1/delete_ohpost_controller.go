package ohpost_controller1

import (
	"OverHere/server/responses"
	"OverHere/server/services/database"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func DeleteOHPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if request isn't processed in 10 seconds
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		ohpostID := c.Param("ohpostid")
		defer cancel()

		fmt.Print("Deleting OHPost " + ohpostID)

		err := database.DeleteOHPost_OHPostID(ohpostID)

		if err != nil {
			c.JSON(http.StatusBadRequest, BadRequestOHPostResponse(err.Error()))
		}

		//Remove OHPost from all images that point to it.
		images, err := database.GetImage_OHPostID(ohpostID)

		for _, image := range images {
			database.PutImage_OHPostID(image.ImageID, "")
		}

		if err == nil {
			c.JSON(http.StatusOK, DeleteOHPostResponse())
		} else {
			c.JSON(http.StatusBadRequest, BadRequestOHPostResponse(err.Error()))
		}
	}
}

func DeleteOHPostResponse() responses.OHPostResponse {
	return responses.OHPostResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    map[string]interface{}{},
	}
}
