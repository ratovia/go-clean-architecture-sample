package usecase

import "gorm.io/gorm"

type DB interface {
	Begin() (*gorm.DB, func() error, func(), error)
	GetQuerier() *gorm.DB
}
