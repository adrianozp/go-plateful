package create

import (
	"context"

	"github.com/adrianozp/go-plateful/app/review/entities"
)

type ReviewRepository interface {
	Create(ctx context.Context, review entities.Review) (string, error)
}

type CreateReviewUsecase struct {
	repository ReviewRepository
}

func NewCreateReviewUsecase(repository ReviewRepository) CreateReviewUsecase {
	return CreateReviewUsecase{repository: repository}
}

func (u CreateReviewUsecase) Create(ctx context.Context, review entities.Review) (string, error) {
	return u.repository.Create(ctx, review)
}
