package mysql

import (
	"github.com/adrianozp/go-plateful/app/review/entities"

	"gorm.io/gorm"
)

func createWhereClause(query *gorm.DB, filter entities.ReviewFilter) *gorm.DB {
	if filter.UserID != "" {
		query = query.Where("user_id = ?", filter.UserID)
	}
	if filter.PlaceID != "" {
		query = query.Where("place_id = ?", filter.PlaceID)
	}
	if filter.Content != "" {
		query = query.Where("content LIKE ?", "%"+filter.Content+"%")
	}
	if filter.Rating != 0 {
		query = query.Where("rating = ?", filter.Rating)
	}
	if !filter.Before.IsZero() {
		query = query.Where("created_at <= ?", filter.Before)
	}
	if !filter.After.IsZero() {
		query = query.Where("created_at >= ?", filter.After)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	return query
}
