package post_repo

import (
	"exp/internal/domain"
	"fmt"
	"gorm.io/gorm"
)

type PostRepositoryI interface {
	FindPostById(idPost int) (domain.Post, error)
	CreatePost(newPost *domain.Post) (domain.Post, error)
	UpdatePost(idPost int, updatePost domain.Post) (domain.Post, error)
	DeletePost(idPost int) (domain.Post, error)
}

type PostRepository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) PostRepository {
	return PostRepository{DB: db}
}

func (r *PostRepository) FindPostById(idPost int) (domain.Post, error) {
	var post domain.Post
	if err := r.DB.First(&post, "id = ?", idPost).Error; err != nil {
		return domain.Post{}, err
	}
	return post, nil
}

func (r *PostRepository) CreatePost(newPost domain.Post) (domain.Post, error) {
	fmt.Println(newPost)
	nPost := r.DB.Create(&newPost)
	if nPost.Error != nil {
		return domain.Post{}, nPost.Error
	}
	return newPost, nil
}
func (r *PostRepository) UpdatePost(idPost int, updatePost domain.Post) (domain.Post, error) {
	var uPost domain.Post
	if err := r.DB.First(&uPost, "id = ?", idPost).Error; err != nil {
		return domain.Post{}, err
	}
	r.DB.Model(&uPost).Updates(updatePost)
	return updatePost, nil
}
func (r *PostRepository) DeletePost(idPost int) (domain.Post, error) {
	var deletePost domain.Post
	if err := r.DB.Delete(&deletePost, "id = ?", idPost).Error; err != nil {
		return domain.Post{}, err
	}
	return deletePost, nil
}
