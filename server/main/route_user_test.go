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

func TestPostUserRoute(t *testing.T) {
	_ = database.ConnectMongoDBAtlas()

	//Setup
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

func TestGetUserRoute(t *testing.T) {
	_ = database.ConnectMongoDBAtlas()

	//Setup
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
