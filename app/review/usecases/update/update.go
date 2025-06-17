package update

import (
	"context"

	"github.com/adrianozp/go-plateful/app/review/entities"
)

type ReviewRepository interface {
	Update(ctx context.Context, id string, review entities.Review) error
}

type UpdateReviewUsecase struct {
	repository ReviewRepository
}

func NewUpdateReviewUsecase(repository ReviewRepository) UpdateReviewUsecase {
	return UpdateReviewUsecase{repository: repository}
}

func (u UpdateReviewUsecase) Update(ctx context.Context, id string, review entities.Review) error {
	return u.repository.Update(ctx, id, review)
}
