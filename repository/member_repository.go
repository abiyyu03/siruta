package repository

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/gofiber/fiber/v2"
)

type MemberRepository struct{}

func (m *MemberRepository) Fetch(ctx *fiber.Ctx) (*model.Member, error) {
	var member *model.Member

	config.DB.Find(&member)

	return member, nil
}
