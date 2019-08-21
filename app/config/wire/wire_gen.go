// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package wire

import (
	"github.com/rossifedericoe/diproject/app"
	"github.com/rossifedericoe/diproject/app/controllers/itemController"
	"github.com/rossifedericoe/diproject/app/services/itemService"
)

// Injectors from wire.go:

func InitializeApp() app.App {
	itemItemService := itemService.NewItemServiceImpl()
	itemControllerItemController := itemController.NewItemController(itemItemService)
	appApp := app.NewApp(itemControllerItemController)
	return appApp
}
