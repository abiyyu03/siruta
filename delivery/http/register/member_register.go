package register

import (
	"github.com/abiyyu03/siruta/entity/request"
	"github.com/abiyyu03/siruta/usecase"
	"github.com/gofiber/fiber/v2"
)

type MemberRegisterHttp struct{}

var memberRegisterUsecase = new(usecase.MemberRegisterUsecase)

func (m *MemberRegisterHttp) Register(ctx *fiber.Ctx) error {
	var userMember *request.RegisterRequest

	if err := ctx.BodyParser(&userMember); err != nil {
		return err
	}

	return memberRegisterUsecase.RegisterMember(ctx, userMember)
}
