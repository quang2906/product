package repository

import (
	"errors"
	"fmt"
	"time"

	"example.com/product/model"
)

type ProductRepo struct {
	products map[int64]*model.Product
	autoID   int64
}

var Products ProductRepo

func init() {
	Products = ProductRepo{autoID: 0}
	Products.products = make(map[int64]*model.Product)
	Products.InitData("sql:45312")
}

func (r *ProductRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *ProductRepo) CreateNewProduct(product *model.Product) int64 {
	nextID := r.getAutoID()
	product.Id = nextID
	product.CreatedAt = time.Now().Unix()
	product.ModifiedAt = time.Now().Unix()
	product.Rating = 0
	r.products[nextID] = product
	return nextID
}

func (r *ProductRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewProduct(&model.Product{
		Id:         1,
		CategoryId: 2,
		Image: []model.Image{
			{ProductId: 1, ImageUrl: "/uploads/images/item-02.jpg"},
		},
		Name:       "Herschel supply co 25l",
		Price:      75,
		IsSale:     true,
		CreatedAt:  1614362898000,
		ModifiedAt: 1615410795000,
	})
	r.CreateNewProduct(&model.Product{
		Id:         2,
		CategoryId: 1,
		Image: []model.Image{
			{ProductId: 1, ImageUrl: "/uploads/images/item-03.jpg"},
		},
		Name:       "Denim jacket blue",
		Price:      92.5,
		IsSale:     false,
		CreatedAt:  1610281342000,
		ModifiedAt: 1619283693000,
	})
	r.CreateNewProduct(&model.Product{
		Id:         3,
		CategoryId: 3,
		Image: []model.Image{
			{ProductId: 1, ImageUrl: "/uploads/images/item-04.jpg"},
		},
		Name:       "Coach slim easton black",
		Price:      165.9,
		IsSale:     false,
		CreatedAt:  1615745962000,
		ModifiedAt: 1615976362000,
	})
	r.CreateNewProduct(&model.Product{
		Id:         4,
		CategoryId: 1,
		Image: []model.Image{
			{ProductId: 1, ImageUrl: "/uploads/images/item-05.jpg"},
		},
		Name:       "Frayed denim shorts",
		Price:      15.9,
		IsSale:     true,
		CreatedAt:  1615746962000,
		ModifiedAt: 1615977362000,
	})
	r.CreateNewProduct(&model.Product{
		Id:         5,
		CategoryId: 2,
		Image: []model.Image{
			{ProductId: 1, ImageUrl: "/uploads/images/item-02.jpg"},
		},
		Name:       "Herschel supply co 25l",
		Price:      75,
		IsSale:     false,
		CreatedAt:  1614362898000,
		ModifiedAt: 1615410795000,
	})
	r.CreateNewProduct(&model.Product{
		Id:         6,
		CategoryId: 1,
		Image: []model.Image{
			{ProductId: 1, ImageUrl: "/uploads/images/item-03.jpg"},
		},
		Name:       "Denim jacket blue",
		Price:      92.5,
		IsSale:     false,
		CreatedAt:  1610281342000,
		ModifiedAt: 1619283693000,
	})
	r.CreateNewProduct(&model.Product{
		Id:         7,
		CategoryId: 3,
		Image: []model.Image{
			{ProductId: 1, ImageUrl: "/uploads/images/item-04.jpg"},
		},
		Name:       "Coach slim easton black",
		Price:      165.9,
		IsSale:     false,
		CreatedAt:  1615745962000,
		ModifiedAt: 1615976362000,
	})
	r.CreateNewProduct(&model.Product{
		Id:         8,
		CategoryId: 1,
		Image: []model.Image{
			{ProductId: 1, ImageUrl: "/uploads/images/item-05.jpg"},
		},
		Name:       "Frayed denim shorts",
		Price:      15.9,
		IsSale:     false,
		CreatedAt:  1615746962000,
		ModifiedAt: 1615977362000,
	})
}

func (r *ProductRepo) GetAllProducts() map[int64]*model.Product {
	return r.products
}

func (r *ProductRepo) FindProductById(Id int64) (*model.Product, error) {
	if product, ok := r.products[Id]; ok {
		return product, nil
	} else {
		return nil, errors.New("product not found")
	}
}

func (r *ProductRepo) FindProductById2(Id int64) *model.Product {
	if product, ok := r.products[Id]; ok {
		return product
	} else {
		return nil
	}
}

func (r *ProductRepo) DeleteProductById(Id int64) error {
	if _, ok := r.products[Id]; ok {
		delete(r.products, Id)
		return nil
	} else {
		return errors.New("product not found")
	}
}

func (r *ProductRepo) UpdateProduct(product *model.Product) error {
	if _, ok := r.products[product.Id]; ok {
		r.products[product.Id] = product
		return nil
	} else {
		return errors.New("product not found")
	}
}

func (r *ProductRepo) Upsert(product *model.Product) int64 {
	if _, ok := r.products[product.Id]; ok {
		r.products[product.Id] = product
		return product.Id
	} else {
		return r.CreateNewProduct(product)
	}
}

func UpdateProductRating(product *model.Product) *model.Product {
	var sumRating float64 = 0
	countRating := 0
	reviews := Reviews.GetAllReviews()
	for _, review := range reviews {
		if review.ProductId == product.Id {
			countRating++
			sumRating += review.Rating
		}
	}
	product.Rating = float64(sumRating) / float64(countRating)
	return product
}
