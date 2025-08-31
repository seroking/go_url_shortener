package models

import "time"

type Link struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ShortCode   string    `gorm:"uniqueIndex;not null;size:10" json:"short_code"`
	OriginalUrl string    `gorm:"not null" json:"original_url"`
	Clicks      int       `gorm:"default:0" json:"clicks"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UserID      uint      `gorm:"index;not null" json:"user_id"`

	User *User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"user,omitempty"`
}
