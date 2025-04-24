package http

import (
	"github.com/abiyyu03/siruta/usecase/rt_profile"
	"github.com/gofiber/fiber/v2"
)

type RTProfileHttp struct{}

var rtProfileUsecase *rt_profile.RTProfileUsecase

func (r *RTProfileHttp) GetData(ctx *fiber.Ctx) error {
	return rtProfileUsecase.Fetch(ctx)
}

func (r *RTProfileHttp) GetDataByRWProfileId(ctx *fiber.Ctx) error {
	rwProfileId := ctx.Params("rw_profile_id")

	return rtProfileUsecase.FetchByRWProfileId(ctx, rwProfileId)
}

func (r *RTProfileHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return rtProfileUsecase.FetchById(ctx, id)
}
