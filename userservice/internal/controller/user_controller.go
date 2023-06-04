package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ivo100/go-practice/userservice/internal/core"
	"github.com/ivo100/go-practice/userservice/internal/service"
	svc "github.com/ivo100/go-practice/userservice/pkg"
	"log"
	"net/http"
)

// UserController handles requests for user resources.
type UserController struct {
	userService svc.UserService
}

func NewUserController() *UserController {
	return &UserController{userService: service.NewUserService()}
}

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

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user svc.User
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
