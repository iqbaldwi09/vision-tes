package repository

import (
	"vision-service/entity"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *entity.Post) error
	FindAll(limit, offset int) ([]entity.Post, int, error)
	FindByID(id uint) (entity.Post, error)
	Update(post *entity.Post) error
	Delete(id uint) error
}

type postRepo struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepo{db: db}
}

func (r *postRepo) Create(post *entity.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepo) FindAll(limit, offset int) ([]entity.Post, int, error) {
	var posts []entity.Post
	var total int64

	if err := r.db.Model(&entity.Post{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Limit(limit).Offset(offset).Find(&posts).Error; err != nil {
		return nil, 0, err
	}

	return posts, int(total), nil
}

func (r *postRepo) FindByID(id uint) (entity.Post, error) {
	var post entity.Post
	err := r.db.First(&post, id).Error
	return post, err
}

func (r *postRepo) Update(post *entity.Post) error {
	return r.db.Save(post).Error
}

func (r *postRepo) Delete(id uint) error {
	return r.db.Delete(&entity.Post{}, id).Error
}
