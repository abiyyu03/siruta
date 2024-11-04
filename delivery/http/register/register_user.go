package register

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository"
	"github.com/gofiber/fiber/v2"
)

func Register(ctx *fiber.Ctx) error {
	var user *model.User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	userRepository := new(repository.UserRepository)
	registeredUser, err := userRepository.RegisterUser(user, user.RoleID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return entity.Success(ctx, &registeredUser, "Register user successfully")

}
