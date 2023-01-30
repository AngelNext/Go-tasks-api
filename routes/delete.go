package routes

import (
	"github.com/AngelNext/tasks-api/database"
	"github.com/AngelNext/tasks-api/models"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

func DeleteTask(c *fiber.Ctx) error {
	db := database.DB
	db.Delete(&models.Task{}, "id = ?", c.Params("id"))
	return c.JSON(models.Info{
		Success: true,
		Message: "Task Deleted Successfully",
	})
}

func DeleteTasks(c *fiber.Ctx) error {
	db := database.DB
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Task{})
	return c.JSON(models.Info{
		Success: true,
		Message: "All Tasks Deleted Successfully",
	})
}
