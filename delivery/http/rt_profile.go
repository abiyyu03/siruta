package http

import (
	"github.com/gofiber/fiber/v2"
)

type RTProfileHttp struct{}

func (r *RTProfileHttp) GetData(ctx *fiber.Ctx) error {
	return rtProfileUsecase.Fetch(ctx)
}

func (r *RTProfileHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return rtProfileUsecase.FetchById(ctx, id)
}
