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
func Test_IMAGE_PostGetPutDelete(t *testing.T) {
	//Setup
	_ = database.ConnectMongoDBAtlas()
	router := routes.CreateRouter()
	routes.Route(router)

	//var inputUsername string = "TESTING-MyUser"
	//var inputUpdatedUsername string = "TESTING-MyNewUser"

	var mockUserID = "TESTING-MockUserID"
	var mockEncoding = "TESTING-MockEncoding"
	var mockXCoord = float64(20)
	var mockYCoord = float64(80)

	//Act
	//**Post**
	image := models.Image{
		UserID:   mockUserID,
		Encoding: mockEncoding,
		XCoord:   mockXCoord,
		YCoord:   mockYCoord,
	}

	postJSONPayload, _ := json.Marshal(image)
	postReq, _ := http.NewRequest("POST", "/images/post", bytes.NewBuffer(postJSONPayload))
	RecordRequest_LookForCode_ImageResponseNotEmpty(postReq, http.StatusCreated, router, t)

	postedImages, err := database.GetImage_UserID(mockUserID)
	assert.Equal(t, nil, err, err)

	//Ensure OHPost is empty
	assert.Equal(t, postedImages[0].OHPostID, "", postedImages[0])

	//**Get**
	getReq, _ := http.NewRequest("GET", "/images/get/"+postedImages[0].ImageID, nil)
	RecordRequest_LookForCode_ImageResponseNotEmpty(getReq, http.StatusOK, router, t)

	//**Put**
	updatedInfoImage := models.Image{
		ImageID:  postedImages[0].ImageID,
		UserID:   "Other" + mockUserID,
		OHPostID: "",
		Encoding: postedImages[0].Base64Encode,
		XCoord:   postedImages[0].XCoord,
		YCoord:   postedImages[0].YCoord,
	}
	putJSONPayload, _ := json.Marshal(updatedInfoImage)
	putReq, _ := http.NewRequest("PUT", "/images/put/"+postedImages[0].ImageID, bytes.NewBuffer(putJSONPayload))
	RecordRequest_LookForCode_ImageResponseNotEmpty(putReq, http.StatusOK, router, t)

	//**Delete**
	deleteReq, _ := http.NewRequest("DELETE", "/images/delete/"+postedImages[0].ImageID, nil)
	RecordRequest_LookForCode_ImageResponseNotEmpty(deleteReq, http.StatusOK, router, t)
}

func Test_IMAGE_GetDeleteMultiple(t *testing.T) {
	//Setup
	_ = database.ConnectMongoDBAtlas()
	router := routes.CreateRouter()
	routes.Route(router)

	var mockUserID = "TESTING-MockUserID"
	var mockEncoding = "TESTING-MockEncoding"
	var mockXCoord1 = float64(100)
	var mockYCoord1 = float64(100)
	var mockXCoord2 = float64(200)
	var mockYCoord2 = float64(200)
	var mockXCoord3 = float64(0)
	var mockYCoord3 = float64(0)
	//var expectedAvgX = float64(100)
	//var expectedAvgY = float64(100)

	//Act
	//**Post**
	//(multiple images)
	image := models.Image{
		UserID:   mockUserID,
		Encoding: mockEncoding,
		XCoord:   mockXCoord1,
		YCoord:   mockYCoord1,
	}

	postJSONPayload, _ := json.Marshal(image)
	postReq, _ := http.NewRequest("POST", "/images/post", bytes.NewBuffer(postJSONPayload))
	RecordRequest_LookForCode_ImageResponseNotEmpty(postReq, http.StatusCreated, router, t)

	image.XCoord = mockXCoord2
	image.YCoord = mockYCoord2

	postJSONPayload, _ = json.Marshal(image)
	postReq, _ = http.NewRequest("POST", "/images/post", bytes.NewBuffer(postJSONPayload))
	RecordRequest_LookForCode_ImageResponseNotEmpty(postReq, http.StatusCreated, router, t)

	image.XCoord = mockXCoord3
	image.YCoord = mockYCoord3

	postReq, _ = http.NewRequest("POST", "/images/post", bytes.NewBuffer(postJSONPayload))
	RecordRequest_LookForCode_ImageResponseNotEmpty(postReq, http.StatusCreated, router, t)

	getMultipleReq, _ := http.NewRequest("GET", "/images/get/byuserid"+mockUserID, nil)
	RecordRequest_LookForCode_ImageResponseNotEmpty(getMultipleReq, http.StatusOK, router, t)

	postedImages, _ := database.GetImage_UserID(mockUserID)
	assert.Equal(t, int(3), len(postedImages), postedImages)

	deleteMultipleReq, _ := http.NewRequest("GET", "/delete/byuserid/"+mockUserID, nil)
	RecordRequest_LookForCode_ImageResponseNotEmpty(deleteMultipleReq, http.StatusOK, router, t)

	postedImages, _ = database.GetImage_UserID(mockUserID)
	assert.Equal(t, int(0), len(postedImages), postedImages)
}

func RecordRequest_LookForCode_ImageResponseNotEmpty(request *http.Request, codeToLookFor int, router *gin.Engine, t *testing.T) {
	var ohpostResponse responses.ImageResponse
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	json.Unmarshal(w.Body.Bytes(), &ohpostResponse)

	assert.Equal(t, codeToLookFor, w.Code, ohpostResponse)
	assert.NotEmpty(t, ohpostResponse, ohpostResponse)
}

/*
func Test_IMAGE_Post(t *testing.T) {
	//Setup
	_ = database.ConnectMongoDBAtlas()
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

func Test_IMAGE_PostGet(t *testing.T) {
	//Setup
	_ = database.ConnectMongoDBAtlas()
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

*/
