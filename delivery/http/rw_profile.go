package http

import (
	"github.com/abiyyu03/siruta/usecase/rw_profile"
	"github.com/gofiber/fiber/v2"
)

type RWProfileHttp struct {
	rwProfileUsecase rw_profile.RWProfileUsecaseInterface
}

func NewRWProfileHttp(rwProfileUsecase rw_profile.RWProfileUsecaseInterface) *RWProfileHttp {
	return &RWProfileHttp{
		rwProfileUsecase: rwProfileUsecase,
	}
}

func (r *RWProfileHttp) GetData(ctx *fiber.Ctx) error {
	return r.rwProfileUsecase.Fetch(ctx)
}

func (r *RWProfileHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return r.rwProfileUsecase.FetchById(ctx, id)
}
