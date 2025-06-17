package entities

import "time"

type Place struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	Location    string    `json:"location"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Rating      float64   `json:"rating"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
