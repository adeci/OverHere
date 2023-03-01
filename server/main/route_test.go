package main

import (
	"OverHere/server/models"
	"OverHere/server/responses"
	"OverHere/server/routes"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUserRoute(t *testing.T) {
	//Setup
	router := routes.CreateRouter()
	routes.Route(router)

	//Act
	user := models.User{
		UserID:   "123456",
		Username: "Mary Sue",
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users/create", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//Assert
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetUserRoute(t *testing.T) {
	//Setup
	router := routes.CreateRouter()
	routes.Route(router)

	//Act
	req, _ := http.NewRequest("GET", "/users/get/123456", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var userResponse responses.UserResponse
	// Try to return response body into object
	json.Unmarshal(w.Body.Bytes(), &userResponse)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, userResponse)
}

func TestCreateImageRoute(t *testing.T) {
	//Setup
	router := routes.CreateRouter()
	routes.Route(router)

	//Act
	image := models.Image{
		ImageID:  "123456",
		OHPostID: "123456",
		Encoding: "Test",
	}

	jsonValue, _ := json.Marshal(image)
	req, _ := http.NewRequest("POST", "/images/create", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//Assert
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetImageRoute(t *testing.T) {
	//Setup
	router := routes.CreateRouter()
	routes.Route(router)

	//Act
	req, _ := http.NewRequest("GET", "/images/get/123456", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var imageResponse responses.ImageResponse
	// Try to return response body into object
	json.Unmarshal(w.Body.Bytes(), &imageResponse)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, imageResponse)
}
