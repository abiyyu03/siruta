package http

import (
	"strconv"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase/letter_type"
	"github.com/gofiber/fiber/v2"
)

type LetterTypeHttp struct{}

var letterTypeUsecase *letter_type.LetterTypeUsecase

func (l *LetterTypeHttp) GetData(ctx *fiber.Ctx) error {
	return letterTypeUsecase.Fetch(ctx)
}

func (l *LetterTypeHttp) GetDataById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return letterTypeUsecase.FetchById(ctx, id)
}

func (l *LetterTypeHttp) StoreData(ctx *fiber.Ctx) error {
	var letterType *model.LetterType

	if err := ctx.BodyParser(&letterType); err != nil {
		return err
	}

	return letterTypeUsecase.Store(letterType, ctx)
}

func (l *LetterTypeHttp) UpdateData(ctx *fiber.Ctx) error {
	var letterType *model.LetterType
	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := ctx.BodyParser(&letterType); err != nil {
		return err
	}

	return letterTypeUsecase.Update(letterType, ctx, id)
}

func (l *LetterTypeHttp) DeleteData(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return letterTypeUsecase.Delete(ctx, id)
}
