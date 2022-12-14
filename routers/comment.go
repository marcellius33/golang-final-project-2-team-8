package routers

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func InitCommentRoutes(Routes *gin.Engine, controller *controllers.CommentController) {
	commentRouter := Routes.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/", controller.GetComments)
		commentRouter.POST("/", controller.CreateComment)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controller.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controller.DeleteComment)
	}
}
