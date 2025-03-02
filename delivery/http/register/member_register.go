package register

import (
	"github.com/abiyyu03/siruta/entity/request"
	"github.com/abiyyu03/siruta/usecase/register"
	"github.com/gofiber/fiber/v2"
)

type MemberRegisterHttp struct {
	memberRegisterUsecase *register.MemberRegisterUsecase
}

func (m *MemberRegisterHttp) Register(ctx *fiber.Ctx) error {
	var userMember *request.RegisterRequest

	if err := ctx.BodyParser(&userMember); err != nil {
		return err
	}

	return m.memberRegisterUsecase.RegisterMember(ctx, userMember)
}
