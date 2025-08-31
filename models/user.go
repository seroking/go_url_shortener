package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"uniqueIndex;not null;type:varchar(50)" json:"username"`
	Email        string    `gorm:"uniqueIndex;not null;type:varchar(255)" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`

	Links []Link `gorm:"foreignKey:UserID" json:"links,omitempty"`
}
