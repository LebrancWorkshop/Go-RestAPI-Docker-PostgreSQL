package main 

import (
	"github.com/gofiber/fiber/v2"
	"github.com/LebrancWorkshop/Go-RestAPI-Docker-PostgreSQL/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home); 
	app.Get("/quizes", handlers.ReadQuiz); 
	app.Post("/quizes", handlers.CreateQuiz); 
}