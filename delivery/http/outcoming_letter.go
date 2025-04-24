package http

import (
	"github.com/abiyyu03/siruta/usecase/letter"
	"github.com/gofiber/fiber/v2"
)

type OutcomingLetterHttp struct{}

var outcomingLetterUsecase *letter.OutcomingLetterUsecase

func (o *OutcomingLetterHttp) GetData(ctx *fiber.Ctx) error {
	return outcomingLetterUsecase.Fetch(ctx)
}

func (o *OutcomingLetterHttp) GetPreview(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return outcomingLetterUsecase.FetchPreview(ctx, id)
}

func (o *OutcomingLetterHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return outcomingLetterUsecase.FetchById(ctx, id)
}

func (o *OutcomingLetterHttp) GetDataByRTProfileId(ctx *fiber.Ctx) error {
	rtProfileId := ctx.Params("rt_profile_id")

	return outcomingLetterUsecase.FetchByRTProfileId(ctx, rtProfileId)
}

func (o *OutcomingLetterHttp) DeleteData(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return outcomingLetterUsecase.Delete(ctx, id)
}
