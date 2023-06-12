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
