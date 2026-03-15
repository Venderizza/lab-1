package router

import (
	"lab-1/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

type TaskRouter struct {
	app     *fiber.App
	handler *handlers.TaskHandler
}

func NewTaskRouter(app *fiber.App, handler *handlers.TaskHandler) *TaskRouter {
	return &TaskRouter{app: app, handler: handler}
}

func (r *TaskRouter) RegisterRoutes() {
	taskGroup := r.app.Group("/api/tasks")

	taskGroup.Post("/", r.handler.CreateTask)
	taskGroup.Get("/", r.handler.GetAllTasks)
	taskGroup.Get("/:id", r.handler.GetTaskByID)
	taskGroup.Get("/priority/:priority", r.handler.GetTasksByPriority)
	taskGroup.Put("/:id", r.handler.UpdateTask)
	taskGroup.Patch("/:id/toggle", r.handler.ToggleTask)
	taskGroup.Delete("/:id", r.handler.DeleteTask)
}
