package repositories

import (
	"github.com/RosaLeaveMeAlone/gin-crud/src/database/models"
	"github.com/RosaLeaveMeAlone/gin-crud/src/plugins"
)

type PostRepository struct{}

// NewPostRepository crea una nueva instancia del repositorio.
func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

func (r *PostRepository) Create(post *models.Post) error {
	return plugins.DB.Create(post).Error
}

func (r *PostRepository) Update(post *models.Post) error {
	return plugins.DB.Save(post).Error
}

func (r *PostRepository) FindAll() ([]models.Post, error) {
	var posts []models.Post
	err := plugins.DB.Find(&posts).Error
	return posts, err
}

func (r *PostRepository) FindByID(id uint) (*models.Post, error) {
	var post models.Post
	err := plugins.DB.First(&post, id).Error
	return &post, err
}

func (r *PostRepository) Delete(id uint) error {
	return plugins.DB.Delete(&models.Post{}, id).Error
}
