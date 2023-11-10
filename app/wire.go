//go:build wireinject
// +build wireinject

package app

import (
	"gosuper/app/http/controllers"
	"gosuper/app/repositories"
	"gosuper/app/services"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeApp(db *gorm.DB) *App {
	panic(wire.Build(
		NewApp,
	))
}

func InitializeUserRepository(db *gorm.DB) *repositories.UserRepository {
	panic(wire.Build(
		repositories.NewUserRepository,
	))
}

func InitializeUserService(db *gorm.DB) *services.UserService {
	panic(wire.Build(
		services.NewUserService,
		InitializeUserRepository,
	))
}

func InitializeAuthService(db *gorm.DB) *services.AuthService {
	panic(wire.Build(
		services.NewAuthService,
		InitializeUserService,
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
