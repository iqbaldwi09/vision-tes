package usecase

import (
	"vision-service/entity"
	"vision-service/repository"
)

type PostUsecase interface {
	CreateArticle(post *entity.Post) error
	GetAllArticle(limit, offset int) ([]entity.Post, int, error)
	GetArticleByID(id uint) (entity.Post, error)
	UpdateArticle(post *entity.Post) error
	DeleteArticle(id uint) error
}

type postUsecase struct {
	repo repository.PostRepository
}

func NewPostUsecase(repo repository.PostRepository) PostUsecase {
	return &postUsecase{repo}
}

func (u *postUsecase) CreateArticle(post *entity.Post) error {
	return u.repo.Create(post)
}

func (u *postUsecase) GetAllArticle(limit, offset int) ([]entity.Post, int, error) {
	return u.repo.FindAll(limit, offset)
}

func (u *postUsecase) GetArticleByID(id uint) (entity.Post, error) {
	return u.repo.FindByID(id)
}

func (u *postUsecase) UpdateArticle(post *entity.Post) error {
	return u.repo.Update(post)
}

func (u *postUsecase) DeleteArticle(id uint) error {
	return u.repo.Delete(id)
}
