package http

import (
	"github.com/abiyyu03/siruta/usecase"
	"github.com/gofiber/fiber/v2"
)

type LetterTypeHttp struct{}

var letterTypeUsecase = new(usecase.LetterTypeUsecase)

func (l *LetterTypeHttp) GetData(ctx *fiber.Ctx) error {
	return letterTypeUsecase.Fetch(ctx)
}

func (l *LetterTypeHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return letterTypeUsecase.FetchById(ctx, id)
}
