package gateways

import (
	"time"
)

type Model struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"-"`
	UpdatedAt time.Time `gorm:"-"`
}
