package handlers

import "github.com/adrianozp/go-plateful/app/place/entities"

type dto struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Location    string `json:"location"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func (d dto) toEntity() entities.Place {
	return entities.Place{
		Name:        d.Name,
		Address:     d.Address,
		Phone:       d.Phone,
		Email:       d.Email,
		Location:    d.Location,
		Category:    d.Category,
		Description: d.Description,
		Image:       d.Image,
	}
}
