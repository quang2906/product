package repository

import (
	"errors"
	"fmt"

	"example.com/product/model"
)

type CategoryRepo struct {
	categories map[int64]*model.Category
	autoID     int64
}

var Categories CategoryRepo

func init() {
	Categories = CategoryRepo{autoID: 0}
	Categories.categories = make(map[int64]*model.Category)
	Categories.InitData("sql:45312")
}

func (r *CategoryRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *CategoryRepo) CreateNewCategory(category *model.Category) int64 {
	nextID := r.getAutoID()
	category.Id = nextID
	r.categories[nextID] = category
	return nextID
}

func (r *CategoryRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewCategory(&model.Category{
		Id:   1,
		Name: "Women",
	})
	r.CreateNewCategory(&model.Category{
		Id:   2,
		Name: "Men",
	})
	r.CreateNewCategory(&model.Category{
		Id:   3,
		Name: "Kids",
	})
}

func (r *CategoryRepo) GetAllCategories() map[int64]*model.Category {
	return r.categories
}

func (r *CategoryRepo) FindCategoryById(Id int64) (*model.Category, error) {
	if category, ok := r.categories[Id]; ok {
		return category, nil
	} else {
		return nil, errors.New("category not found")
	}
}

func (r *CategoryRepo) FindCategoryById2(Id int64) *model.Category {
	if category, ok := r.categories[Id]; ok {
		return category
	} else {
		return nil
	}
}

func (r *CategoryRepo) DeleteCategoryById(Id int64) error {
	if _, ok := r.categories[Id]; ok {
		delete(r.categories, Id)
		return nil
	} else {
		return errors.New("category not found")
	}
}

func (r *CategoryRepo) UpdateCategory(category *model.Category) error {
	if _, ok := r.categories[category.Id]; ok {
		r.categories[category.Id] = category
		return nil
	} else {
		return errors.New("category not found")
	}
}

func (r *CategoryRepo) Upsert(category *model.Category) int64 {
	if _, ok := r.categories[category.Id]; ok {
		r.categories[category.Id] = category
		return category.Id
	} else {
		return r.CreateNewCategory(category)
	}
}
