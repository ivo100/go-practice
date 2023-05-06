package service

import (
	"github.com/gin-gonic/gin"
	"myservice/internal/core"
	myservice "myservice/pkg"
	"net/http"
)

// UserService provides a REST API for working with users.
type UserService struct {
	repo core.UserRepository
}

// GetUser returns the user with the given ID, or an error if the user is not found.
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags users
// @Param id query string true "ID of the user to get"
// @Success 200 {object} User
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /users [get]
func (service *UserService) GetUser(c *gin.Context) {
	id := c.Query("id")
	user, err := service.repo.GetUser(id)
	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// AddUser adds a new user to the system.
// @Summary Add a user
// @Description Add a new user
// @Tags users
// @Accept json
// @Param user body User true "The user to add"
// @Success 201 {string} string "User created"
// @Failure 400 {string} string "Invalid request"
// @Failure 409 {string} string "User already exists"
// @Failure 500 {string} string "Internal server error"
// @Router /users [post]
func (service *UserService) AddUser(c *gin.Context) {
	var user myservice.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.repo.SaveUser(user); err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func getStatusCode(err error) int {
	return 400
}
