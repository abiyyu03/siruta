package http

import (
	"github.com/abiyyu03/siruta/usecase/rt_profile"
	"github.com/gofiber/fiber/v2"
)

type RTProfileHttp struct {
	rtProfileUsecase *rt_profile.RTProfileUsecase
}

func (r *RTProfileHttp) GetData(ctx *fiber.Ctx) error {
	return r.rtProfileUsecase.Fetch(ctx)
}

func (r *RTProfileHttp) GetDataByRtProfileId(ctx *fiber.Ctx) error {
	rtProfileId := ctx.Params("rt_profile_id")

	return r.rtProfileUsecase.FetchByRtProfileId(ctx, rtProfileId)
}

func (r *RTProfileHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return r.rtProfileUsecase.FetchById(ctx, id)
}
