package services

import (
	"github.com/psinthorn/gostack_items-api/domain/items"
)

var (
	// ItemsService tyep itemsInterfaceService and pointer to itemsService that implement method of interface require
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsService struct{}

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, error)
	Get(items.Item) (*items.Item, error)
}

func (i *itemsService) Create(item items.Item) (*items.Item, error) {
	return nil, nil
}
func (i *itemsService) Get(item items.Item) (*items.Item, error) {
	return nil, nil
}
