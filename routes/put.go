package routes

import (
	"github.com/angelnext/tasks-api/database"
	"github.com/angelnext/tasks-api/models"
	"github.com/gofiber/fiber/v2"
)

func UpdateTask(c *fiber.Ctx) error {
	db := database.DB

	var parsedTask models.Task

	if getReqBodyErr := c.BodyParser(&parsedTask); getReqBodyErr != nil {
		return c.Status(503).SendString("Invalid JSON Request Body")
  }

	var task models.Task

	db.First(&task, "id = ?", c.Params("id"))

  if parsedTask.Title != "" {
    task.Title = parsedTask.Title
  } 

  if parsedTask.Description != "" {
    task.Title = parsedTask.Title
  }

	task.Done = parsedTask.Done

	db.Save(&task)

	return c.SendString("Task Updated Successfully")
}
