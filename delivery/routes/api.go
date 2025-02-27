package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/http/auth"
	"github.com/abiyyu03/siruta/delivery/http/middleware"
	"github.com/abiyyu03/siruta/delivery/http/register"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/entity/request"
	"github.com/gofiber/fiber/v2"
)

func HttpRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	member := new(http.MemberHttp)
	rwProfileRegister := new(register.RWProfileRegisterHttp)
	rtProfileRegister := new(register.RTProfileRegisterHttp)
	village := new(http.VillageHttp)
	letterType := new(http.LetterTypeHttp)
	rwProfile := new(http.RWProfileHttp)
	rtProfile := new(http.RTProfileHttp)
	role := new(http.RoleHttp)
	incomingLetter := new(http.IncomingLetterHttp)
	letterReq := new(http.LetterReqHttp)
	OutcomingLetter := new(http.OutcomingLetterHttp)
	memberStatus := new(http.MemberStatusHttp)
	user := new(http.UserHttp)
	religion := new(http.ReligionHttp)
	referalCode := new(http.ReferalCodeHttp)

	adminOnly := middleware.JWTMiddleware([]int{1})
	// rwLeaderOnly := middleware.JWTMiddleware([]int{2})
	// rwAdministratorOnly := middleware.JWTMiddleware([]int{3})
	// validator := middleware.ValidateField()

	//authentication
	v1.Post("/login", middleware.ValidateField[request.LoginRequest](), auth.Login)
	// v1.Post("/logout", ,auth.Login)
	// v1.Post("/register", register.Register)

	// members
	v1.Get("/members", adminOnly, member.GetData)
	v1.Get("/members/:id", adminOnly, middleware.ValidateField[model.Member](), member.GetDataById)
	v1.Put("/members/:id", adminOnly, member.UpdateData)
	v1.Post("/members", adminOnly, member.StoreData)

	// register
	v1.Post("/registers/rw", middleware.ValidateField[model.RWProfile](), rwProfileRegister.RegisterRWProfile)
	v1.Put("/registers/rw/:rwProfileId/approve", rwProfileRegister.ApproveRegistration)
	v1.Post("/registers/rw/user-account", rwProfileRegister.RegisterUserRw)
	v1.Post("/registers/rt", middleware.ValidateField[request.RTProfileRegisterRequest](), rtProfileRegister.RegisterRTProfile)
	v1.Put("/registers/rt/:rtProfileId/approve", rtProfileRegister.ApproveRegistration)
	v1.Post("/registers/rt/user-account", rtProfileRegister.RegisterUserRt)
	// v1.Post("/registers/rt/:referalCode", rtProfileRegister.RegisterRTProfile)

	//register tokens
	// v1.Get("/token-registrations", adminOnly, rwProfile.GetData)
	// v1.Get("/token-registrations/validate", adminOnly, rwProfile.GetData)

	//referal code
	v1.Get("/referal-codes", adminOnly, referalCode.GetData)
	v1.Get("/referal-codes/:id", adminOnly, referalCode.GetDataById)
	v1.Post("/referal-codes/validate", referalCode.ValidateReferalCode)

	//----------------------------------------------------------------
	//
	// Admin Authority
	//
	//----------------------------------------------------------------

	//rw profiles
	v1.Get("/rw-profiles", adminOnly, rwProfile.GetData)
	v1.Get("/rw-profiles/:id", adminOnly, rwProfile.GetDataById)

	//rt profiles
	v1.Get("/rt-profiles", adminOnly, rtProfile.GetData)
	v1.Get("/rt-profiles/:id", adminOnly, rtProfile.GetDataById)

	//users
	v1.Get("/users", adminOnly, user.GetData)
	v1.Get("/users/:id", adminOnly, user.GetDataById)

	//roles
	v1.Get("/roles", adminOnly, role.GetData)
	v1.Get("/roles/:id", adminOnly, role.GetDataById)
	v1.Post("/roles", adminOnly, role.StoreData)
	v1.Put("/roles/:id", adminOnly, role.UpdateData)
	v1.Delete("/roles/:id", adminOnly, role.DeleteData)

	//incoming letter
	v1.Get("/incoming-letters", adminOnly, incomingLetter.GetData)
	v1.Get("/incoming-letters/:id", adminOnly, incomingLetter.GetDataById)
	v1.Post("/incoming-letters", adminOnly, incomingLetter.StoreData)
	v1.Put("/incoming-letters/:id", adminOnly, incomingLetter.UpdateData)
	v1.Delete("/incoming-letters/:id", adminOnly, incomingLetter.DeleteData)

	//member status
	v1.Get("/member-status", adminOnly, memberStatus.GetData)
	v1.Get("/member-status/:id", adminOnly, memberStatus.GetDataById)
	v1.Post("/member-status", adminOnly, memberStatus.StoreData)
	v1.Put("/member-status/:id", adminOnly, memberStatus.UpdateData)
	v1.Delete("/member-status/:id", adminOnly, memberStatus.DeleteData)

	//village
	v1.Get("/villages", adminOnly, village.GetData)
	v1.Get("/villages/:id", adminOnly, village.GetDataById)
	v1.Post("/villages", adminOnly, village.StoreData)
	v1.Put("/villages/:id", adminOnly, village.UpdateData)
	v1.Delete("/villages/:id", adminOnly, village.DeleteData)

	//religion
	v1.Get("/religions", adminOnly, religion.GetData)
	v1.Get("/religions/:id", adminOnly, religion.GetDataById)
	v1.Post("/religions", adminOnly, religion.StoreData)
	v1.Put("/religions/:id", adminOnly, religion.UpdateData)
	v1.Delete("/religions/:id", adminOnly, religion.DeleteData)

	//letter type
	v1.Get("/letter-types", adminOnly, letterType.GetData)
	v1.Get("/letter-types/:id", adminOnly, letterType.GetDataById)
	v1.Post("/letter-types", adminOnly, letterType.StoreData)
	v1.Put("/letter-types/:id", adminOnly, letterType.UpdateData)
	v1.Delete("/letter-types/:id", adminOnly, letterType.DeleteData)

	//outcoming letter
	v1.Get("/outcoming-letters", adminOnly, OutcomingLetter.GetData)
	// v1.Get("/outcoming-letters/:id", adminOnly, OutcomingLetter.GetDataById)

	//letter req
	v1.Post("/request-letters", adminOnly, letterReq.CreateData)
	v1.Put("/request-letters/approve/:letter_req_id", adminOnly, letterReq.UpdateApprovalStatus)

}
