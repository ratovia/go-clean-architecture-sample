package server

import (
	"context"
	"ratovia/go-clean-architecture-sample/app/src/gateways"
)

type Controller struct {
	db *DB
}

func NewController(db *DB, env string) *Controller {
	// 他のコントローラの初期化処理を追加する場合はここに記述する
	// 環境変数とかシークレットとか
	controller := &Controller{
		db: db,
	}

	return controller
}

// GetDB get db
func (c *Controller) GetDB(ctx context.Context) gateways.DB {
	return c.db.WithContext(ctx)
}
