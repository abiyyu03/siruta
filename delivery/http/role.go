package http

import (
	"strconv"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase"
	"github.com/gofiber/fiber/v2"
)

type RoleHttp struct{}

var roleUsecase = new(usecase.RoleUsecase)

func (r *RoleHttp) GetData(ctx *fiber.Ctx) error {
	return roleUsecase.Fetch(ctx)
}
func (r *RoleHttp) GetDataById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return roleUsecase.FetchById(ctx, id)
}

func (r *RoleHttp) StoreData(ctx *fiber.Ctx) error {
	var role *model.Role

	if err := ctx.BodyParser(&role); err != nil {
		return err
	}

	return roleUsecase.Store(role, ctx)
}

func (r *RoleHttp) UpdateData(ctx *fiber.Ctx) error {
	var role *model.Role
	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := ctx.BodyParser(&role); err != nil {
		return err
	}

	return roleUsecase.Update(role, ctx, id)
}

func (r *RoleHttp) DeleteData(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return roleUsecase.Delete(ctx, id)
}
