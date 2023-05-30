package post

import (
	"github.com/jinzhu/gorm"
	"go-auth/pkg/model"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (p *Repository) All() ([]model.Post, error) {
	var posts []model.Post
	err := p.db.Find(&posts).Error
	return posts, err
}

func (p *Repository) FindById(id uint) (*model.Post, error) {
	post := new(model.Post)
	err := p.db.Where(`id = ?`, id).First(&post).Error
	return post, err
}

func (p *Repository) Save(post *model.Post) (*model.Post, error) {
	err := p.db.Save(&post).Error
	return post, err
}

func (p *Repository) Delete(id uint) error {
	err := p.db.Delete(&model.Post{ID: id}).Error
	return err
}

func (p *Repository) Migrate() error {
	return p.db.AutoMigrate(&model.Post{}).Error
}
