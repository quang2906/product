package model

type Review struct {
	Id         int64   `json:"id"`
	ProductId  int64   `json:"productId"`
	Rating     float64 `json:"rating"`
	CreatedAt  int64   `json:"createdAt"`
	ModifiedAt int64   `json:"modifiedAt"`
}
