package image_controller

import (
	"OverHere/server/responses"
	"net/http"
)

func BadRequestImageResponse(errorData string) responses.ImageResponse {
	return responses.ImageResponse{
		Status:  http.StatusBadRequest,
		Message: "error",
		Data: map[string]interface{}{
			"data": errorData,
		},
	}
}
