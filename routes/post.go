package routes

import (
	"time"

	"github.com/angelnext/tasks-api/database"
	"github.com/angelnext/tasks-api/models"

	"github.com/gofiber/fiber/v2"
	nanoid "github.com/matoous/go-nanoid"
)

func CreateTask(c *fiber.Ctx) error {
  db := database.DB
	var task models.Task

	if getBodyReqErr := c.BodyParser(&task); getBodyReqErr != nil {
		return c.Status(503).SendString("Invalid JSON Request Body")
	}
  
  id, idGenerationErr := nanoid.ID(12)

  if idGenerationErr != nil {
    return c.Status(500).SendString("Server Error")
  }

	task.ID = id
  task.CreatedAt = int(time.Now().UnixNano())

	db.Create(&task)
	return c.Status(201).SendString("Task Created Successfully")
}
