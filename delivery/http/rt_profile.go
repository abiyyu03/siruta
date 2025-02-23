package http

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/repository"
	"github.com/gofiber/fiber/v2"
)

type RTProfileHttp struct{}

var rtProfileRepository = new(repository.RTProfileRepository)

func (r *RTProfileHttp) GetData(ctx *fiber.Ctx) error {
	rtProfiles, err := rtProfileRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return entity.Success(ctx, &rtProfiles, "Data Fetched successfully")
}

func (r *RTProfileHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	rtProfile, err := rtProfileRepository.FetchById(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusNotFound, fiber.ErrNotFound.Message)
	}

	return entity.Success(ctx, &rtProfile, "Data Fetched successfully")
}
