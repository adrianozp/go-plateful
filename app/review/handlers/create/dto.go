package handlers

import (
	"github.com/adrianozp/go-plateful/app/review/entities"
)

type dto struct {
	ID      string  `json:"id"`
	UserID  string  `json:"user_id"`
	PlaceID string  `json:"place_id"`
	Content string  `json:"content"`
	Rating  float64 `json:"rating"`
}

func (d dto) toEntity() entities.Review {
	return entities.Review{
		ID:      d.ID,
		UserID:  d.UserID,
		PlaceID: d.PlaceID,
		Content: d.Content,
		Rating:  d.Rating,
	}
}
