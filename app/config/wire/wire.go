// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/rossifedericoe/diproject/app"
	"github.com/rossifedericoe/diproject/app/controllers/itemController"
	"github.com/rossifedericoe/diproject/app/services/itemService"
)

func InitializeApp() app.App {
	wire.Build(
		app.NewApp,
		itemController.NewItemController,
		itemService.NewItemServiceImpl,
		)
	return app.App{}
}