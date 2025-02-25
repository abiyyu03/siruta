package http

import (
	"github.com/abiyyu03/siruta/usecase"
	"github.com/gofiber/fiber/v2"
)

type RWProfileHttp struct{}

var rwProfileUsecase = new(usecase.RWProfileUsecase)

func (r *RWProfileHttp) GetData(ctx *fiber.Ctx) error {
	return rwProfileUsecase.Fetch(ctx)
}

func (r *RWProfileHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return rwProfileUsecase.FetchById(ctx, id)
}

// func (r *RWProfileHttp) DeleteData(ctx *fiber.Ctx) error {}

// func (r *RWProfileHttp) UpdateData(ctx *fiber.Ctx) error {}
