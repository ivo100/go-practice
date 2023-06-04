package controller_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/ivo100/go-practice/userservice/internal/controller"
	"github.com/ivo100/go-practice/userservice/internal/routes"
	svc "github.com/ivo100/go-practice/userservice/pkg"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setup() *gin.Engine {
	gin.SetMode(gin.TestMode)
	userController := controller.NewUserController()
	router := routes.InitializeRoutes(userController)
	return router
}

func TestGetUserNotFound(t *testing.T) {
	engine := setup()
	// Create get request
	req, _ := http.NewRequest("GET", "/users/99999", nil)
	// Create  response recorder
	w := httptest.NewRecorder()
	// Send the request to the test router
	engine.ServeHTTP(w, req)
	// Check that the response code is correct
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestInvalidRequests(t *testing.T) {
	engine := setup()
	users := []svc.User{
		{ID: "123", FirstName: "john", LastName: ""},
		{ID: "123", FirstName: "", LastName: "doe"},
		{ID: "", FirstName: "john", LastName: "doe"},
		{ID: "", FirstName: "", LastName: ""},
		{ID: "1", FirstName: "2", LastName: "3"},
	}
	for _, user := range users {
		userJson, err := json.Marshal(user)
		assert.NoError(t, err)
		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userJson))
		w := httptest.NewRecorder()
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	}
}

func TestCreateThenGetUser(t *testing.T) {
	engine := setup()
	// Create a test user object
	user := svc.User{ID: "123", FirstName: "John", LastName: "Doe"}
	// Encode the user object as JSON
	userJson, err := json.Marshal(user)
	assert.NoError(t, err)
	// Create post request
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userJson))
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	// read back
	getReq1, _ := http.NewRequest("GET", "/users/123", nil)
	// get recorder
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, getReq1)
	assert.Equal(t, http.StatusOK, w.Code)
	// Decode the response body into a User object
	var actual svc.User
	err = json.NewDecoder(w.Body).Decode(&actual)
	assert.NoError(t, err)
	assert.EqualValues(t, user, actual)
}
