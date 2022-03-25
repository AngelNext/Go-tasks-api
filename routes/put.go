package routes

import (
	"github.com/AngelNext/tasks/database"
	"github.com/AngelNext/tasks/models"
	"github.com/gofiber/fiber/v2"
)

func UpdateTask(c *fiber.Ctx) error {
	db := database.DB
	var parsedTask models.Task
	if getReqBodyErr := c.BodyParser(&parsedTask); getReqBodyErr != nil {
		return c.JSON(models.Info{
			Success: false,
			Message: "Invalid JSON Request Body",
		})
	}
	var task models.Task
	db.First(&task, "id = ?", c.Params("id"))
	if parsedTask.Name == "" || parsedTask.Content == "" {
		return c.JSON(models.Info{
			Success: false,
			Message: "Task name and content cannot be empty",
		})
	}
	task.Name = parsedTask.Name
	task.Content = parsedTask.Content
	task.Completed = parsedTask.Completed
	db.Save(&task)
	return c.JSON(models.Info{
		Success: true,
		Message: "Task Updated Successfully",
	})
}
