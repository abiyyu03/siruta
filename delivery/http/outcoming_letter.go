package http

import (
	"github.com/abiyyu03/siruta/usecase/letter"
	"github.com/gofiber/fiber/v2"
)

type OutcomingLetterHttp struct {
	outcomingLetterUsecase *letter.OutcomingLetterUsecase
}

func (o *OutcomingLetterHttp) GetData(ctx *fiber.Ctx) error {
	return o.outcomingLetterUsecase.Fetch(ctx)
}

func (o *OutcomingLetterHttp) GetDataByRtProfileId(ctx *fiber.Ctx) error {
	rtProfileId := ctx.Params("rt_profile_id")

	return o.outcomingLetterUsecase.FetchByRtProfileId(ctx, rtProfileId)
}
