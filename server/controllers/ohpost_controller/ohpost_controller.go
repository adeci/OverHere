package ohpost_controller

import (
	"OverHere/server/responses"
	"net/http"
)

func BadRequestOHPostResponse(errorData string) responses.OHPostResponse {
	return responses.OHPostResponse{
		Status:  http.StatusBadRequest,
		Message: "error",
		Data: map[string]interface{}{
			"data": errorData,
		},
	}
}
