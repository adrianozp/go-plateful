package mysql

import (
	"time"

	"github.com/adrianozp/go-plateful/app/review/entities"
)

type (
	reviews []review
	review  struct {
		ID        string
		UserID    string
		PlaceID   string
		Content   string
		Rating    float64
		CreatedAt time.Time
		UpdatedAt time.Time
		Status    string
	}
)

func (r review) ToEntity() entities.Review {
	return entities.Review{
		ID:        r.ID,
		UserID:    r.UserID,
		PlaceID:   r.PlaceID,
		Content:   r.Content,
		Rating:    r.Rating,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
		Status:    entities.Status(r.Status),
	}
}

func (rs reviews) ToEntity() entities.Reviews {
	reviews := make(entities.Reviews, len(rs))
	for i, r := range rs {
		reviews[i] = r.ToEntity()
	}
	return reviews
}

func toModel(p entities.Review) review {
	return review{
		ID:      p.ID,
		UserID:  p.UserID,
		PlaceID: p.PlaceID,
		Content: p.Content,
		Rating:  p.Rating,
		Status:  p.Status.String(),
	}
}
