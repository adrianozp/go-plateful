package entities

import "time"

type Status string

const StatusActive Status = "active"
const StatusInactive Status = "inactive"

func (s Status) String() string {
	return string(s)
}

type Reviews []Review

type Review struct {
	ID        string
	UserID    string
	PlaceID   string
	Content   string
	Rating    float64
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    Status
}

type ReviewFilter struct {
	UserID  string
	PlaceID string
	Content string
	Rating  float64
	Before  time.Time
	After   time.Time
	Status  Status
}
