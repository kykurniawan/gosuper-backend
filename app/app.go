package app

import (
	"context"
	"gosuper/app/exception"
	"gosuper/app/http/middlewares"
	"gosuper/app/libs/queue/consumers"
	"gosuper/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type App struct {
	Fiber    *fiber.App
	Database *gorm.DB
	Amqp     *amqp091.Connection
}

func NewApp(db *gorm.DB, amqp *amqp091.Connection) *App {
	fiber := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
		ErrorHandler:      exception.GlobalErrorHandler,
	})

	fiber.Use(recover.New())

	return &App{
		Fiber:    fiber,
		Database: db,
		Amqp:     amqp,
	}
}

func (app *App) Run() {
	app.runQueueWorker()
	app.registerRoutes(app.Fiber.Group("/api"))

	err := app.Fiber.Listen(":" + config.App.Port)

	if err != nil {
		log.Fatal(err)
	}
}

func (app *App) registerRoutes(api fiber.Router) {
	authService := InitializeAuthService(app.Database, app.Amqp)
	userService := InitializeUserService(app.Database)

	authController := InitializeAuthController(authService)
	userController := InitializeUserController(userService)

	v1 := api.Group("/v1").Name("v1.")

	v1.Post("/auth/login", authController.Login).Name("auth.login")
	v1.Post("/auth/register", authController.Register).Name("auth.register")
	v1.Post("/auth/logout", middlewares.Authenticate(authService), authController.Logout).Name("auth.logout")
	v1.Post("/auth/refresh", authController.Refresh).Name("auth.refresh")
	v1.Get("/auth/user", middlewares.Authenticate(authService), authController.User).Name("auth.user")
	v1.Post("/auth/forgot-password", authController.ForgotPassword).Name("auth.forgot-password")
	v1.Post("/auth/reset-password", authController.ResetPassword).Name("auth.reset-password")

	v1.Get("/users", middlewares.Authenticate(authService), userController.Index).Name("users.index")
}

func (app *App) runQueueWorker() {
	mailService := InitializeMailService()

	ctx := context.Background()
	q := InitializeQueue(app.Amqp)
	q.RegisterConsumer(consumers.NewSendEmailQueueConsumer(mailService))
	q.Run(ctx)
}
