package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/http/auth"
	"github.com/abiyyu03/siruta/delivery/http/register"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/entity/request"
	"github.com/gofiber/fiber/v2"
)

type HandlerDefinition struct {
	member            *http.MemberHttp
	memberRegister    *register.MemberRegisterHttp
	rwProfileRegister *register.RWProfileRegisterHttp
	rtProfileRegister *register.RTProfileRegisterHttp
	village           *http.VillageHttp
	letterType        *http.LetterTypeHttp
	rwProfile         *http.RWProfileHttp
	rtProfile         *http.RTProfileHttp
	role              *http.RoleHttp
	incomingLetter    *http.IncomingLetterHttp
	letterReq         *http.LetterReqHttp
	OutcomingLetter   *http.OutcomingLetterHttp
	memberStatus      *http.MemberStatusHttp
	user              *http.UserHttp
	religion          *http.ReligionHttp
	referalCode       *http.ReferalCodeHttp
	rtLeader          *http.RTLeaderHttp
	rwLeader          *http.RWLeaderHttp
	inventory         *http.InventoryHttp
	guestList         *http.GuestListHttp
	cashflow          *http.CashflowHttp
	resetPassword     *auth.ResetPasswordHttp
}

func (handler *HandlerDefinition) HttpRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	adminOnly := middleware.JWTMiddleware([]int{1})
	rwLeaderOnly := middleware.JWTMiddleware([]int{2})
	rtLeaderOnly := middleware.JWTMiddleware([]int{3})
	memberOnly := middleware.JWTMiddleware([]int{4})
	// validator := middleware.ValidateField()

	//authentication
	v1.Post("/login", middleware.ValidateField[request.LoginRequest](), auth.Login)
	v1.Post("/forgot-password", handler.resetPassword.SendForgotPasswordLink)
	v1.Put("/reset-password", middleware.ValidateField[request.ResetPassword](), handler.resetPassword.ResetPassword)

	// members
	v1.Get("/members", adminOnly, handler.member.GetData)
	v1.Get("/members/:id", adminOnly, handler.member.GetDataById)
	v1.Put("/members/:id", adminOnly, handler.member.UpdateData)
	v1.Get("/members/:rt_profile_id/rt", adminOnly, rtLeaderOnly, handler.member.GetDataByRTProfileId)

	// register
	v1.Post("/registers/rw", middleware.ValidateField[model.RWProfile](), handler.rwProfileRegister.RegisterRWProfile)
	v1.Put("/registers/rw/:rwProfileId/approve", handler.rwProfileRegister.ApproveRegistration)
	v1.Post("/registers/rw/user-account", handler.rwProfileRegister.RegisterUserRw)
	v1.Post("/registers/rt", middleware.ValidateField[request.RTProfileRegisterRequest](), handler.rtProfileRegister.RegisterRTProfile)
	v1.Put("/registers/rt/:rtProfileId/approve", handler.rtProfileRegister.ApproveRegistration)
	v1.Post("/registers/rt/user-account", handler.rtProfileRegister.RegisterUserRt)
	v1.Post("/registers/member", middleware.ValidateField[request.MemberRegisterRequest](), handler.memberRegister.Register)

	//inventory
	v1.Get("/inventories", rtLeaderOnly, handler.inventory.GetData)
	v1.Get("/inventories/:id", rtLeaderOnly, handler.inventory.GetDataById)
	v1.Post("/inventories", rtLeaderOnly, handler.inventory.StoreData)
	v1.Put("/inventories/:id", rtLeaderOnly, handler.inventory.UpdateData)
	v1.Delete("/inventories/:id", rtLeaderOnly, handler.inventory.DeleteData)

	//referal codes
	v1.Get("/referal-codes/:profile_id/rt", rtLeaderOnly, handler.referalCode.GetDataByRTProfileId)
	v1.Put("/referal-codes/:profile_id/rt/regenerate/:code", rtLeaderOnly, handler.referalCode.RegenerateCode)

	// incoming letter
	v1.Get("/incoming-letters/:rt_profile_id/rt", rtLeaderOnly, handler.incomingLetter.GetDataByRTProfileId)

	// guest list
	v1.Get("/guest-lists", adminOnly, handler.guestList.GetData)
	v1.Get("/guest-lists/:rt_profile_id/rt", adminOnly, rtLeaderOnly, handler.guestList.GetDataByRTProfileId)
	v1.Get("/guest-lists/:id", adminOnly, rtLeaderOnly, handler.guestList.GetDataById)
	v1.Put("/guest-lists/:id", adminOnly, rtLeaderOnly, handler.guestList.UpdateData)
	v1.Delete("/guest-lists/:id", adminOnly, rtLeaderOnly, handler.guestList.DeleteData)

	// Request letter
	v1.Put("/request-letters/approve/:letter_req_id", adminOnly, rtLeaderOnly, handler.letterReq.UpdateApprovalStatus)

	// cashflow
	v1.Get("/finances/cashflow", adminOnly, handler.cashflow.GetData)
	v1.Get("/finances/cashflow/:rt_profile_id/rt", adminOnly, rtLeaderOnly, handler.cashflow.GetDataByRTProfileId)
	v1.Get("/finances/cashflow/:id", adminOnly, rtLeaderOnly, handler.cashflow.GetDataById)
	v1.Put("/finances/cashflow/:id", adminOnly, rtLeaderOnly, handler.cashflow.UpdateData)
	v1.Post("/finances/cashflow", handler.cashflow.StoreData)
	v1.Delete("/finances/cashflow/:id", adminOnly, rtLeaderOnly, handler.cashflow.DeleteData)
	//----------------------------------------------------------------
	//
	// Admin Authority
	//
	//----------------------------------------------------------------
	// members
	v1.Get("/members", adminOnly, handler.member.GetData)
	v1.Get("/members/:id", adminOnly, handler.member.GetDataById)
	v1.Put("/members/:id", adminOnly, handler.member.UpdateData)

	//rw profiles
	v1.Get("/rw-profiles", adminOnly, handler.rwProfile.GetData)
	v1.Get("/rw-profiles/:id", adminOnly, rtLeaderOnly, handler.rwProfile.GetDataById)

	//rt profiles
	v1.Get("/rt-profiles", adminOnly, handler.rtProfile.GetData)
	v1.Get("/rt-profiles/:id", adminOnly, rwLeaderOnly, handler.rtProfile.GetDataById)
	v1.Get("/rt-profiles/:rw_profile_id", adminOnly, rwLeaderOnly, handler.rtProfile.GetDataByRWProfileId)

	//referal code
	v1.Get("/referal-codes", adminOnly, handler.referalCode.GetData)
	v1.Get("/referal-codes/:id", adminOnly, handler.referalCode.GetDataById)
	v1.Get("/referal-codes/:rt_profile_id", adminOnly, handler.referalCode.GetDataByRTProfileId)
	v1.Post("/referal-codes/validate", handler.referalCode.ValidateReferalCode)
	v1.Delete("/referal-codes/:id", adminOnly, handler.referalCode.GetDataById)

	//users
	v1.Get("/users", adminOnly, handler.user.GetData)
	v1.Get("/users/:id", adminOnly, handler.user.GetDataById)

	//leaders
	v1.Get("/rt-leaders", adminOnly, handler.rtLeader.GetData)
	v1.Get("/rt-leaders/:id", adminOnly, rtLeaderOnly, handler.rtLeader.GetDataById)
	v1.Put("/rt-leaders/:id", adminOnly, rtLeaderOnly, handler.rtLeader.UpdateData)
	v1.Get("/rw-leaders", adminOnly, handler.rwLeader.GetData)
	v1.Get("/rw-leaders/:id", adminOnly, rwLeaderOnly, handler.rwLeader.GetDataById)
	v1.Put("/rw-leaders/:id", adminOnly, rwLeaderOnly, handler.rwLeader.UpdateData)

	//roles
	v1.Get("/roles", adminOnly, handler.role.GetData)
	v1.Get("/roles/:id", adminOnly, handler.role.GetDataById)
	v1.Post("/roles", adminOnly, handler.role.StoreData)
	v1.Put("/roles/:id", adminOnly, handler.role.UpdateData)
	v1.Delete("/roles/:id", adminOnly, handler.role.DeleteData)

	//incoming letter
	v1.Get("/incoming-letters", adminOnly, handler.incomingLetter.GetData)
	v1.Get("/incoming-letters/:id", adminOnly, handler.incomingLetter.GetDataById)
	v1.Post("/incoming-letters", adminOnly, handler.incomingLetter.StoreData)
	v1.Put("/incoming-letters/:id", adminOnly, handler.incomingLetter.UpdateData)
	v1.Delete("/incoming-letters/:id", adminOnly, handler.incomingLetter.DeleteData)

	//member status
	v1.Get("/member-status", adminOnly, handler.memberStatus.GetData)
	v1.Get("/member-status/:id", adminOnly, handler.memberStatus.GetDataById)
	v1.Post("/member-status", adminOnly, handler.memberStatus.StoreData)
	v1.Put("/member-status/:id", adminOnly, handler.memberStatus.UpdateData)
	v1.Delete("/member-status/:id", adminOnly, handler.memberStatus.DeleteData)

	//village
	v1.Get("/villages", adminOnly, rtLeaderOnly, memberOnly, handler.village.GetData)
	v1.Get("/villages/:id", adminOnly, handler.village.GetDataById)
	v1.Post("/villages", adminOnly, handler.village.StoreData)
	v1.Put("/villages/:id", adminOnly, handler.village.UpdateData)
	v1.Delete("/villages/:id", adminOnly, handler.village.DeleteData)

	//religion
	v1.Get("/religions", adminOnly, rtLeaderOnly, memberOnly, handler.religion.GetData)
	v1.Get("/religions/:id", adminOnly, handler.religion.GetDataById)
	v1.Post("/religions", adminOnly, handler.religion.StoreData)
	v1.Put("/religions/:id", adminOnly, handler.religion.UpdateData)
	v1.Delete("/religions/:id", adminOnly, handler.religion.DeleteData)

	//letter type
	v1.Get("/letter-types", adminOnly, rtLeaderOnly, memberOnly, handler.letterType.GetData)
	v1.Get("/letter-types/:id", adminOnly, handler.letterType.GetDataById)
	v1.Post("/letter-types", adminOnly, handler.letterType.StoreData)
	v1.Put("/letter-types/:id", adminOnly, handler.letterType.UpdateData)
	v1.Delete("/letter-types/:id", adminOnly, handler.letterType.DeleteData)

	//outcoming letter
	v1.Get("/outcoming-letters", adminOnly, handler.OutcomingLetter.GetData)
	v1.Get("/outcoming-letters/:rt_profile_id", adminOnly, rtLeaderOnly, handler.OutcomingLetter.GetDataByRTProfileId)
	// 	v1.Get("/outcoming-letters/:id", adminOnly, handler.OutcomingLetter.GetDataById)

	//letter req
	v1.Post("/request-letters", rtLeaderOnly, memberOnly, handler.letterReq.CreateData)
	v1.Put("/request-letters/approve/:letter_req_id", rtLeaderOnly, handler.letterReq.UpdateApprovalStatus)

	v1.Get("/outcoming-letters/:id", adminOnly, rtLeaderOnly, handler.OutcomingLetter.GetDataById)
	v1.Get("/outcoming-letters/:id/preview", handler.OutcomingLetter.GetPreview)
	v1.Get("/outcoming-letters/:rt_profile_id/rt", adminOnly, rtLeaderOnly, handler.OutcomingLetter.GetDataByRTProfileId)

	//letter req
	v1.Post("/request-letters", adminOnly, memberOnly, handler.letterReq.CreateData)
}
