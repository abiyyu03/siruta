package http

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/repository"
	"github.com/gofiber/fiber/v2"
)

type RWProfileHttp struct{}

var rwProfileRepository = new(repository.RWProfileRepository)

func (r *RWProfileHttp) GetData(ctx *fiber.Ctx) error {
	rwProfiles, err := rwProfileRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, "Server error")
	}

	return entity.Success(ctx, &rwProfiles, "Data fetched successfully")
}

func (r *RWProfileHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	rwProfile, err := rwProfileRepository.FetchById(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusNotFound, "RW Profile not found")
	}

	return entity.Success(ctx, &rwProfile, "Data fetched successfully")
}

// func (r *RWProfileHttp) DeleteData(ctx *fiber.Ctx) error {}

// func (r *RWProfileHttp) UpdateData(ctx *fiber.Ctx) error {}
