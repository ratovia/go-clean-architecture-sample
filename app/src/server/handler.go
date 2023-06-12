package server

import (
	"net/http"

	"ratovia/go-clean-architecture-sample/app/src/controllers"

	"github.com/gin-gonic/gin"
)

type rootAPI struct {
	path    string
	method  string
	handler apiHandler
}

func newRootAPI(path string, method string, handler apiHandler) *rootAPI {
	return &rootAPI{
		path:    path,
		method:  method,
		handler: handler,
	}
}

func handleRootGroup(controller *Controller, group *gin.RouterGroup) {
	itemController := controllers.NewItemController(controller)

	apiV1 := group.Group("/v1")

	get := http.MethodGet

	apis := []*rootAPI{
		newRootAPI("/items", get, itemController.IndexItem),
	}

	for _, a := range apis {
		handler := a.handler
		h := func(ctx *gin.Context) {
			status, res, _ := handler(ctx)
			ctx.JSON(status, res)
		}

		apiV1.Handle(a.method, a.path, h)
	}
}
