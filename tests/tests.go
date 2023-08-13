package tests

//
//import (
//	"net/http"
//	"net/http/httptest"
//	"testing"
//
//	"github.com/gin-gonic/gin"
//	"github.com/shirrashko/BuildingAServer-step2/pkg/api"
//	"github.com/stretchr/testify/assert"
//)
//
//func TestCreateUserProfile(t *testing.T) {
//	// Set up a mock Gin router
//	router := gin.Default()
//	h := api.Router().handlers[0:1]
//	router.POST("/users", h.createUserProfile)
//
//	// Create a request to the endpoint
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("POST", "/users", nil)
//	router.ServeHTTP(w, req)
//
//	// Assertions
//	assert.Equal(t, http.StatusCreated, w.Code)
//	// You can add more assertions here to check the response body or any other expected behavior
//}
//
//func TestGetUserProfile(t *testing.T) {
//	// Set up a mock Gin router
//	router := gin.Default()
//	router.GET("/users/:id", api.getAUserProfile)
//
//	// Create a request to the endpoint
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("GET", "/users/1", nil)
//	router.ServeHTTP(w, req)
//
//	// Assertions
//	assert.Equal(t, http.StatusOK, w.Code)
//	// You can add more assertions here to check the response body or any other expected behavior
//}
