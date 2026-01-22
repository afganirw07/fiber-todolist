package routes

import (
	"todolist-backend/internal/handler"
	"todolist-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(
	app *fiber.App,
	auth *handler.AuthHandler,
	user *handler.UserHandler,
	task *handler.TaskHandler,
) {
	api := app.Group("/api")

	api.Post("/register", auth.Register)
	api.Post("/login", auth.Login)

	api.Get("/users", user.GetAll)

	taskRoute := api.Group("/tasks", middleware.JWTProtected())
	taskRoute.Post("/", task.Create)
	taskRoute.Get("/", task.GetMyTasks)
}
