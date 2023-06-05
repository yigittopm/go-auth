package model

import "time"

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Email     string     `json:"email"`
}

func (User) GetTableName() string {
	return "users"
}
