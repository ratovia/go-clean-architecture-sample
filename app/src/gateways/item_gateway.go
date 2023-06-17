package gateways

import (
	"fmt"

	"ratovia/go-clean-architecture-sample/app/src/entity"

	"gorm.io/gorm"
)

type Item struct {
	Model
	Price uint   `gorm:"not null"`
	Name  string `gorm:"not null"`
}

func (i *Item) toEntity() *entity.Item {
	return &entity.Item{
		ID:    i.ID,
		Price: i.Price,
		Name:  i.Name,
	}
}

func newItem(i *entity.Item) *Item {
	return &Item{
		Model: Model{
			ID: i.ID,
		},
		Price: i.Price,
		Name:  i.Name,
	}
}

func newItems(es []entity.Item) []Item {
	items := []Item{}
	for _, e := range es {
		items = append(items, *newItem(&e))
	}
	return items
}

func FindItemAll(db *gorm.DB) ([]entity.Item, error) {
	var items []Item
	if err := db.Find(&items).Error; err != nil {
		return nil, fmt.Errorf("failed to find all items: %w", err)
	}

	var entityItems []entity.Item
	for _, item := range items {
		entityItems = append(entityItems, *item.toEntity())
	}

	return entityItems, nil
}

func UpdateItem(db *gorm.DB, id uint, eItem *entity.Item) (*entity.Item, error) {
	item := newItem(eItem)
	item.ID = id

	if err := db.Save(item).Error; err != nil {
		return nil, fmt.Errorf("failed to update item: %w", err)
	}

	return item.toEntity(), nil
}

func CreateItem(db *gorm.DB, eItem *entity.Item) (*entity.Item, error) {
	item := newItem(eItem)

	if err := db.Create(item).Error; err != nil {
		return nil, fmt.Errorf("failed to create item: %w", err)
	}

	return item.toEntity(), nil
}

func DeleteItem(db *gorm.DB, id uint) error {
	item := &Item{Model: Model{ID: id}}

	if err := db.Delete(item).Error; err != nil {
		return fmt.Errorf("failed to delete item: %w", err)
	}

	return nil
}
