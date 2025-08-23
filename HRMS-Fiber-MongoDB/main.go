package main

import (
	"log"

	"github.com/AlperSeyman/hrms-fiber-mongodb/database"
	"github.com/AlperSeyman/hrms-fiber-mongodb/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/employee", handlers.GetAllEmployees)
	app.Get("/employee/:id", handlers.GetEmployee)
	app.Post("/employee", handlers.CreateEmployee)
	app.Put("/employee/:id", handlers.UpdateEmployee)
	app.Delete("/employee/:id", handlers.DeleteEmployee)

}

func main() {

	database.ConnectMongoDB()

	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
