package tests

import (
	"assessment/controller"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/find-pairs", controller.PairedNumberAPI)
	return r
}

func TestFindPairSuccess(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	reqBody := `{"numbers": [1, 2, 3, 4, 5], "target": 6}`
	req, _ := http.NewRequest("POST", "/find-pairs", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	expectedResponse := `{"solutions":[[1,4],[0,5]]}`
	assert.Equal(t, expectedResponse, w.Body.String())
}

func TestFindPairValidationError(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	reqBody := `{"numbers": [1, 2, 3, 4, 5]}`
	req, _ := http.NewRequest("POST", "/find-pairs", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	expectedResponse := `{"error":"Key: 'RequestPayload.Target' Error:Field validation for 'Target' failed on the 'required' tag"}`
	assert.Equal(t, expectedResponse, w.Body.String())
}

func TestFindPairValidationError2(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	reqBody := `{"target": 6}`
	req, _ := http.NewRequest("POST", "/find-pairs", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	expectedResponse := `{"error":"Key: 'RequestPayload.Numbers' Error:Field validation for 'Numbers' failed on the 'required' tag"}`
	assert.Equal(t, expectedResponse, w.Body.String())
}
