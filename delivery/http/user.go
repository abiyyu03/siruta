package http

import (
	"github.com/abiyyu03/siruta/usecase/user"
	"github.com/gofiber/fiber/v2"
)

type UserHttp struct {
	userUsecase *user.UserUsecase
}

func (u *UserHttp) GetData(ctx *fiber.Ctx) error {
	return u.userUsecase.Fetch(ctx)
}

func (u *UserHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return u.userUsecase.FetchById(ctx, id)
}
