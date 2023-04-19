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

	"github.com/stretchr/testify/assert"
)

func Test_USER_PostGetPutDelete(t *testing.T) {
	//Setup
	_ = database.ConnectMongoDBAtlas()
	router := routes.CreateRouter()
	routes.Route(router)
	w := httptest.NewRecorder()

	//Act

	//**Delete old user**
	deleteJSONPayload, _ := json.Marshal("")
	deleteReq, _ := http.NewRequest("DELETE", "/users/delete/USER-123", bytes.NewBuffer(deleteJSONPayload))

	router.ServeHTTP(w, deleteReq)
	//Don't assert check, doesn't matter if deleted or not

	//**Post**
	user := models.User{
		Username: "USER-123",
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users/post", bytes.NewBuffer(jsonValue))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	//Ensure returned username is correct

	//**Get**
	var response responses.UserResponse
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusCreated, response.Status)

	userId := response.Data["data"]

	//**Get**
	/*
		user := models.User{
			Username: "MyNewUser",
		}

		jsonValue, _ := json.Marshal(user)
		req, _ := http.NewRequest("POST", "/users/get/", bytes.NewBuffer(jsonValue))

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		//Assert
		assert.Equal(t, http.StatusCreated, w.Code)
	*/
}
func Test_USER_Post(t *testing.T) {
	//Setup
	_ = database.ConnectMongoDBAtlas()
	router := routes.CreateRouter()
	routes.Route(router)

	//Act
	user := models.User{
		Username: "Test2",
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users/post", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//Assert
	assert.Equal(t, http.StatusCreated, w.Code)
}

func Test_USER_Get(t *testing.T) {
	//Setup
	_ = database.ConnectMongoDBAtlas()
	router := routes.CreateRouter()
	routes.Route(router)

	//Act
	req, _ := http.NewRequest("GET", "/users/get/Test1", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var userResponse responses.UserResponse
	// Try to return response body into object
	json.Unmarshal(w.Body.Bytes(), &userResponse)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, userResponse)
}
