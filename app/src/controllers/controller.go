package controllers

import (
	"context"
	"ratovia/go-clean-architecture-sample/app/src/gateways"
)
type Controller interface {
	GetDB(context.Context) gateways.DB
}
