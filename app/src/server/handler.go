package server

import (
	"net/http"

	"ratovia/go-clean-architecture-sample/app/src/controllers"

	"github.com/gin-gonic/gin"
)

type apiHandler func(*gin.Context) (int, interface{}, error)

func handleRootGroup(controller *Controller, group *gin.RouterGroup) {
	itemController := controllers.NewItemController(controller)

	apiV1 := group.Group("/v1")

	registerHandler(apiV1, http.MethodGet, "/items", itemController.IndexItem)
	registerHandler(apiV1, http.MethodPost, "/items", itemController.CreateItem)
	registerHandler(apiV1, http.MethodDelete, "/items/:id", itemController.DeleteItem)
	registerHandler(apiV1, http.MethodPut, "/items/:id", itemController.UpdateItem)
}

func registerHandler(group *gin.RouterGroup, method string, path string, handler apiHandler) {
	h := func(ctx *gin.Context) {
		status, res, _ := handler(ctx)
		ctx.JSON(status, res)
	}

	group.Handle(method, path, h)
}
