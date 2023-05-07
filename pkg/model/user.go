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

type UserDTO struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func ToUser(userDTO *UserDTO) *User {
	return &User{
		Username: userDTO.Username,
		Password: userDTO.Password,
		Email:    userDTO.Email,
	}
}

func ToUserDTO(user *User) *UserDTO {
	return &UserDTO{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}
}
