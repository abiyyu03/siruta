package http

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase"
	"github.com/gofiber/fiber/v2"
)

type MemberHttp struct{}

var memberUsecase = new(usecase.MemberUsecase)

func (m *MemberHttp) GetData(ctx *fiber.Ctx) error {
	return memberUsecase.Fetch(ctx)
}

func (m *MemberHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return memberUsecase.FetchById(ctx, id)
}

func (m *MemberHttp) StoreData(ctx *fiber.Ctx) error {
	var member *model.Member

	if err := ctx.BodyParser(&member); err != nil {
		return entity.Error(ctx, fiber.ErrBadRequest.Code, err.Error())
	}

	return memberUsecase.Store(ctx, member)
}

func (m *MemberHttp) UpdateData(ctx *fiber.Ctx) error {
	var member *model.Member
	id := string(ctx.Params("id"))

	if err := ctx.BodyParser(&member); err != nil {
		return entity.Error(ctx, fiber.ErrBadRequest.Code, err.Error())
	}

	return memberUsecase.Update(ctx, id, member)
}
