package usecase

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository"
	"github.com/gofiber/fiber/v2"
)

type RoleUsecase struct{}

var roleRepository = new(repository.RoleRepository)

func (v *RoleUsecase) Fetch(ctx *fiber.Ctx) error {
	roles, err := roleRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &roles, "Data fetched successfully")
}

func (v *RoleUsecase) FetchById(ctx *fiber.Ctx, id int) error {
	role, err := roleRepository.FetchById(id)

	if role == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &role, "Data fetched successfully")
}

func (v *RoleUsecase) Store(role *model.Role, ctx *fiber.Ctx) error {
	createdRole := &model.Role{
		Name: role.Name,
	}

	storedRole, err := roleRepository.Store(createdRole)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &storedRole, "Data updated successfully")
}

func (v *RoleUsecase) Update(role *model.Role, ctx *fiber.Ctx, id int) error {
	updateRole := &model.Role{
		Name: role.Name,
	}

	updatedRole, err := roleRepository.Update(updateRole, id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &updatedRole, "Data updated successfully")
}

func (v *RoleUsecase) Delete(ctx *fiber.Ctx, id int) error {
	_, err := roleRepository.Delete(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, nil, "Data deleted successfully")
}
