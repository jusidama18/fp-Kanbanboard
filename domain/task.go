package domain

import "time"

type Task struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Title       string    `gorm:"notNull" json:"title"`
	Description string    `gorm:"notNull" json:"description"`
	Status      bool      `gorm:"notNull" json:"status"`
	UserID      int       `json:"user_id"`
	User        User      `json:"-"`
	CategoryID  int       `json:"category_id"`
	Category    Category  `json:"-"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
