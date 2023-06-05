package registry

import (
	"github.com/jinzhu/gorm"
	"go-auth/pkg-v2/adapter/controller"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		User: r.NewUserController(),
	}
}
