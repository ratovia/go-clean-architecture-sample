package controllers

import (
	"net/http"
	"ratovia/go-clean-architecture-sample/app/src/entity"
	"ratovia/go-clean-architecture-sample/app/src/usecase"
	"strconv"

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

func (c *ItemController) CreateItem(ctx *gin.Context) (int, interface{}, error) {
	var item entity.Item

	if err := ctx.ShouldBindJSON(&item); err != nil {
		return http.StatusBadRequest, nil, err
	}

	newItem, err := c.Usecase.CreateItem(ctx, c.GetDB(ctx), &item)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, newItem, nil
}

func (c *ItemController) DeleteItem(ctx *gin.Context) (int, interface{}, error) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	if err := c.Usecase.DeleteItem(ctx, c.GetDB(ctx), uint(id)); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}

func (c *ItemController) UpdateItem(ctx *gin.Context) (int, interface{}, error) {
	var item entity.Item
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	if err := ctx.ShouldBindJSON(&item); err != nil {
		return http.StatusBadRequest, nil, err
	}

	updatedItem, err := c.Usecase.UpdateItem(ctx, c.GetDB(ctx), uint(id), &item)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, updatedItem, nil
}
