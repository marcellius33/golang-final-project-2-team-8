package main

import (
	"os"

	"mygram/database"
	_ "mygram/initializer"

	"mygram/controllers"
	"mygram/repositories"
	"mygram/routers"
	"mygram/services"

	"github.com/gin-gonic/gin"
)

func init() {
	database.Connect()
}

func main() {
	Routes := gin.Default()

	userRepository := repositories.NewUserRepository(database.GetDB())
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	routers.InitUserRoutes(Routes, userController)

	photoRepository := repositories.NewPhotoRepository(database.GetDB())
	photoService := services.NewPhotoService(photoRepository, userRepository)
	photoController := controllers.NewPhotoController(photoService)
	routers.InitPhotoRoutes(Routes, photoController)

	commentRepository := repositories.NewCommentRepository(database.GetDB())
	commentService := services.NewCommentService(commentRepository, userRepository)
	commentController := controllers.NewCommentController(commentService)
	routers.InitCommentRoutes(Routes, commentController)

	// socmedRepository := repositories.NewSocmedRepository(database.GetDB())
	// socmedService := services.NewSocmedService(socmedRepository, userRepository)
	// socmedController := controllers.NewSocialMediaController(socmedService)
	// routers.InitSocmedRouter(Routes, socmedController)

	Routes.Run(os.Getenv("SERVER_PORT"))
}
