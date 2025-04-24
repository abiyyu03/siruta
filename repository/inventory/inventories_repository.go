package inventory

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
)

type InventoryRepository struct{}

func (l *InventoryRepository) Fetch() ([]*model.Inventory, error) {
	var inventories []*model.Inventory

	if err := config.DB.Find(&inventories).Error; err != nil {
		return nil, err
	}

	return inventories, nil
}

func (l *InventoryRepository) FetchById(id string) (*model.Inventory, error) {
	var inventory *model.Inventory

	if err := config.DB.Where("id = ?", id).First(&inventory).Error; err != nil {
		return nil, err
	}

	return inventory, nil
}

func (v *InventoryRepository) Store(inventory *model.Inventory) (*model.Inventory, error) {
	if err := config.DB.Create(&inventory).Error; err != nil {
		return nil, err
	}

	return inventory, nil
}

func (v *InventoryRepository) Update(inventory *model.Inventory, id int) (*model.Inventory, error) {
	if err := config.DB.Where("id = ?", id).Updates(&inventory).Error; err != nil {
		return nil, err
	}

	return inventory, nil
}

func (v *InventoryRepository) Delete(id int) error {
	var inventory model.Inventory

	if err := config.DB.Where("id = ?", id).Delete(&inventory).Error; err != nil {
		return err
	}

	return nil
}
