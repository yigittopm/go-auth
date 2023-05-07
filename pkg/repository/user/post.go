package user

import (
	"github.com/jinzhu/gorm"
	"go-auth/pkg/model"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (u *Repository) All() ([]model.User, error) {
	var posts []model.User
	err := u.db.Find(&posts).Error
	return posts, err
}

func (u *Repository) FindById(id uint) (*model.User, error) {
	user := new(model.User)
	err := u.db.Where(`id = ?`, id).First(&user).Error
	return user, err
}

func (u *Repository) Save(user *model.User) (*model.User, error) {
	err := u.db.Save(&user).Error
	return user, err
}

func (u *Repository) Delete(id uint) error {
	err := u.db.Delete(&model.User{ID: id}).Error
	return err
}

func (u *Repository) Migrate() error {
	return u.db.AutoMigrate(&model.User{}).Error
}
