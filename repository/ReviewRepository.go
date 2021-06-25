package repository

import (
	"errors"
	"fmt"
	"time"

	"example.com/product/model"
)

type ReviewRepo struct {
	reviews map[int64]*model.Review
	autoID  int64
}

var Reviews ReviewRepo

func init() {
	Reviews = ReviewRepo{autoID: 0}
	Reviews.reviews = make(map[int64]*model.Review)
	Reviews.InitData("sql:45312")
}

func (r *ReviewRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *ReviewRepo) CreateNewReview(review *model.Review) int64 {
	nextID := r.getAutoID()
	review.Id = nextID
	review.CreatedAt = time.Now().Unix()
	review.ModifiedAt = time.Now().Unix()
	r.reviews[nextID] = review
	return nextID
}

func (r *ReviewRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewReview(&model.Review{
		Id:         1,
		ProductId:  2,
		Rating:     3,
		CreatedAt:  1614362898000,
		ModifiedAt: 1615410795000,
	})
	r.CreateNewReview(&model.Review{
		Id:         2,
		ProductId:  2,
		Rating:     4,
		CreatedAt:  1610281342000,
		ModifiedAt: 1619283693000,
	})
}

func (r *ReviewRepo) GetAllReviews() map[int64]*model.Review {
	return r.reviews
}

func (r *ReviewRepo) FindReviewById(Id int64) (*model.Review, error) {
	if review, ok := r.reviews[Id]; ok {
		return review, nil
	} else {
		return nil, errors.New("review not found")
	}
}

func (r *ReviewRepo) FindReviewById2(Id int64) *model.Review {
	if review, ok := r.reviews[Id]; ok {
		return review
	} else {
		return nil
	}
}

func (r *ReviewRepo) DeleteReviewById(Id int64) error {
	if _, ok := r.reviews[Id]; ok {
		delete(r.reviews, Id)
		return nil
	} else {
		return errors.New("review not found")
	}
}

func (r *ReviewRepo) UpdateReview(review *model.Review) error {
	if _, ok := r.reviews[review.Id]; ok {
		r.reviews[review.Id] = review
		return nil
	} else {
		return errors.New("review not found")
	}
}

func (r *ReviewRepo) Upsert(review *model.Review) int64 {
	if _, ok := r.reviews[review.Id]; ok {
		r.reviews[review.Id] = review
		return review.Id
	} else {
		return r.CreateNewReview(review)
	}
}
