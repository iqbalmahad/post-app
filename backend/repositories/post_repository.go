package repositories

import (
	"github.com/iqbalmahad/post-app.git/backend/models"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *models.Post) error
	FindAll() ([]models.Post, error)
	FindByID(id string) (*models.Post, error)
	Update(post *models.Post) error
	Delete(id string) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

func (r *postRepository) Create(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) FindAll() ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Find(&posts).Error
	return posts, err
}

func (r *postRepository) FindByID(id string) (*models.Post, error) {
	var post models.Post
	err := r.db.First(&post, "id = ?", id).Error
	return &post, err
}

func (r *postRepository) Update(post *models.Post) error {
	return r.db.Save(post).Error
}

func (r *postRepository) Delete(id string) error {
	return r.db.Delete(&models.Post{}, "id = ?", id).Error
}
