package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/http/auth"
	"github.com/abiyyu03/siruta/delivery/http/register"
	"github.com/gofiber/fiber/v2"

	_ "github.com/swaggo/fiber-swagger/example/docs"
)

type HandlerDefinition struct {
	Member            *http.MemberHttp
	MemberRegister    *register.MemberRegisterHttp
	RWProfileRegister *register.RWProfileRegisterHttp
	RTProfileRegister *register.RTProfileRegisterHttp
	Village           *http.VillageHttp
	LetterType        *http.LetterTypeHttp
	RWProfile         *http.RWProfileHttp
	RTProfile         *http.RTProfileHttp
	Role              *http.RoleHttp
	IncomingLetter    *http.IncomingLetterHttp
	LetterReq         *http.LetterReqHttp
	OutcomingLetter   *http.OutcomingLetterHttp
	MemberStatus      *http.MemberStatusHttp
	User              *http.UserHttp
	Religion          *http.ReligionHttp
	ReferalCode       *http.ReferalCodeHttp
	RTLeader          *http.RTLeaderHttp
	RWLeader          *http.RWLeaderHttp
	Inventory         *http.InventoryHttp
	GuestList         *http.GuestListHttp
	Cashflow          *http.CashflowHttp
	ResetPassword     *auth.ResetPasswordHttp
	Auth              *auth.AuthHttp
}

func (handler *HandlerDefinition) HttpRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	RegisterRoleRoutes(v1, handler.Role)
	RegisterCashflowRoutes(v1, handler.Cashflow)
	RegisterAuthRoutes(v1, handler.Auth)
	RegisterVillageRoutes(v1, handler.Village)
	RegisterLetterTypeRoutes(v1, handler.LetterType)
	RegisterRWProfileRoutes(v1, handler.RWProfile)
	RegisterReligionRoutes(v1, handler.Religion)
	RegisterMemberStatusRoutes(v1, handler.MemberStatus)
	RegisterMemberRoutes(v1, handler.Member)
	RegisterInventoryRoutes(v1, handler.Inventory)
	RegisterReferalCodeRoutes(v1, handler.ReferalCode)
	RegisterGuestListRoutes(v1, handler.GuestList)
	RegisterOutcomingLetterRoutes(v1, handler.OutcomingLetter)
	RegisterUserRoutes(v1, handler.User)
	RegisterResetPasswordRoutes(v1, handler.ResetPassword)
	RegisterLetterReqRoutes(v1, handler.LetterReq)
	RegisterRWProfileRoutes(v1, handler.RWProfile)
	RegisterRTProfileRoutes(v1, handler.RTProfile)
	RegisterRWLeaderRoutes(v1, handler.RWLeader)
	RegisterRTLeaderRoutes(v1, handler.RTLeader)
	RegisterRegisterRoutes(v1, handler.RTProfileRegister, handler.RWProfileRegister, handler.MemberRegister)
}
