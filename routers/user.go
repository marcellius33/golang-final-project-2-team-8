package routers

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(Routes *gin.Engine, controller *controllers.UserController) {
	userRouter := Routes.Group("/users")
	{
		
		userRouter.POST("/register", controller.UserRegisterController)
		userRouter.POST("/login", controller.UserLoginController)
		userRouter.Use(middlewares.Authentication())
		userRouter.PUT("/:userId", middlewares.UserAuthorization(), controller.UserUpdateController)
		userRouter.DELETE("/:userId", middlewares.UserAuthorization(), controller.DeleteUserController)
	}
}