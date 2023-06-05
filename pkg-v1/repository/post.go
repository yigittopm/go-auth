package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-auth/pkg-v1/model"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (p *PostRepository) All() ([]model.Post, error) {
	var posts []model.Post
	err := p.db.Find(&posts).Error
	return posts, err
}

func (p *PostRepository) FindById(id uint) (*model.Post, error) {
	post := new(model.Post)
	err := p.db.Where(`id = ?`, id).First(&post).Error
	return post, err
}

func (p *PostRepository) Save(post *model.Post) (*model.Post, error) {
	err := p.db.Save(&post).Error
	return post, err
}

func (p *PostRepository) Delete(id uint) error {
	err := p.db.Delete(&model.Post{ID: id}).Error
	return err
}

func (p *PostRepository) Migrate() error {
	return p.db.AutoMigrate(&model.Post{}).Error
}
