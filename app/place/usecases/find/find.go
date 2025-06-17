package find

import (
	"context"

	"github.com/adrianozp/go-plateful/app/place/entities"
)

type PlaceRepository interface {
	Find(ctx context.Context, id string) (entities.Place, error)
}

type FindPlaceUsecase struct {
	repository PlaceRepository
}

func NewFindPlaceUsecase(repository PlaceRepository) FindPlaceUsecase {
	return FindPlaceUsecase{repository: repository}
}

func (u FindPlaceUsecase) Find(ctx context.Context, id string) (entities.Place, error) {
	return u.repository.Find(ctx, id)
}
