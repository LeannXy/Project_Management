package routes

import (
	"log"

	"github.com/LeannXy/Project_Management/controllers"
	"github.com/LeannXy/Project_Management/middleware"
	
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func Setup(app *fiber.App,uc *controllers.UserController) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app.Post("/v1/auth/register", uc.Register)
	app.Post("/v1/auth/login", uc.Login)

	//JWT protected routes
	api := app.Group("/api/v1", middleware.AuthMiddleware)

	userGroup := api.Group("/users")
	userGroup.Get("/page", uc.GetUserPagination)// http://127.0.0.1:3030/api/v1/users/page?filter=&sort=-id&page=1&limit=20
	userGroup.Get("/:id", uc.GetUser)// /api/v1/users/:id
	userGroup.Put("/:id", uc.UpdateUser)// /api/v1/users/:id
	
}