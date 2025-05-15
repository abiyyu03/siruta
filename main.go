package main

import (
	_ "github.com/abiyyu03/siruta/config/timezone"
	"github.com/abiyyu03/siruta/delivery/http"
	authHttp "github.com/abiyyu03/siruta/delivery/http/auth"
	registerHttp "github.com/abiyyu03/siruta/delivery/http/register"
	"github.com/abiyyu03/siruta/delivery/routes"
	"github.com/abiyyu03/siruta/repository/database/seeder"
	"github.com/abiyyu03/siruta/usecase/auth"
	"github.com/abiyyu03/siruta/usecase/finance"
	"github.com/abiyyu03/siruta/usecase/guest_list"
	"github.com/abiyyu03/siruta/usecase/inventory"
	"github.com/abiyyu03/siruta/usecase/letter"
	"github.com/abiyyu03/siruta/usecase/letter_type"
	"github.com/abiyyu03/siruta/usecase/member"
	"github.com/abiyyu03/siruta/usecase/member_status"
	"github.com/abiyyu03/siruta/usecase/referal_code"
	"github.com/abiyyu03/siruta/usecase/register"
	"github.com/abiyyu03/siruta/usecase/religion"
	"github.com/abiyyu03/siruta/usecase/role"
	"github.com/abiyyu03/siruta/usecase/rt_profile"
	"github.com/abiyyu03/siruta/usecase/rw_profile"
	"github.com/abiyyu03/siruta/usecase/user"
	"github.com/abiyyu03/siruta/usecase/village"

	"github.com/abiyyu03/siruta/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitDB()

	seed := new(seeder.SeederStruct)
	seed.RunSeeders()

	app := fiber.New()
	routeDefinition := &routes.HandlerDefinition{
		Auth:              authHttp.NewAuthHttp(&auth.AuthUsecase{}),
		MemberRegister:    registerHttp.NewMemberRegisterHttp(&register.MemberRegisterUsecase{}),
		RWProfileRegister: registerHttp.NewRWProfileRegisterHttp(&rw_profile.RWProfileRegisterUsecase{}),
		RTProfileRegister: registerHttp.NewRTProfileRegisterHttp(&rt_profile.RTProfileRegisterUsecase{}),
		Village:           http.NewVillageHttp(&village.VillageUsecase{}),
		LetterType:        http.NewLetterTypeHttp(&letter_type.LetterTypeUsecase{}),
		RWProfile:         http.NewRWProfileHttp(&rw_profile.RWProfileUsecase{}),
		RTProfile:         http.NewRTProfileHttp(&rt_profile.RTProfileUsecase{}),
		Member:            http.NewMemberHttp(&member.MemberUsecase{}),
		Role:              http.NewRoleHttp(&role.RoleUsecase{}),
		IncomingLetter:    http.NewIncomingLetterHttp(&letter.IncomingLetterUsecase{}),
		LetterReq:         http.NewLetterReqHttp(&letter.LetterReqUsecase{}),
		OutcomingLetter:   http.NewOutcomingLetterHttp(&letter.OutcomingLetterUsecase{}),
		MemberStatus:      http.NewMemberStatusHttp(&member_status.MemberStatusUsecase{}),
		User:              http.NewUserHttp(&user.UserUsecase{}),
		Religion:          http.NewReligionHttp(&religion.ReligionUsecase{}),
		ReferalCode:       http.NewReferalCodeHttp(&referal_code.ReferalCodeUsecase{}),
		RTLeader:          http.NewRTLeaderHttp(&rt_profile.RTLeaderUsecase{}),
		RWLeader:          http.NewRWLeaderHttp(&rw_profile.RWLeaderUsecase{}),
		Inventory:         http.NewInventoryHttp(&inventory.InventoryUsecase{}),
		GuestList:         http.NewGuestListHttp(&guest_list.GuestListUsecase{}),
		Cashflow:          http.NewCashflowHttp(&finance.CashflowUsecase{}),
		ResetPassword:     authHttp.NewResetPasswordHttp(&auth.ResetPasswordUsecase{}),
	}

	routeDefinition.HttpRoutes(app)

	app.Listen(":8080")
}
