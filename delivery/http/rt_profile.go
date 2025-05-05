package http

import (
	"github.com/abiyyu03/siruta/usecase/rt_profile"
	"github.com/gofiber/fiber/v2"
)

type RTProfileHttp struct {
	rtProfileUsecase rt_profile.RTProfileUsecaseInterface
}

func NewRTProfileHttp(rtProfileUsecase rt_profile.RTProfileUsecaseInterface) *RTProfileHttp {
	return &RTProfileHttp{
		rtProfileUsecase: rtProfileUsecase,
	}
}

func (r *RTProfileHttp) GetData(ctx *fiber.Ctx) error {
	return r.rtProfileUsecase.Fetch(ctx)
}

func (r *RTProfileHttp) GetDataByRWProfileId(ctx *fiber.Ctx) error {
	rwProfileId := ctx.Params("rw_profile_id")

	return r.rtProfileUsecase.FetchByRWProfileId(ctx, rwProfileId)
}

func (r *RTProfileHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return r.rtProfileUsecase.FetchById(ctx, id)
}
