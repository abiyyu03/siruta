package http

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase/member"
	"github.com/gofiber/fiber/v2"
)

type MemberHttp struct{}

var memberUsecase *member.MemberUsecase

func (m *MemberHttp) GetData(ctx *fiber.Ctx) error {
	return memberUsecase.Fetch(ctx)
}

func (m *MemberHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return memberUsecase.FetchById(ctx, id)
}

func (m *MemberHttp) GetDataByRTProfileId(ctx *fiber.Ctx) error {
	rtProfileId := ctx.Params("rt_profile_id")

	return memberUsecase.FetchByRTProfileId(ctx, rtProfileId)
}

func (m *MemberHttp) UpdateData(ctx *fiber.Ctx) error {
	var member *model.Member
	id := string(ctx.Params("id"))

	if err := ctx.BodyParser(&member); err != nil {
		return entity.Error(ctx, fiber.StatusUnprocessableEntity, constant.Errors["UnprocessableEntity"].Message, constant.Errors["UnprocessableEntity"].Clue)
	}

	return memberUsecase.Update(ctx, id, member)
}
