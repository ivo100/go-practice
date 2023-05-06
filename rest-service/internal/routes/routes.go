package routes

import (
	"github.com/gin-gonic/gin"
	"userservice/internal/controller"
)

func InitializeRoutes(userCtl *controller.UserController) *gin.Engine {

	// Set Gin to production mode
	//gin.SetMode(gin.ReleaseMode)

	gin.SetMode(gin.DebugMode)

	var router = gin.Default()

	// Group user related routes together
	users := router.Group("/users")
	{
		users.GET("/:id", userCtl.GetUser)
		users.POST("", userCtl.CreateUser)
		//userRoutes.PUT("/:id", userCtl.UpdateUser)
		//userRoutes.DELETE("/:id", userCtl.DeleteUser)
	}

	_ = router.SetTrustedProxies(nil)

	return router
}
