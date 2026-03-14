package handlers

import (
	"lab-1/internal/dto"
	"lab-1/internal/services"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	taskService services.TaskService
}

func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	var taskdto dto.CreateTaskDTO

	if err := c.BodyParser(&taskdto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	task, err := h.taskService.CreateTask(&taskdto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create task"})
	}

	return c.Status(fiber.StatusCreated).JSON(
		dto.NewTaskResponseDTO(task),
	)
}
