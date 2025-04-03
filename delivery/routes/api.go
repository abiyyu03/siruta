package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/http/auth"
	"github.com/abiyyu03/siruta/delivery/http/finance"
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
	memberRegister := new(register.MemberRegisterHttp)
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
	rtLeader := new(http.RTLeaderHttp)
	rwLeader := new(http.RWLeaderHttp)
	inventory := new(http.InventoryHttp)
	guestList := new(http.GuestListHttp)
	incomePlan := new(finance.IncomePlanHttp)
	incomeLog := new(finance.IncomeHttp)
	expenseLog := new(finance.ExpenseHttp)
	resetPassword := new(auth.ResetPasswordHttp)

	adminOnly := middleware.JWTMiddleware([]int{1})
	// rwLeaderOnly := middleware.JWTMiddleware([]int{2})
	rtLeaderOnly := middleware.JWTMiddleware([]int{3})
	memberOnly := middleware.JWTMiddleware([]int{4})
	// validator := middleware.ValidateField()

	//authentication
	v1.Post("/login", middleware.ValidateField[request.LoginRequest](), auth.Login)
	v1.Post("/forgot-password", resetPassword.SendForgotPasswordLink)
	v1.Put("/reset-password", middleware.ValidateField[request.ResetPassword](), resetPassword.ResetPassword)

	// members
	v1.Get("/members", adminOnly, member.GetData)
	v1.Get("/members/:id", adminOnly, member.GetDataById)
	v1.Put("/members/:id", adminOnly, member.UpdateData)
	v1.Get("/members/:rt_profile_id/rt", adminOnly, rtLeaderOnly, member.GetDataByRTProfileId)

	// register
	v1.Post("/registers/rw", middleware.ValidateField[model.RWProfile](), rwProfileRegister.RegisterRWProfile)
	v1.Put("/registers/rw/:rwProfileId/approve", rwProfileRegister.ApproveRegistration)
	v1.Post("/registers/rw/user-account", rwProfileRegister.RegisterUserRw)
	v1.Post("/registers/rt", middleware.ValidateField[request.RTProfileRegisterRequest](), rtProfileRegister.RegisterRTProfile)
	v1.Put("/registers/rt/:rtProfileId/approve", rtProfileRegister.ApproveRegistration)
	v1.Post("/registers/rt/user-account", rtProfileRegister.RegisterUserRt)
	v1.Post("/registers/member", middleware.ValidateField[request.MemberRegisterRequest](), memberRegister.Register)

	//----------------------------------------------------------------
	//
	// RW Authority
	//
	//----------------------------------------------------------------

	//----------------------------------------------------------------
	//
	// RT Authority
	//
	//----------------------------------------------------------------

	//inventory
	v1.Get("/inventories", rtLeaderOnly, inventory.GetData)
	v1.Get("/inventories/:id", rtLeaderOnly, inventory.GetDataById)
	v1.Post("/inventories", rtLeaderOnly, inventory.StoreData)
	v1.Put("/inventories/:id", rtLeaderOnly, inventory.UpdateData)
	v1.Delete("/inventories/:id", rtLeaderOnly, inventory.DeleteData)

	//referal codes
	v1.Get("/referal-codes/:profile_id/rt", rtLeaderOnly, referalCode.GetDataByRTProfileId)
	v1.Put("/referal-codes/:profile_id/rt/regenerate/:code", rtLeaderOnly, referalCode.RegenerateCode)

	// incoming letter
	v1.Get("/incoming-letters/:rt_profile_id/rt", rtLeaderOnly, incomingLetter.GetDataByRTProfileId)

	// guest list
	v1.Get("/guest-lists", adminOnly, guestList.GetData)
	v1.Get("/guest-lists/:rt_profile_id/rt", adminOnly, rtLeaderOnly, guestList.GetDataByRTProfileId)
	v1.Get("/guest-lists/:id", adminOnly, rtLeaderOnly, guestList.GetDataById)
	v1.Put("/guest-lists/:id", adminOnly, rtLeaderOnly, guestList.UpdateData)
	v1.Delete("/guest-lists/:id", adminOnly, rtLeaderOnly, guestList.DeleteData)

	// Request letter
	v1.Put("/request-letters/approve/:letter_req_id", adminOnly, rtLeaderOnly, letterReq.UpdateApprovalStatus)

	// income finance plan
	v1.Get("/finances/income-plan", adminOnly, incomePlan.GetData)
	v1.Get("/finances/income-plan/:rt_profile_id/rt", adminOnly, rtLeaderOnly, incomePlan.GetDataByRTProfileId)
	v1.Get("/finances/income-plan/:id", adminOnly, rtLeaderOnly, incomePlan.GetDataById)
	v1.Post("/finances/income-plan", incomePlan.StoreData)
	v1.Put("/finances/income-plan/:id", adminOnly, rtLeaderOnly, incomePlan.UpdateData)
	v1.Delete("/finances/income-plan/:id", adminOnly, rtLeaderOnly, incomePlan.DeleteData)

	// income finance logs
	v1.Get("/finances/income-log", adminOnly, incomeLog.GetData)
	v1.Get("/finances/income-log/:rt_profile_id/rt", adminOnly, rtLeaderOnly, incomeLog.GetDataByRTProfileId)
	v1.Get("/finances/income-log/:id", adminOnly, rtLeaderOnly, incomeLog.GetDataById)
	v1.Put("/finances/income-log/:id", adminOnly, rtLeaderOnly, incomeLog.UpdateData)
	v1.Post("/finances/income-log", incomeLog.StoreData)
	v1.Delete("/finances/income-log/:id", adminOnly, rtLeaderOnly, incomeLog.DeleteData)

	// income expense logs
	v1.Get("/finances/expense-log", adminOnly, expenseLog.GetData)
	v1.Get("/finances/expense-log/:rt_profile_id/rt", adminOnly, rtLeaderOnly, expenseLog.GetDataByRTProfileId)
	v1.Get("/finances/expense-log/:id", adminOnly, rtLeaderOnly, expenseLog.GetDataById)
	v1.Put("/finances/expense-log/:id", adminOnly, rtLeaderOnly, expenseLog.UpdateData)
	v1.Post("/finances/expense-log", expenseLog.StoreData)
	v1.Delete("/finances/expense-log/:id", adminOnly, rtLeaderOnly, expenseLog.DeleteData)

	//----------------------------------------------------------------
	//
	// Admin Authority
	//
	//----------------------------------------------------------------

	//rw profiles
	v1.Get("/rw-profiles", adminOnly, rwProfile.GetData)
	v1.Get("/rw-profiles/:id", adminOnly, rtLeaderOnly, rwProfile.GetDataById)

	//rt profiles
	v1.Get("/rt-profiles", adminOnly, rtProfile.GetData)
	v1.Get("/rt-profiles/:id", adminOnly, rtLeaderOnly, rtProfile.GetDataById)

	//referal code
	v1.Get("/referal-codes", adminOnly, referalCode.GetData)
	v1.Get("/referal-codes/:id", adminOnly, referalCode.GetDataById)
	v1.Post("/referal-codes/validate", referalCode.ValidateReferalCode)

	//users
	v1.Get("/users", adminOnly, user.GetData)
	v1.Get("/users/:id", adminOnly, user.GetDataById)

	//leaders
	v1.Get("/rt-leaders", adminOnly, rtLeader.GetData)
	v1.Get("/rt-leaders/:id", adminOnly, rtLeader.GetDataById)
	v1.Put("/rt-leaders/:id", adminOnly, rtLeader.UpdateData)
	v1.Get("/rw-leaders", adminOnly, rwLeader.GetData)
	v1.Get("/rw-leaders/:id", adminOnly, rwLeader.GetDataById)
	v1.Put("/rw-leaders/:id", adminOnly, rwLeader.UpdateData)

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
	v1.Get("/outcoming-letters/:id", adminOnly, rtLeaderOnly, OutcomingLetter.GetDataById)
	v1.Get("/outcoming-letters/:rt_profile_id/rt", adminOnly, rtLeaderOnly, OutcomingLetter.GetDataByRTProfileId)

	//letter req
	v1.Post("/request-letters", adminOnly, memberOnly, letterReq.CreateData)
}
