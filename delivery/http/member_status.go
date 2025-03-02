package http

import (
	"strconv"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase/member_status"
	"github.com/gofiber/fiber/v2"
)

type MemberStatusHttp struct {
	memberStatusUsecase *member_status.MemberStatusUsecase
}

func (m *MemberStatusHttp) GetData(ctx *fiber.Ctx) error {
	return m.memberStatusUsecase.Fetch(ctx)
}

func (m *MemberStatusHttp) GetDataById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return m.memberStatusUsecase.FetchById(ctx, id)
}

func (m *MemberStatusHttp) StoreData(ctx *fiber.Ctx) error {
	var memberStatus *model.MemberStatus

	if err := ctx.BodyParser(&memberStatus); err != nil {
		return err
	}

	return m.memberStatusUsecase.Store(memberStatus, ctx)
}

func (m *MemberStatusHttp) UpdateData(ctx *fiber.Ctx) error {
	var memberStatus *model.MemberStatus
	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := ctx.BodyParser(&memberStatus); err != nil {
		return err
	}

	return m.memberStatusUsecase.Update(memberStatus, ctx, id)
}

func (m *MemberStatusHttp) DeleteData(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return m.memberStatusUsecase.Delete(ctx, id)
}
