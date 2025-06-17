package mysql

import (
	"context"
	"time"

	"github.com/adrianozp/go-plateful/app/place/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlaceRepository struct {
	db *gorm.DB
}

func NewPlaceRepository(db *gorm.DB) PlaceRepository {
	return PlaceRepository{db: db}
}

func (r PlaceRepository) Find(ctx context.Context, id string) (entities.Place, error) {
	var model place
	err := r.db.WithContext(ctx).Where("id = ?", id).Take(&model).Error
	if err != nil {
		return entities.Place{}, err
	}
	return model.ToEntity(), nil
}

func (r PlaceRepository) Create(ctx context.Context, place entities.Place) (string, error) {
	if place.ID == "" {
		place.ID = uuid.New().String()
	}
	model := toModel(place)
	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return "", err
	}
	return place.ID, nil
}

func (r PlaceRepository) Update(ctx context.Context, id string, place entities.Place) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Updates(place).Error
}
