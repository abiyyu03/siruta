package http

import (
	"github.com/abiyyu03/siruta/usecase"
	"github.com/gofiber/fiber/v2"
)

type OutcomingLetterHttp struct{}

var outcomingLetterUsecase = new(usecase.OutcomingLetterUsecase)

func (l *OutcomingLetterHttp) GetData(ctx *fiber.Ctx) error {
	return outcomingLetterUsecase.Fetch(ctx)
}
