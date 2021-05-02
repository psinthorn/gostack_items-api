package services

var (
	// ItemsService tyep itemsInterfaceService and pointer to itemsService that implement method of interface require
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsService struct{}

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, error)
	Get(items.Item) (*items.Item, error)
}

func (i *itemsService) Create(item items.Items) (*items.Item, error) {}
func (i *itemsService) Get(item items.Item) (*items.Item, error)     {}
