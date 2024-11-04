package http

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/repository"
	"github.com/gofiber/fiber/v2"
)

type MemberHttp struct{}

func (m *MemberHttp) GetData(ctx *fiber.Ctx) error {
	memberRepository := new(repository.MemberRepository)

	members, err := memberRepository.Fetch(ctx)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, "Error fetching data")
	}

	return entity.Success(ctx, &members, "Data fetched successfully")
}
