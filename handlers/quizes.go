package handlers 

import (
	"github.com/gofiber/fiber/v2"
	"github.com/LebrancWorkshop/Go-RestAPI-Docker-PostgreSQL/models"
	"github.com/LebrancWorkshop/Go-RestAPI-Docker-PostgreSQL/database"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Test Routes."); 
}

func ReadQuiz(c *fiber.Ctx) error {
	quizes := []models.Quiz{};
	database.DB.Db.Find(&quizes);
	return c.Status(200).JSON(quizes); 
}

func CreateQuiz(c *fiber.Ctx) error {
	quiz := new(models.Quiz);

	if err := c.BodyParser(quiz); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&quiz);
	return c.Status(200).JSON(quiz);
}