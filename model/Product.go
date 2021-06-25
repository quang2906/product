package model

type Product struct {
	Id         int64   `json:"id"`
	CategoryId int64   `json:"categoryId"`
	Image      []Image `json:"image"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	IsSale     bool    `json:"isSales"`
	CreatedAt  int64   `json:"createdAt"`
	ModifiedAt int64   `json:"modifiedAt"`
	Rating     float64 `json:"rating"`
}
