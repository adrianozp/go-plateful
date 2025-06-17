package create

import (
	"context"

	"github.com/adrianozp/go-plateful/app/place/entities"
)

type PlaceRepository interface {
	Create(ctx context.Context, place entities.Place) (string, error)
}

type CreatePlaceUsecase struct {
	repository PlaceRepository
}

func NewCreatePlaceUsecase(repository PlaceRepository) CreatePlaceUsecase {
	return CreatePlaceUsecase{repository: repository}
}

func (u CreatePlaceUsecase) Create(ctx context.Context, place entities.Place) (string, error) {
	return u.repository.Create(ctx, place)
}
