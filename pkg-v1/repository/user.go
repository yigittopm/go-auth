package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-auth/pkg-v1/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) All() ([]model.User, error) {
	var users []model.User
	err := u.db.Find(&users).Error
	return users, err
}

func (u *UserRepository) FindById(id uint) (*model.User, error) {
	user := new(model.User)
	err := u.db.Where(`id = ?`, id).First(&user).Error
	return user, err
}

func (u *UserRepository) Save(user *model.User) (*model.User, error) {
	err := u.db.Save(&user).Error
	return user, err
}

func (u *UserRepository) Delete(id uint) error {
	err := u.db.Delete(&model.User{ID: id}).Error
	return err
}

func (u *UserRepository) Migrate() error {
	return u.db.AutoMigrate(&model.User{}).Error
}
