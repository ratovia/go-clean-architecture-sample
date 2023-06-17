package usecase

import (
	"context"
	"fmt"

	"ratovia/go-clean-architecture-sample/app/src/gateways"
	"ratovia/go-clean-architecture-sample/app/src/entity"
)

type ItemUsecase struct {
}

func (*ItemUsecase) IndexItems(ctx context.Context, db gateways.DB) ([]entity.Item, error) {

	items, err := gateways.FindItemAll(db.GetQuerier())

	if err != nil {
		fmt.Println("Failed to select items:", err)
		return nil, err
	}

	return items, nil
}

func (u *ItemUsecase) CreateItem(ctx context.Context, db gateways.DB, item *entity.Item) (*entity.Item, error) {
	newItem, err := gateways.CreateItem(db.GetQuerier(), item)

	if err != nil {
		fmt.Println("Failed to create item:", err)
		return nil, err
	}

	return newItem, nil
}

func (u *ItemUsecase) DeleteItem(ctx context.Context, db gateways.DB, id uint) error {
	err := gateways.DeleteItem(db.GetQuerier(), id)

	if err != nil {
		fmt.Println("Failed to delete item:", err)
		return err
	}

	return nil
}

func (u *ItemUsecase) UpdateItem(ctx context.Context, db gateways.DB, id uint, item *entity.Item) (*entity.Item, error) {
	updatedItem, err := gateways.UpdateItem(db.GetQuerier(), id, item)

	if err != nil {
		fmt.Println("Failed to update item:", err)
		return nil, err
	}

	return updatedItem, nil
}
