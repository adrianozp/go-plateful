package update

import (
	"context"

	"github.com/adrianozp/go-plateful/app/place/entities"
)

type PlaceRepository interface {
	Update(ctx context.Context, id string, place entities.Place) error
}

type UpdatePlaceUsecase struct {
	repository PlaceRepository
}

func NewUpdatePlaceUsecase(repository PlaceRepository) UpdatePlaceUsecase {
	return UpdatePlaceUsecase{repository: repository}
}

func (u UpdatePlaceUsecase) Update(ctx context.Context, id string, place entities.Place) error {
	return u.repository.Update(ctx, id, place)
}
