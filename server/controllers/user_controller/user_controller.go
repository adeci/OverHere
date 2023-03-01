package user_controller

import (
	"OverHere/server/responses/user_response"
	"net/http"
)

func BadRequestUserResponse(errorData string) user_response.UserResponse {
	return user_response.UserResponse{
		Status:  http.StatusBadRequest,
		Message: "error",
		Data: map[string]interface{}{
			"data": errorData,
		},
	}
}
