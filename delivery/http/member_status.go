package http

import (
	"strconv"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/gofiber/fiber/v2"
)

type MemberStatusHttp struct{}

func (m *MemberStatusHttp) GetData(ctx *fiber.Ctx) error {
	return memberStatusUsecase.Fetch(ctx)
}

func (m *MemberStatusHttp) GetDataById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return memberStatusUsecase.FetchById(ctx, id)
}

func (m *MemberStatusHttp) StoreData(ctx *fiber.Ctx) error {
	var memberStatus *model.MemberStatus

	if err := ctx.BodyParser(&memberStatus); err != nil {
		return err
	}

	return memberStatusUsecase.Store(memberStatus, ctx)
}

func (m *MemberStatusHttp) UpdateData(ctx *fiber.Ctx) error {
	var memberStatus *model.MemberStatus
	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := ctx.BodyParser(&memberStatus); err != nil {
		return err
	}

	return memberStatusUsecase.Update(memberStatus, ctx, id)
}

func (m *MemberStatusHttp) DeleteData(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return memberStatusUsecase.Delete(ctx, id)
}
