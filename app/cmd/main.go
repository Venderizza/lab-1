package main

import (
	"lab-1/internal/database"
	"lab-1/internal/handlers"
	"lab-1/internal/repo"
	"lab-1/internal/router"
	"lab-1/internal/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.Connect()

	taskRepo := repo.NewTaskRepo(database.DB)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)
	taskRouter := router.NewTaskRouter(app, taskHandler)
	taskRouter.RegisterRoutes()

	app.Listen(":3000")
}
