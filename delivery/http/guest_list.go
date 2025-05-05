package http

import (
	"strconv"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase/guest_list"
	"github.com/gofiber/fiber/v2"
)

type GuestListHttp struct {
	guestListUsecase guest_list.GuestListUsecaseInterface
}

func NewGuestListHttp(guestListUC guest_list.GuestListUsecaseInterface) *GuestListHttp {
	return &GuestListHttp{
		guestListUsecase: guestListUC,
	}
}

func (g *GuestListHttp) GetData(ctx *fiber.Ctx) error {
	return g.guestListUsecase.Fetch(ctx)
}

func (g *GuestListHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	convertedId, _ := strconv.Atoi(id)

	return g.guestListUsecase.FetchById(ctx, convertedId)
}

func (g *GuestListHttp) GetDataByRTProfileId(ctx *fiber.Ctx) error {
	rtProfileId := ctx.Params("rt_profile_id")

	return g.guestListUsecase.FetchByRTProfileId(ctx, rtProfileId)
}

func (g *GuestListHttp) CreateData(ctx *fiber.Ctx) error {
	var guestListData *model.GuestList

	if err := ctx.BodyParser(&guestListData); err != nil {
		return err
	}

	return g.guestListUsecase.Store(ctx, guestListData)
}

func (g *GuestListHttp) UpdateData(ctx *fiber.Ctx) error {
	var guestList *model.GuestList
	id := ctx.Params("id")
	convertedId, _ := strconv.Atoi(id)

	if err := ctx.BodyParser(&guestList); err != nil {
		return err
	}

	return g.guestListUsecase.Update(ctx, guestList, convertedId)
}

func (g *GuestListHttp) DeleteData(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	convertedId, _ := strconv.Atoi(id)

	return g.guestListUsecase.Delete(ctx, convertedId)
}
