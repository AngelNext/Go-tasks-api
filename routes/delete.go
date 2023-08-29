package routes

import (
	"github.com/angelnext/tasks-api/database"
	"github.com/angelnext/tasks-api/models"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

func DeleteTask(c *fiber.Ctx) error {
	db := database.DB
	db.Delete(&models.Task{}, "id = ?", c.Params("id"))
	return c.SendString("Task Deleted Successfully")
}

func DeleteTasks(c *fiber.Ctx) error {
	db := database.DB
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Task{})
	return c.SendString("All Tasks Deleted Successfully")
}
