package http

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository"
	"github.com/gofiber/fiber/v2"
)

type RoleHttp struct{}

var RoleRepository = new(repository.RolesRepository)

func (r *RoleHttp) GetData(ctx *fiber.Ctx) error {
	var roles []*model.Role

	roles, err := RoleRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, "Server error")
	}

	return entity.Success(ctx, &roles, "Data fetched successfully")

}
