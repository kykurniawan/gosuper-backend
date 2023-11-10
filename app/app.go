package app

import (
	"gosuper/app/exception"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

type App struct {
	Fiber    *fiber.App
	Database *gorm.DB
}

func NewApp(db *gorm.DB) *App {
	fiber := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
		ErrorHandler:      exception.GlobalErrorHandler,
	})

	fiber.Use(recover.New())

	return &App{
		Fiber:    fiber,
		Database: db,
	}
}

func (app *App) Run() {
	authService := InitializeAuthService(app.Database)
	userService := InitializeUserService(app.Database)

	authController := InitializeAuthController(authService)
	userController := InitializeUserController(userService)

	api := app.Fiber.Group("/api")

	v1 := api.Group("/v1")

	v1.Post("/auth/login", authController.Login).Name("auth.login")
	v1.Post("/auth/register", authController.Register).Name("auth.register")
	v1.Post("/auth/logout", authController.Logout).Name("auth.logout")
	v1.Post("/auth/refresh", authController.Refresh).Name("auth.refresh")
	v1.Get("/auth/user", authController.User).Name("auth.user")

	v1.Get("/users", userController.Index).Name("users.index")

	err := app.Fiber.Listen(":" + os.Getenv("APP_PORT"))

	if err != nil {
		panic(err)
	}
}
