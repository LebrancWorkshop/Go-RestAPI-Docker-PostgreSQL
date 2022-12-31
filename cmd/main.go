package main 

import (
	"github.com/gofiber/fiber/v2"
	"github.com/LebrancWorkshop/Go-RestAPI-Docker-PostgreSQL/database"
)

func getIndex(c *fiber.Ctx) error {
	return c.SendString("Test API.");   
}

func main() {
	database.ConnectDb(); 
	app := fiber.New();
	app.Get("/", getIndex);
	app.Listen(":8000");
}