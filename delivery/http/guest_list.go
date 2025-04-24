package http

import (
	"strconv"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase/guest_list"
	"github.com/gofiber/fiber/v2"
)

type GuestListHttp struct{}

var guestListUsecase *guest_list.GuestListUsecase

func (g *GuestListHttp) GetData(ctx *fiber.Ctx) error {
	return guestListUsecase.Fetch(ctx)
}

func (g *GuestListHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	convertedId, _ := strconv.Atoi(id)

	return guestListUsecase.FetchById(ctx, convertedId)
}

func (g *GuestListHttp) GetDataByRTProfileId(ctx *fiber.Ctx) error {
	rtProfileId := ctx.Params("rt_profile_id")

	return guestListUsecase.FetchByRTProfileId(ctx, rtProfileId)
}

func (g *GuestListHttp) CreateData(ctx *fiber.Ctx) error {
	var guestListData *model.GuestList

	if err := ctx.BodyParser(&guestListData); err != nil {
		return err
	}

	return guestListUsecase.Store(ctx, guestListData)
}

func (g *GuestListHttp) UpdateData(ctx *fiber.Ctx) error {
	var guestList *model.GuestList
	id := ctx.Params("id")
	convertedId, _ := strconv.Atoi(id)

	if err := ctx.BodyParser(&guestList); err != nil {
		return err
	}

	return guestListUsecase.Update(ctx, guestList, convertedId)
}

func (g *GuestListHttp) DeleteData(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	convertedId, _ := strconv.Atoi(id)

	return guestListUsecase.Delete(ctx, convertedId)
}
