package guest_list

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/guest_list"
	"github.com/gofiber/fiber/v2"
)

type GuestListUsecase struct{}

var GuestListRepository *guest_list.GuestListRepository

type GuestListUsecaseInterface interface {
	Fetch(ctx *fiber.Ctx) error
	FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string) error
	FetchById(ctx *fiber.Ctx, id int) error
	Store(ctx *fiber.Ctx, guestList *model.GuestList) error
	Update(ctx *fiber.Ctx, guestList *model.GuestList, id int) error
	Delete(ctx *fiber.Ctx, id int) error
}

func (g *GuestListUsecase) Fetch(ctx *fiber.Ctx) error {
	guestLists, err := GuestListRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &guestLists, "Data fetched successfully")
}

func (g *GuestListUsecase) FetchById(ctx *fiber.Ctx, id int) error {
	guestList, err := GuestListRepository.FetchById(id)

	if guestList == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &guestList, "Data fetched successfully")
}
func (g *GuestListUsecase) FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string) error {
	guestList, err := GuestListRepository.FetchByRTProfileId(rtProfileId)

	if guestList == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &guestList, "Data fetched successfully")
}

func (g *GuestListUsecase) Store(ctx *fiber.Ctx, guestListData *model.GuestList) error {
	guestList, err := GuestListRepository.Store(guestListData)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, guestList, "Data stored successfully")
}

func (g *GuestListUsecase) Update(ctx *fiber.Ctx, guestListData *model.GuestList, id int) error {
	guestList, err := GuestListRepository.Update(guestListData, id)

	if guestList == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, guestList, "Data updated successfully")
}

func (g *GuestListUsecase) Delete(ctx *fiber.Ctx, id int) error {
	_, err := GuestListRepository.Delete(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return nil
}
