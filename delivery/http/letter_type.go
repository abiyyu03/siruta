package http

import (
	"strconv"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/gofiber/fiber/v2"
)

type LetterTypeHttp struct{}

func (v *LetterTypeHttp) GetData(ctx *fiber.Ctx) error {
	return letterTypeUsecase.Fetch(ctx)
}

func (v *LetterTypeHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return letterTypeUsecase.FetchById(ctx, id)
}

func (v *LetterTypeHttp) StoreData(ctx *fiber.Ctx) error {
	var letterType *model.LetterType

	if err := ctx.BodyParser(&letterType); err != nil {
		return err
	}

	return letterTypeUsecase.Store(letterType, ctx)
}

func (v *LetterTypeHttp) UpdateData(ctx *fiber.Ctx) error {
	var letterType *model.LetterType
	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := ctx.BodyParser(&letterType); err != nil {
		return err
	}

	return letterTypeUsecase.Update(letterType, ctx, id)
}

func (v *LetterTypeHttp) DeleteData(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return letterTypeUsecase.Delete(ctx, id)
}
