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
func Test_OHPOST_PostGetPutDelete(t *testing.T) {
	//Setup
	_ = database.ConnectMongoDBAtlas()
	router := routes.CreateRouter()
	routes.Route(router)

	//var inputUsername string = "TESTING-MyUser"
	//var inputUpdatedUsername string = "TESTING-MyNewUser"

	var mockUserID = "TESTING-MockUserID"
	//Act
	//**Post**
	ohpost := models.OHPost{
		UserID:  mockUserID,
		Tag:     "My Tag",
		Caption: "My Caption",
	}

	postJSONPayload, _ := json.Marshal(ohpost)
	postReq, _ := http.NewRequest("POST", "/ohpost/post", bytes.NewBuffer(postJSONPayload))
	RecordRequest_LookForCode_OHPostResponseNotEmpty(postReq, http.StatusCreated, router, t)

	postedOHPosts, err := database.GetOHPost_UserID(mockUserID)
	assert.Equal(t, nil, err, err)

	//Ensure freshly created OHPosts with no images are 0
	assert.Equal(t, float64(0), ohpost.AvgXCoord, postedOHPosts[0])
	assert.Equal(t, float64(0), ohpost.AvgYCoord, postedOHPosts[0])

	//**Get**
	getReq, _ := http.NewRequest("GET", "/ohpost/get/"+postedOHPosts[0].OHPostID, nil)
	RecordRequest_LookForCode_OHPostResponseNotEmpty(getReq, http.StatusOK, router, t)

	//**Put**
	updatedInfoOHPost := models.OHPost{
		UserID:  "Other" + mockUserID,
		Tag:     postedOHPosts[0].Tag,
		Caption: postedOHPosts[0].Description,
	}
	putJSONPayload, _ := json.Marshal(updatedInfoOHPost)
	putReq, _ := http.NewRequest("PUT", "/ohpost/put/"+postedOHPosts[0].OHPostID, bytes.NewBuffer(putJSONPayload))
	RecordRequest_LookForCode_OHPostResponseNotEmpty(putReq, http.StatusOK, router, t)

	//**Delete**
	deleteReq, _ := http.NewRequest("DELETE", "/ohpost/delete/"+postedOHPosts[0].OHPostID, nil)
	RecordRequest_LookForCode_OHPostResponseNotEmpty(deleteReq, http.StatusOK, router, t)
}

func RecordRequest_LookForCode_OHPostResponseNotEmpty(request *http.Request, codeToLookFor int, router *gin.Engine, t *testing.T) {
	var ohpostResponse responses.OHPostResponse
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	json.Unmarshal(w.Body.Bytes(), &ohpostResponse)

	assert.Equal(t, codeToLookFor, w.Code, ohpostResponse)
	assert.NotEmpty(t, ohpostResponse, ohpostResponse)
}
