package controller

import (
	"go-auth/pkg-v2/domain/model"
	"go-auth/pkg-v2/usecase/usecase"
	"net/http"
)

type userController struct {
	userUsecase usecase.User
}

type User interface {
	GetUsers(ctx Context) error
}

func NewUserController(us usecase.User) User {
	return &userController{us}
}

func (uc *userController) GetUsers(ctx Context) error {
	var u []*model.User

	u, err := uc.userUsecase.List(u)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, u)
}
