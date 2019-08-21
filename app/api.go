package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rossifedericoe/diproject/app/controllers/itemController"
)

type App struct {
	engine *gin.Engine
	itemController *itemController.ItemController
}

func (app *App) Run() {
	app.engine = gin.Default()
	app.engine.GET("/items/:id/site", app.itemController.GetSite)
	app.engine.Run()
}

func NewApp(itemController *itemController.ItemController) App {
	return App{itemController:itemController}
}
