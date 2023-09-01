package routes

import (
	"github.com/angelnext/tasks/database"
	"github.com/angelnext/tasks/models"
	"github.com/gofiber/fiber/v2"
)

func GetTasks(c *fiber.Ctx) error {
	db := database.DB
	var tasks []models.Task
	db.Order("created_at DESC, id DESC").Find(&tasks)
	return c.JSON(tasks)
}

func GetTask(c *fiber.Ctx) error {
	db := database.DB
	var task models.Task
	db.Find(&task, "id = ?", c.Params("id"))
	if task.ID != "" {
		return c.JSON(task)
	}
	return c.Status(404).JSON(fiber.Map{})
}
