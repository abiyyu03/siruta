package inventory

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/inventory"
	"github.com/gofiber/fiber/v2"
)

type InventoryUsecase struct{}

var inventoryRepository *inventory.InventoryRepository

type InventoryUsecaseInterface interface {
	Fetch(ctx *fiber.Ctx) error
	FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string) error
	FetchById(ctx *fiber.Ctx, id int) error
	Store(inventory *model.Inventory, ctx *fiber.Ctx) error
	Update(inventory *model.Inventory, ctx *fiber.Ctx, id int) error
	Delete(ctx *fiber.Ctx, id int) error
}

func (i *InventoryUsecase) Fetch(ctx *fiber.Ctx) error {
	inventories, err := inventoryRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &inventories, "Data fetched successfully")
}

func (i *InventoryUsecase) FetchById(ctx *fiber.Ctx, id int) error {
	inventory, err := inventoryRepository.FetchById(id)

	if inventory == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &inventory, "Data fetched successfully")
}

func (i *InventoryUsecase) FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string) error {
	inventories, err := inventoryRepository.FetchByRTProfileId(rtProfileId)
	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &inventories, "Data fetched successfully")
}

func (i *InventoryUsecase) Store(inventory *model.Inventory, ctx *fiber.Ctx) error {
	createdInventory := &model.Inventory{
		Name:        inventory.Name,
		Quantity:    inventory.Quantity,
		RTProfileId: inventory.RTProfileId,
		Image:       inventory.Image,
	}

	storedInventory, err := inventoryRepository.Store(createdInventory)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.SuccessCreated(ctx, &storedInventory, "Data updated successfully")
}

func (i *InventoryUsecase) Update(inventory *model.Inventory, ctx *fiber.Ctx, id int) error {

	updateInventory := &model.Inventory{
		Name:        inventory.Name,
		Quantity:    inventory.Quantity,
		RTProfileId: inventory.RTProfileId,
		Image:       inventory.Image,
	}

	updatedInventory, err := inventoryRepository.Update(updateInventory, id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &updatedInventory, "Data updated successfully")
}

func (i *InventoryUsecase) Delete(ctx *fiber.Ctx, id int) error {
	err := inventoryRepository.Delete(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, nil, "Data deleted successfully")
}
