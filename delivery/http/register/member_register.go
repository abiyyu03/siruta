package register

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/request"
	"github.com/abiyyu03/siruta/usecase/referal_code"
	"github.com/abiyyu03/siruta/usecase/register"
	"github.com/gofiber/fiber/v2"
)

type MemberRegisterHttp struct{}

var memberRegisterUsecase *register.MemberRegisterUsecase
var referalCodeUsecase *referal_code.ReferalCodeUsecase

func (m *MemberRegisterHttp) Register(ctx *fiber.Ctx) error {
	params := ctx.Queries()
	//token verif
	_, profileId := referalCodeUsecase.Validate(ctx, params["referal_code"])

	if profileId == "" {
		return entity.Error(ctx, fiber.StatusForbidden, constant.Errors["InvalidReferalCode"].Message, constant.Errors["InvalidReferalCode"].Clue)
	}
	var userMember *request.MemberRegisterRequest

	if err := ctx.BodyParser(&userMember); err != nil {
		return err
	}

	return memberRegisterUsecase.RegisterMember(ctx, userMember, profileId)
}
