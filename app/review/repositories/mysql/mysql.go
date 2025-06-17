package mysql

import (
	"context"
	"time"

	"github.com/adrianozp/go-plateful/app/review/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return ReviewRepository{db: db}
}

func (r ReviewRepository) FindByID(ctx context.Context, id string) (entities.Review, error) {
	var model review
	err := r.db.WithContext(ctx).Where("id = ?", id).Take(&model).Error
	if err != nil {
		return entities.Review{}, err
	}
	return model.ToEntity(), nil
}

func (r ReviewRepository) FindByFilter(ctx context.Context, filter entities.ReviewFilter) (entities.Reviews, error) {
	var model reviews
	tx := r.db.WithContext(ctx).Model(&review{})
	tx = createWhereClause(tx, filter)

	if err := tx.Find(&model).Error; err != nil {
		return entities.Reviews{}, err
	}
	return model.ToEntity(), nil
}

func (r ReviewRepository) Create(ctx context.Context, review entities.Review) (string, error) {
	if review.ID == "" {
		review.ID = uuid.New().String()
	}
	model := toModel(review)
	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return "", err
	}
	return review.ID, nil
}

func (r ReviewRepository) Update(ctx context.Context, id string, review entities.Review) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Updates(review).Error
}
