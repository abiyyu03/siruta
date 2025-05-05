package routes

import (
	"github.com/abiyyu03/siruta/delivery/http/register"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/entity/request"
	"github.com/gofiber/fiber/v2"
)

func RegisterRegisterRoutes(v1 fiber.Router, handlerRT *register.RTProfileRegisterHttp, handlerRW *register.RWProfileRegisterHttp, handlerMember *register.MemberRegisterHttp) {
	v1.Post("/registers/rw", middleware.ValidateField[model.RWProfile](), handlerRW.RegisterRWProfile)
	v1.Put("/registers/rw/:rwProfileId/approve", handlerRW.ApproveRegistration)
	v1.Post("/registers/rw/user-account", handlerRW.RegisterUserRw)
	v1.Post("/registers/rt", middleware.ValidateField[request.RTProfileRegisterRequest](), handlerRT.RegisterRTProfile)
	v1.Put("/registers/rt/:rtProfileId/approve", handlerRT.ApproveRegistration)
	v1.Post("/registers/rt/user-account", handlerRT.RegisterUserRt)
	v1.Post("/registers/member", middleware.ValidateField[request.MemberRegisterRequest](), handlerMember.Register)
}
