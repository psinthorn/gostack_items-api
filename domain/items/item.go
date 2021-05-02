package items

type Item struct {
	Id               string      `json: "id"`
	Seller           int64       `json: "seller"`
	Title            string      `json: "title"`
	Description      Description `json: "description"`
	Pictures         []Picture   `json: "pictures"`
	Videos           []Video     `json: "videos"`
	Price            float32     `json: "price"`
	AvailableQuatity int         `json: "available_quatity"`
	SoldQuantity     int         `json: "sold_quatity"`
	Status           string      `json: "status"`
	CreatedAt        string      `json: "created_at"`
	UpdatedAt        string      `json: "updated_at"`
}

type Description struct {
	Id        string `json: "id"`
	PlainText string `json: "plain_text"`
	Html      string `json: "html"`
}

type Picture struct {
	Id    string `json: "id"`
	Title string `json: "title"`
	Url   string `json: "url"`
}

type Video struct {
	Id    string `json: "id"`
	Title string `json: "title"`
	Url   string `json: "url"`
}
