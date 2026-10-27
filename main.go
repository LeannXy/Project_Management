package main

import (
	"log"

	"github.com/LeannXy/Project_Management/config"
	"github.com/LeannXy/Project_Management/controllers"
	"github.com/LeannXy/Project_Management/database/seed"
	"github.com/LeannXy/Project_Management/repositories"
	"github.com/LeannXy/Project_Management/routes"
	"github.com/LeannXy/Project_Management/services"
	"github.com/gofiber/fiber/v3"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	seed.SeedAdmin()
	app := fiber.New()
	//user
	userRepo := repositories.NewUserRepository()
	UserService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(UserService)

	//board
	boardRepo := repositories.NewBoardRepository()
	boardMemberRepo := repositories.NewBoardMemberRepository()
	boardService := services.NewBoardService(boardRepo, userRepo, boardMemberRepo)
	boardController := controllers.NewBoardController(boardService)
	
	routes.Setup(app, userController, boardController)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port:", port)
	log.Fatal(app.Listen(":" + port))
}