package models

import "time"

type Brend struct {
	ID        int       `json:"id"`
	BrendName string    `json:"brend_name"`
	BrendSlug string    `json:"brend_slug"`
	CreatedAt time.Time `json:"created_at"`
}

func (*Brend) TableName() string {
	return "brends"
}
