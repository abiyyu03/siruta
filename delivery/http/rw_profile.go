package http

import (
	"github.com/gofiber/fiber/v2"
)

type RWProfileHttp struct{}

func (r *RWProfileHttp) GetData(ctx *fiber.Ctx) error {
	return rwProfileUsecase.Fetch(ctx)
}

func (r *RWProfileHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return rwProfileUsecase.FetchById(ctx, id)
}
