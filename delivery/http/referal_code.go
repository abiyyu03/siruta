package http

import "github.com/gofiber/fiber/v2"

type ReferalCodeHttp struct{}

func (r *ReferalCodeHttp) GetData(ctx *fiber.Ctx) error {
	return referalCodeUsecase.Fetch(ctx)
}

func (r *ReferalCodeHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return referalCodeUsecase.FetchById(ctx, id)
}

func (r *ReferalCodeHttp) ValidateReferalCode(ctx *fiber.Ctx) error {
	queryParam := ctx.Queries()

	return referalCodeUsecase.Validate(ctx, queryParam["code"])
}
