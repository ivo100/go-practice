package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"userservice/internal/core"
	"userservice/internal/service"
	userservice "userservice/pkg"
)

// UserController handles requests for user resources.
type UserController struct {
	userService userservice.UserService
}

func NewUserController() *UserController {
	return &UserController{userService: service.NewUserService()}
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
func (ctrl *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	log.Printf("UserController GetUser, id %v", id)
	if id == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	user, err := ctrl.userService.GetUser(id)
	if err == core.ErrNotFound || user == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser adds a new user to the system.
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
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user userservice.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	usr, err := ctrl.userService.GetUser(user.ID)
	if usr != nil {
		c.AbortWithStatus(http.StatusConflict)
		return
	}
	createdUser, err := ctrl.userService.AddUser(user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}
