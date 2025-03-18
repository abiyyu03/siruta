package http

import (
	"strconv"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase/inventory"
	"github.com/gofiber/fiber/v2"
)

type InventoryHttp struct {
	inventoryUsecase *inventory.InventoryUsecase
}

func (l *InventoryHttp) GetData(ctx *fiber.Ctx) error {
	return l.inventoryUsecase.Fetch(ctx)
}

func (l *InventoryHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return l.inventoryUsecase.FetchById(ctx, id)
}

func (l *InventoryHttp) StoreData(ctx *fiber.Ctx) error {
	var inventory *model.Inventory

	if err := ctx.BodyParser(&inventory); err != nil {
		return err
	}

	return l.inventoryUsecase.Store(inventory, ctx)
}

func (l *InventoryHttp) UpdateData(ctx *fiber.Ctx) error {
	var inventory *model.Inventory
	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := ctx.BodyParser(&inventory); err != nil {
		return err
	}

	return l.inventoryUsecase.Update(inventory, ctx, id)
}

func (l *InventoryHttp) DeleteData(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return l.inventoryUsecase.Delete(ctx, id)
}
