package http

import (
	"strconv"

	"github.com/abiyyu03/siruta/usecase/referal_code"
	"github.com/gofiber/fiber/v2"
)

type ReferalCodeHttp struct{}

var referalCodeUsecase *referal_code.ReferalCodeUsecase

func (r *ReferalCodeHttp) GetData(ctx *fiber.Ctx) error {
	return referalCodeUsecase.Fetch(ctx)
}

func (r *ReferalCodeHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return referalCodeUsecase.FetchById(ctx, id)
}

func (r *ReferalCodeHttp) GetDataByRTProfileId(ctx *fiber.Ctx) error {
	profileId := ctx.Params("profile_id")

	return referalCodeUsecase.FetchByRTProfileId(ctx, profileId)
}

func (r *ReferalCodeHttp) RegenerateCode(ctx *fiber.Ctx) error {
	profileId := ctx.Params("profile_id")
	code := ctx.Params("code")

	return referalCodeUsecase.RegenerateReferalCode(ctx, profileId, code)
}

func (r *ReferalCodeHttp) ValidateReferalCode(ctx *fiber.Ctx) error {
	queryParam := ctx.Queries()

	response, _ := referalCodeUsecase.Validate(ctx, queryParam["code"])

	return response
}

func (r *ReferalCodeHttp) DeleteData(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return referalCodeUsecase.Delete(ctx, id)
}
