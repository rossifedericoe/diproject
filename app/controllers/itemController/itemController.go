package itemController

import (
	"github.com/gin-gonic/gin"
	"github.com/rossifedericoe/diproject/app/domain/item"
)

type ItemController struct {
	service item.ItemService
}

func NewItemController(pService item.ItemService) *ItemController {
	return &ItemController{service:pService}
}

func (controller *ItemController) GetSite(context *gin.Context)  {
	id := context.Param("id")
	context.JSON(200, controller.service.GetSite(id))
}
