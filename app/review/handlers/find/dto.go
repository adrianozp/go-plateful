package handlers

import (
	"time"

	"github.com/adrianozp/go-plateful/app/review/entities"
)

type dto struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	PlaceID   string    `json:"place_id"`
	Content   string    `json:"content"`
	Rating    float64   `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func toReviewsDTO(rs entities.Reviews) []dto {
	dtos := make([]dto, len(rs))
	for i, r := range rs {
		dtos[i] = toDTO(r)
	}
	return dtos
}

func toDTO(d entities.Review) dto {
	return dto{
		ID:        d.ID,
		UserID:    d.UserID,
		PlaceID:   d.PlaceID,
		Content:   d.Content,
		Rating:    d.Rating,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

type filterDTO struct {
	UserID  string    `form:"user_id"`
	PlaceID string    `form:"place_id"`
	Content string    `form:"content"`
	Rating  float64   `form:"rating"`
	Before  time.Time `form:"before"`
	After   time.Time `form:"after"`
	Status  string    `form:"status"`
}

func (f filterDTO) toEntity() entities.ReviewFilter {
	return entities.ReviewFilter{
		UserID:  f.UserID,
		PlaceID: f.PlaceID,
		Content: f.Content,
		Rating:  f.Rating,
		Before:  f.Before,
		After:   f.After,
		Status:  entities.Status(f.Status),
	}
}
