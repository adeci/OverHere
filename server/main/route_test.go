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

func TestPostUserRoute(t *testing.T) {
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

func TestPostImageRoute(t *testing.T) {
	//Setup
	router := routes.CreateRouter()
	routes.Route(router)

	//Act
	image := models.Image{
		UserID:   "Test2",
		OHPostID: "Test3",
		Encoding: "Test4",
		XCoord:   29.649934,
		YCoord:   82.348655,
	}

	jsonValue, _ := json.Marshal(image)
	req, _ := http.NewRequest("POST", "/images/post", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//Assert
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPostAndGetImageRoute(t *testing.T) {
	//Setup
	router := routes.CreateRouter()
	routes.Route(router)

	//Act
	image := models.Image{
		UserID:   "Test2",
		OHPostID: "Test3",
		Encoding: "Test4",
		XCoord:   29.649934,
		YCoord:   82.348655,
	}

	jsonValue, _ := json.Marshal(image)
	req, _ := http.NewRequest("POST", "/images/post", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//Assert
	assert.Equal(t, http.StatusCreated, w.Code)

	//Act
	req2, _ := http.NewRequest("GET", "/images/get/123456", nil)

	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	var imageResponse responses.ImageResponse
	// Try to return response body into object
	json.Unmarshal(w.Body.Bytes(), &imageResponse)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, imageResponse)
}
