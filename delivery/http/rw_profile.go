package http

import (
	"github.com/abiyyu03/siruta/usecase/rw_profile"
	"github.com/gofiber/fiber/v2"
)

type RWProfileHttp struct{}

var rwProfileUsecase *rw_profile.RWProfileUsecase

func (r *RWProfileHttp) GetData(ctx *fiber.Ctx) error {
	return rwProfileUsecase.Fetch(ctx)
}

func (r *RWProfileHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return rwProfileUsecase.FetchById(ctx, id)
}
