package find

import (
	"context"

	"github.com/adrianozp/go-plateful/app/review/entities"
)

type ReviewRepository interface {
	FindByID(ctx context.Context, id string) (entities.Review, error)
	FindByFilter(ctx context.Context, filter entities.ReviewFilter) (entities.Reviews, error)
}

type FindReviewUsecase struct {
	repository ReviewRepository
}

func NewFindReviewUsecase(repository ReviewRepository) FindReviewUsecase {
	return FindReviewUsecase{repository: repository}
}

func (u FindReviewUsecase) FindByID(ctx context.Context, id string) (entities.Review, error) {
	return u.repository.FindByID(ctx, id)
}

func (u FindReviewUsecase) FindByFilter(ctx context.Context, filter entities.ReviewFilter) (entities.Reviews, error) {
	if filter.Status == "" {
		filter.Status = entities.StatusActive
	}
	return u.repository.FindByFilter(ctx, filter)
}
