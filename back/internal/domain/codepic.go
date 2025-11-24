package domain

import "time"

type CodePic struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	Title     string    `gorm:"size:1024" json:"title"`
	Content   string    `json:"content"`
	Language  string    `gorm:"size:64" json:"language"`
	ViewCount int       `gorm:"default:0" json:"view_count"`
	MaxViews  int       `gorm:"default:0" json:"max_views"`
}
