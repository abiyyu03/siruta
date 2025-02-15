package http

import (
	"github.com/abiyyu03/siruta/usecase"
	"github.com/gofiber/fiber/v2"
)

type UserHttp struct{}

var UserUsecase = new(usecase.UserUsecase)

func (u *UserHttp) GetData(ctx *fiber.Ctx) error {
	return UserUsecase.Fetch(ctx)
}

func (u *UserHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return UserUsecase.FetchById(ctx, id)
}
