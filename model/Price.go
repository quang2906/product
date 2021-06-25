package model

type Price struct {
	Id         int64   `json:"id"`
	ProductId  int64   `json:"product_id"`
	Value      float64 `json:"value"`
	CreatedAt  int64   `json:"createdAt"`
	ModifiedAt int64   `json:"modifiedAt"`
}
