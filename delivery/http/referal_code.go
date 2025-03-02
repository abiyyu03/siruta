package http

import (
	"github.com/abiyyu03/siruta/usecase/referal_code"
	"github.com/gofiber/fiber/v2"
)

type ReferalCodeHttp struct {
	referalCodeUsecase *referal_code.ReferalCodeUsecase
}

func (r *ReferalCodeHttp) GetData(ctx *fiber.Ctx) error {
	return r.referalCodeUsecase.Fetch(ctx)
}

func (r *ReferalCodeHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return r.referalCodeUsecase.FetchById(ctx, id)
}

func (r *ReferalCodeHttp) ValidateReferalCode(ctx *fiber.Ctx) error {
	queryParam := ctx.Queries()

	return r.referalCodeUsecase.Validate(ctx, queryParam["code"])
}
