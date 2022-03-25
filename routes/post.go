package routes

import (
	"github.com/AngelNext/tasks/database"
	"github.com/AngelNext/tasks/models"

	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/ksuid"
)

func CreateTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(503).JSON(models.Info{
			Success: false,
			Message: "Invalid JSON Request Body",
		})
	}
	task.ID = ksuid.New().String()
	database.DB.Create(&task)
	return c.Status(201).JSON(models.Info{
		Success: true,
		Message: "Task Created Successfully",
	})
}
