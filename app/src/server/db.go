package server

import (
	"context"
	"fmt"
	"ratovia/go-clean-architecture-sample/app/src/gateways"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func InitializeDB() *DB {
	dsn := "root@tcp(127.0.0.1:3306)/clean_architecture_sample_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Errorf("failed to connect to the database: %v", err)
		panic("exit")
	}

	return &DB{
		db: db,
	}
}

func (db *DB) WithContext(ctx context.Context) gateways.DB {
	return &DB{
		db: db.db.WithContext(ctx),
	}
}

// Begin begin new transaction
func (db *DB) Begin() (*gorm.DB, func() error, func(), error) {
	tx := db.db.Begin()
	if tx.Error != nil {
		return nil, nil, nil, tx.Error
	}
	commit := func() error {
		return tx.Commit().Error
	}
	rollback := func() {
		tx.Rollback()
	}
	return tx, commit, rollback, tx.Error
}

// GetQuerier get queyer with no transaction
func (db *DB) GetQuerier() *gorm.DB {
	return db.db
}
