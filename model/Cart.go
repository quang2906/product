package model

type Cart struct {
	Id          int64   `json:"id"`
	CategoryId  int64   `json:"categoryId"`
	Image       string  `json:"image"`
	ProductName string  `json:"name"`
	Price       float64 `json:"price"`
	Quantity    int64   `json:"quantity"`
	CreatedAt   int64   `json:"createdAt"`
	ModifiedAt  int64   `json:"modifiedAt"`
}
