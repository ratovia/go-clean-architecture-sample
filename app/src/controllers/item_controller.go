package controllers

import (
	"net/http"
	// "ratovia/go-clean-architecture-sample/app/src/entity"
	"ratovia/go-clean-architecture-sample/app/src/usecase"
	"github.com/gin-gonic/gin"
)

type ItemController struct {
	Controller
	Usecase usecase.ItemUsecase
}

func NewItemController(controller Controller) *ItemController {
	return &ItemController{
		Controller: controller,
	}
}

func (c *ItemController) IndexItem(ctx *gin.Context) (int, interface{}, error) {
	items, err := c.Usecase.IndexItems(ctx, c.GetDB(ctx))

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, items, nil
}
