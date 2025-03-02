package http

import (
	"strconv"

	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase/letter"
	"github.com/gofiber/fiber/v2"
)

type IncomingLetterHttp struct {
	incomingLetterUsecase *letter.IncomingLetterUsecase
}

func (i *IncomingLetterHttp) GetData(ctx *fiber.Ctx) error {
	return i.incomingLetterUsecase.Fetch(ctx)
}

func (i *IncomingLetterHttp) GetDataById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return i.incomingLetterUsecase.FetchById(ctx, id)
}

func (i *IncomingLetterHttp) UpdateData(ctx *fiber.Ctx) error {
	var incomingLetter *model.IncomingLetter
	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := ctx.BodyParser(&incomingLetter); err != nil {
		return entity.Error(ctx, fiber.StatusUnprocessableEntity, constant.Errors["UnprocessableEntity"].Message, constant.Errors["UnprocessableEntity"].Clue)
	}

	return i.incomingLetterUsecase.Update(incomingLetter, ctx, id)
}

func (i *IncomingLetterHttp) StoreData(ctx *fiber.Ctx) error {
	var incomingLetter *model.IncomingLetter

	if err := ctx.BodyParser(&incomingLetter); err != nil {
		return entity.Error(ctx, fiber.StatusUnprocessableEntity, constant.Errors["UnprocessableEntity"].Message, constant.Errors["UnprocessableEntity"].Clue)
	}

	return i.incomingLetterUsecase.Store(incomingLetter, ctx)
}

func (i *IncomingLetterHttp) DeleteData(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return i.incomingLetterUsecase.Delete(ctx, id)
}
