package user_controller

import (
	"OverHere/server/responses"
	"net/http"
)

func BadRequestUserResponse(errorData string) responses.UserResponse {
	return responses.UserResponse{
		Status:  http.StatusBadRequest,
		Message: "error",
		Data: map[string]interface{}{
			"data": errorData,
		},
	}
}
