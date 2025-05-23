package http

import (
	"strconv"

	"github.com/abiyyu03/siruta/usecase/referal_code"
	"github.com/gofiber/fiber/v2"
)

type ReferalCodeHttp struct {
	referalCodeUsecase referal_code.ReferalCodeUsecaseInterface
}

func NewReferalCodeHttp(referalCodeUC referal_code.ReferalCodeUsecaseInterface) *ReferalCodeHttp {
	return &ReferalCodeHttp{
		referalCodeUsecase: referalCodeUC,
	}
}

func (r *ReferalCodeHttp) GetData(ctx *fiber.Ctx) error {
	return r.referalCodeUsecase.Fetch(ctx)
}

func (r *ReferalCodeHttp) GetDataById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return r.referalCodeUsecase.FetchById(ctx, id)
}

func (r *ReferalCodeHttp) GetDataByRTProfileId(ctx *fiber.Ctx) error {
	profileId := ctx.Params("profile_id")

	return r.referalCodeUsecase.FetchByRTProfileId(ctx, profileId)
}

func (r *ReferalCodeHttp) RegenerateCode(ctx *fiber.Ctx) error {
	profileId := ctx.Params("profile_id")
	code := ctx.Params("code")

	return r.referalCodeUsecase.RegenerateReferalCode(ctx, profileId, code)
}

func (r *ReferalCodeHttp) ValidateReferalCode(ctx *fiber.Ctx) error {
	queryParam := ctx.Queries()

	response, _ := r.referalCodeUsecase.Validate(ctx, queryParam["code"])

	return response
}

func (r *ReferalCodeHttp) DeleteData(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return r.referalCodeUsecase.Delete(ctx, id)
}
