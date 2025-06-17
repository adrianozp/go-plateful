package mysql

import (
	"time"

	"github.com/adrianozp/go-plateful/app/place/entities"
)

type (
	place struct {
		ID          string
		Name        string
		Address     string
		Phone       string
		Email       string
		Location    string
		Category    string
		Description string
		Image       string
		Rating      float64
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)

func (p place) ToEntity() entities.Place {
	return entities.Place(p)
}

func toModel(p entities.Place) place {
	return place(p)
}
