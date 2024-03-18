package models

import "time"

type Contact struct {
	ID               int       `json:"id" gorm:"column:"`
	PhoneNumber      string    `json:"phone_number" gorm:"column:phone_number"`
	YouTubeAccount   string    `json:"youtube_account" gorm:"column:youtube_account"`
	InstagramAccount string    `json:"instagram_account" gorm:"column:instagram_account"`
	TiktokAccount    string    `json:"tiktok_account" gorm:"column:tiktok_account"`
	ImoAccount       string    `json:"imo_account" gorm:"column:imo_account"`
	Address          string    `json:"address" gorm:"column:address"`
	CreatedAt        time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (*Contact) TableName() string {
	return "contacts"
}
