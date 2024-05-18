package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbalmahad/post-app.git/backend/controllers"
	"github.com/iqbalmahad/post-app.git/backend/repositories"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	postRepository := repositories.NewPostRepository(db)
	postController := controllers.NewPostController(postRepository)

	app.Post("/posts", postController.CreatePost)
	app.Get("/posts", postController.GetPosts)
	app.Get("/posts/:id", postController.GetPost)
	app.Put("/posts/:id", postController.UpdatePost)
	app.Delete("/posts/:id", postController.DeletePost)
}
