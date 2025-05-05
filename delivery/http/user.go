package http

import (
	"github.com/abiyyu03/siruta/usecase/user"
	"github.com/gofiber/fiber/v2"
)

type UserHttp struct {
	userUsecase user.UserUsecaseInterface
}

func NewUserHttp(userUC user.UserUsecaseInterface) *UserHttp {
	return &UserHttp{
		userUsecase: userUC,
	}
}

func (u *UserHttp) GetData(ctx *fiber.Ctx) error {
	return u.userUsecase.Fetch(ctx)
}

func (u *UserHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return u.userUsecase.FetchById(ctx, id)
}

func (u *UserHttp) RevokeUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return u.userUsecase.RevokeUserAccess(ctx, id)
}
