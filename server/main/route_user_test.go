package main

import (
	"OverHere/server/models"
	"OverHere/server/responses"
	"OverHere/server/routes"
	"OverHere/server/services/database"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Integration tests
func Test_USER_PostGetPutDelete(t *testing.T) {
	//Setup
	_ = database.ConnectMongoDBAtlas()
	router := routes.CreateRouter()
	routes.Route(router)
	w := httptest.NewRecorder()

	var inputUsername string = "TESTING-MyUser"
	var inputUpdatedUsername string = "TESTING-MyNewUser"

	//Act
	//**Delete old user**
	deleteReq, _ := http.NewRequest("DELETE", "/users/delete/byusername/"+inputUsername, nil)

	router.ServeHTTP(w, deleteReq)
	//Don't assert check, doesn't matter if deleted or not

	//**Post**
	user := models.User{
		Username: inputUsername,
	}

	postJSONPayload, _ := json.Marshal(user)
	postReq, _ := http.NewRequest("POST", "/users/post", bytes.NewBuffer(postJSONPayload))
	RecordRequest_LookForCode_UserResponseNotEmpty(postReq, http.StatusCreated, router, t)

	postedUser, err := database.GetUser_Username(inputUsername)
	assert.Equal(t, nil, err, err)

	//**Get**
	getReq, _ := http.NewRequest("GET", "/users/get/"+postedUser.UserID, nil)
	RecordRequest_LookForCode_UserResponseNotEmpty(getReq, http.StatusOK, router, t)

	//**Put**
	updatedInfoUser := models.User{
		Username: inputUpdatedUsername,
	}
	putJSONPayload, _ := json.Marshal(updatedInfoUser)
	putReq, _ := http.NewRequest("PUT", "/users/put/"+postedUser.UserID, bytes.NewBuffer(putJSONPayload))
	RecordRequest_LookForCode_UserResponseNotEmpty(putReq, http.StatusOK, router, t)

	//**Delete**
	deleteReq, _ = http.NewRequest("DELETE", "/users/delete/"+postedUser.UserID, nil)
	RecordRequest_LookForCode_UserResponseNotEmpty(deleteReq, http.StatusOK, router, t)
}

func RecordRequest_LookForCode_UserResponseNotEmpty(request *http.Request, codeToLookFor int, router *gin.Engine, t *testing.T) {
	var userResponse responses.UserResponse
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	json.Unmarshal(w.Body.Bytes(), &userResponse)

	assert.Equal(t, codeToLookFor, w.Code, userResponse)
	assert.NotEmpty(t, userResponse, userResponse)
}

func Unmarshal_UserResponse(response responses.UserResponse) (models.User, error) {
	x, _ := json.Marshal(response.Data["data"])
	y, _ := json.Marshal(x)

	var returnedUser models.User
	err := json.Unmarshal(y, &returnedUser)

	//Ensure returned username is correct
	//json.Unmarshal(w.Body.Bytes(), &userResponse)

	//var target interface{}
	//json.NewDecoder(w.Body).Decode(target)

	//x, _ := json.Marshal(target)
	//y, _ := json.Marshal(x)

	//var returnedUser models.User
	//_ = json.Unmarshal(x, &returnedUser)

	return returnedUser, err
}
