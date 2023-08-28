package routes

import (
	"github.com/AngelNext/tasks-api/database"
	"github.com/AngelNext/tasks-api/models"

	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/ksuid"
)

func CreateTask(c *fiber.Ctx) error {
  db := database.DB
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(503).SendString("Invalid JSON Request Body")
	}

	task.ID = ksuid.New().String()
	db.Create(&task)
	return c.Status(201).SendString("Task Created Successfully")
}
