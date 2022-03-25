package main

import (
	"fmt"
	"log"

	"github.com/AngelNext/tasks/database"
	"github.com/AngelNext/tasks/models"
	"github.com/AngelNext/tasks/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func main() {
	var dbOpenErr error
	if database.DB, dbOpenErr = gorm.Open(mysql.Open(fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/gotasks", os.Getenv("MYSQL_PASSWORD")))); dbOpenErr != nil {
		log.Fatal(dbOpenErr)
	}
	if databaseMigrateErr := database.DB.AutoMigrate(&models.Task{}); databaseMigrateErr != nil {
		log.Fatal("Error while migrating database")
	}
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		Prefork:       true,
		AppName:       "Tasks API",
	})
	app.Static("/", "./public")
	app.Get("/", routes.Home)
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
		return nil
	})
	app.Use(func(c *fiber.Ctx) error {
		c.Set("X-Content-Type-Options", "nosniff")
		return c.Next()
	})
	log.Fatal(app.Listen(":4000"))
}
