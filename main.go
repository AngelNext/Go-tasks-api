package main

import (
	"fmt"
	"log"

	"github.com/angelnext/tasks-api/database"
	"github.com/angelnext/tasks-api/models"
	"github.com/angelnext/tasks-api/routes"

	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var dbOpenErr error
	if database.DB, dbOpenErr = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DB")))); dbOpenErr != nil {
		log.Fatalln(dbOpenErr)
	}

	if databaseMigrateErr := database.DB.AutoMigrate(&models.Task{}); databaseMigrateErr != nil {
		log.Fatalln("Error while migrating database")
	}

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		AppName:       "Tasks API",
	})

	app.Get("/tasks", routes.GetTasks)
	app.Get("/tasks/:id", routes.GetTask)
	app.Post("/tasks", routes.CreateTask)
	app.Delete("/tasks/:id", routes.DeleteTask)
	app.Delete("/tasks", routes.DeleteTasks)
	app.Put("/tasks/:id", routes.UpdateTask)
	app.Options("*", func(c *fiber.Ctx) error {
		c.Set("Allow", "GET, POST, DELETE, PUT, OPTIONS")
		c.Set("Cache-Control", "max-age=604800")
		c.Set("Access-Control-Allow-Headers", "X-PINGOTHER, Content-Type")
		c.Set("X-Content-Type-Options", "nosniff")
		return nil
	})

	log.Fatalln(app.Listen(":4000"))
}
