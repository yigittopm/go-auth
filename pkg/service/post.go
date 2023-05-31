package service

import (
	"go-auth/pkg/model"
	"go-auth/pkg/repository"
)

type PostService struct {
	PostRepository *repository.PostRepository
}

func NewPostService(p *repository.PostRepository) PostService {
	return PostService{PostRepository: p}
}

func (p *PostService) All() ([]model.Post, error) {
	return p.PostRepository.All()
}

func (p *PostService) FindById(id uint) (*model.Post, error) {
	return p.PostRepository.FindById(id)
}

func (p *PostService) Save(post *model.Post) (*model.Post, error) {
	return p.PostRepository.Save(post)
}

func (p *PostService) Delete(id uint) error {
	return p.PostRepository.Delete(id)
}

func (p *PostService) Migrate() error {
	return p.PostRepository.Migrate()
}
