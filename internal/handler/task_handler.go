package handler

import (
	"todolist-backend/internal/model"
	"todolist-backend/internal/service"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	Service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{Service: service}
}

func (h *TaskHandler) Create(c *fiber.Ctx) error {
	user := c.Locals("user").(*fiber.Map)
	userID := uint((*user)["user_id"].(float64))

	task := new(model.Task)
	c.BodyParser(task)
	task.UserID = userID

	h.Service.Create(task)
	return c.JSON(task)
}

func (h *TaskHandler) GetMyTasks(c *fiber.Ctx) error {
	user := c.Locals("user").(*fiber.Map)
	userID := uint((*user)["user_id"].(float64))

	var tasks []model.Task
	h.Service.GetByUser(userID, &tasks)

	return c.JSON(tasks)
}
