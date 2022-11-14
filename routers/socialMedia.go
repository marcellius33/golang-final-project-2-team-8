package routers

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func InitSocmedRouter(Routes *gin.Engine, controller *controllers.SocialMediaController) {
	socmedRouter := Routes.Group("/socialmedias")
	{
		socmedRouter.Use(middlewares.Authentication())
		socmedRouter.GET("/", controller.GetSocialMedias)
		socmedRouter.POST("/", controller.CreateSocialMedia)
		socmedRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controller.UpdateSocialMedia)
		socmedRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controller.DeleteSocialMedia)
	}
}
