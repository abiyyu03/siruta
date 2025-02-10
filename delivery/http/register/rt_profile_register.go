package register

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase"
	"github.com/gofiber/fiber/v2"
)

type RTProfileRegisterHttp struct{}

// type RTProfileRequest struct {
// 	RTProfile   *model.RTProfile `json:"rt_profile"`
// 	ReferalCode string           `json:"referal_code" validate:"required"`
// }

var rtProfileRegistrationUsecase = new(usecase.RTProfileRegistrationUsecase)
var userUsecase = new(usecase.UserUsecase)

func (r *RTProfileRegisterHttp) RegisterRTProfile(ctx *fiber.Ctx) error {
	var request *model.RTProfile
	queryParams := ctx.Queries()

	if err := ctx.BodyParser(&request); err != nil {
		return entity.Error(ctx, fiber.StatusUnprocessableEntity, fiber.ErrUnprocessableEntity.Message)
	}

	return rtProfileRegistrationUsecase.Register(ctx, request, queryParams["referal_code"])
}

func (r *RTProfileRegisterHttp) RegisterUserAccount(ctx *fiber.Ctx) error {
	var user *model.User
	queryParams := ctx.Queries()

	if err := ctx.BodyParser(&user); err != nil {
		return entity.Error(ctx, fiber.StatusUnprocessableEntity, fiber.ErrUnprocessableEntity.Message)
	}

	return userUsecase.RegisterUserWithTokenVerification(ctx, user, queryParams["token"])

}

func (r *RTProfileRegisterHttp) ApproveRegistrant(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return rtProfileRegistrationUsecase.ApproveRegistrant(ctx, id)
}

// func (r *RTProfileRegisterHttp) VerifyRWReferalCode(ctx *fiber.Ctx) error {
// 	code := ctx.Params("code")

// 	RTProfileRegistrationRepository := new(repository.RTProfileRegistrationRepository)
// 	verified, err := RTProfileRegistrationRepository.VerifyRWReferalCode(code)

// 	if err != nil {
// 		return entity.Error(ctx, fiber.StatusInternalServerError, err.Error()) //fiber.ErrInternalServerError.Message)
// 	}

// 	return entity.Success(ctx, &verified, "Referral code verified successfully")
// }

// func (r *RTProfileRegisterHttp) ApproveRegistration(ctx *fiber.Ctx) error {

// }
