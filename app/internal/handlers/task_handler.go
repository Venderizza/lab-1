package handlers

import (
	"lab-1/internal/dto"
	"lab-1/internal/services"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	taskService *services.TaskService
}

func NewTaskHandler(taskService *services.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var taskdto dto.CreateTaskDTO

	if err := c.BodyParser(&taskdto); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Invalid request body"})
	}

	task, err := h.taskService.CreateTask(ctx, &taskdto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Failed to create task"})
	}

	return c.Status(fiber.StatusCreated).JSON(
		dto.NewTaskResponseDTO(*task),
	)
}

func (h *TaskHandler) GetTaskByID(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Invalid task ID"})
	}

	taskResponse, err := h.taskService.GetTaskByID(ctx, uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Failed to retrieve task"})
	}
	if taskResponse.ID == 0 {
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{"error": "Task not found"})
	}

	return c.JSON(taskResponse)
}

func (h *TaskHandler) GetAllTasks(c *fiber.Ctx) error {
	ctx := c.UserContext()

	tasks, err := h.taskService.GetAllTasks(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Failed to retrieve tasks"})
	}

	return c.JSON(tasks)
}

func (h *TaskHandler) GetTasksByPriority(c *fiber.Ctx) error {
	ctx := c.UserContext()
	priority, err := c.ParamsInt("priority")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid priority value"})
	}

	tasks, err := h.taskService.GetTasksByPriority(ctx, priority)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Failed to retrieve tasks"})
	}

	return c.JSON(tasks)
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Invalid task ID"})
	}

	var taskdto dto.UpdateTaskDTO
	if err := c.BodyParser(&taskdto); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Invalid request body"})
	}

	err = h.taskService.UpdateTask(ctx, uint(id), &taskdto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Failed to update task"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *TaskHandler) ToggleTask(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Invalid task ID"})
	}

	err = h.taskService.ToggleTask(ctx, uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Failed to toggle task status"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Invalid task ID"})
	}

	err = h.taskService.DeleteTask(ctx, uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Failed to delete task"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
