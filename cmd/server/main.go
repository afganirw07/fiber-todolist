package main

import (
	"todolist-backend/internal/config"
	"todolist-backend/internal/handler"
	"todolist-backend/internal/model"
	"todolist-backend/internal/repository"
	"todolist-backend/internal/routes"
	"todolist-backend/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.ConnectDatabase()
	config.DB.AutoMigrate(&model.User{}, &model.Task{})

	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	taskRepo := repository.NewTaskRepository(config.DB)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	routes.Setup(app, authHandler, userHandler, taskHandler)

	app.Listen(":3000")
}
