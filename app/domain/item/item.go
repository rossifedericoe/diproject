package item

type Item struct {

}

type ItemService interface {
	GetSite(id string) string
}
