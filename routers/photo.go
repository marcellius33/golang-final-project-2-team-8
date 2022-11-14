package routers

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func InitPhotoRoutes(Routes *gin.Engine, controller *controllers.PhotoController) {
	photoRouter := Routes.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/", controller.GetPhotos)
		photoRouter.POST("/", controller.CreatePhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controller.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controller.DeletePhoto)
	}
}
