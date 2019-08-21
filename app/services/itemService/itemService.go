package itemService

import "github.com/rossifedericoe/diproject/app/domain/item"

type ItemServiceImpl struct {

}

func NewItemServiceImpl() item.ItemService {
	return &ItemServiceImpl{}
}

func (service *ItemServiceImpl) GetSite(id string) string {
	return id[:3]
}
