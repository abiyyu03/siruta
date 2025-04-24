package http

import (
	"github.com/abiyyu03/siruta/usecase/user"
	"github.com/gofiber/fiber/v2"
)

type UserHttp struct{}

var userUsecase *user.UserUsecase

func (u *UserHttp) GetData(ctx *fiber.Ctx) error {
	return userUsecase.Fetch(ctx)
}

func (u *UserHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return userUsecase.FetchById(ctx, id)
}

func (u *UserHttp) RevokeUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return userUsecase.RevokeUserAccess(ctx, id)
}
