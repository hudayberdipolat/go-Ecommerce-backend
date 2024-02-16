package models

import "time"

type Category struct {
	ID           int       `json:"id"`
	CategoryName string    `json:"category_name"`
	CategorySlug string    `json:"category_slug"`
	CreatedAt    time.Time `json:"created_at"`
}

func (*Category) TableName() string {
	return "categories"
}
