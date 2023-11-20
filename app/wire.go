//go:build wireinject
// +build wireinject

package app

import (
	"gosuper/app/http/controllers"
	"gosuper/app/libs/queue"
	"gosuper/app/repositories"
	"gosuper/app/services"

	"github.com/google/wire"
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

func InitializeApp(db *gorm.DB, amqp *amqp091.Connection) *App {
	panic(wire.Build(
		NewApp,
	))
}

func InitializeUserRepository(db *gorm.DB) *repositories.UserRepository {
	panic(wire.Build(
		repositories.NewUserRepository,
	))
}

func InitializeRefreshTokenRepository(db *gorm.DB) *repositories.RefreshTokenRepository {
	panic(wire.Build(
		repositories.NewRefreshTokenRepository,
	))
}

func InitializeUserService(db *gorm.DB) *services.UserService {
	panic(wire.Build(
		services.NewUserService,
		InitializeUserRepository,
	))
}

func InitializeAuthService(db *gorm.DB, amqp *amqp091.Connection) *services.AuthService {
	panic(wire.Build(
		services.NewAuthService,
		InitializeUserService,
		InitializeOtpService,
		InitializeMailService,
		InitializeRefreshTokenRepository,
		InitializeQueue,
	))
}

func InitializeAuthController(authService *services.AuthService) *controllers.AuthController {
	panic(wire.Build(
		controllers.NewAuthController,
	))
}

func InitializeUserController(userService *services.UserService) *controllers.UserController {
	panic(wire.Build(
		controllers.NewUserController,
	))
}

func InitializeOtpRepository(db *gorm.DB) *repositories.OtpRepository {
	panic(wire.Build(
		repositories.NewOtpRepository,
	))
}

func InitializeOtpService(db *gorm.DB) *services.OtpService {
	panic(wire.Build(
		services.NewOtpService,
		InitializeOtpRepository,
	))
}

func InitializeMailService() *services.MailService {
	panic(wire.Build(
		services.NewMailService,
	))
}

func InitializeQueue(amqp *amqp091.Connection) *queue.Queue {
	panic(wire.Build(
		queue.NewQueue,
	))
}
