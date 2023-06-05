package model

import (
	"time"
)

type Post struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
	Description string     `json:"description"`
	Image       string     `json:"image"`
	Likes       uint       `json:"likes"`
}

type PostDTO struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Likes       uint   `json:"likes"`
}

func ToPost(postDTO *PostDTO) *Post {
	return &Post{
		Description: postDTO.Description,
		Image:       postDTO.Image,
		Likes:       postDTO.Likes,
	}
}

func ToPostDTO(post *Post) *PostDTO {
	return &PostDTO{
		ID:          post.ID,
		Description: post.Description,
		Image:       post.Image,
		Likes:       post.Likes,
	}
}
