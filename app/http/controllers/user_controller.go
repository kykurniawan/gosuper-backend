package controllers

import (
	"gosuper/app/helpers"
	"gosuper/app/http/responses"
	"gosuper/app/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (controller *UserController) Index(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Query param page is not valid!")
	}

	perPage, err := strconv.Atoi(c.Query("per_page", "10"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Query param per_page is not valid!")
	}

	users, total, totalPage, err := controller.userService.GetAllPaginate(page, perPage)

	if err != nil {
		return err
	}

	userResponse := make([]responses.UserResponse, len(users))

	for i, user := range users {
		userResponse[i] = responses.UserResponse{
			ID:              user.ID,
			Name:            user.Name,
			Email:           user.Email,
			EmailVerifiedAt: helpers.NilOrTIme(user.EmailVerifiedAt),
			CreatedAt:       helpers.NilOrTIme(user.CreatedAt),
			UpdatedAt:       helpers.NilOrTIme(user.UpdatedAt),
			DeletedAt:       helpers.NilOrTIme(user.DeletedAt),
		}
	}

	return c.JSON(fiber.Map{
		"message": "Get all users success!",
		"data": responses.UserIndexResponse{
			Users: userResponse,
			Meta: responses.PaginationMetaResponse{
				Total:     total,
				TotalPage: totalPage,
				Page:      page,
				PerPage:   perPage,
			},
		},
		"errors": nil,
	})
}
