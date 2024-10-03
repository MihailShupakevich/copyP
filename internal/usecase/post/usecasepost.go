package post

import (
	"exp/internal/domain"
	"exp/internal/repository/post"
)

type UsecasePost interface {
	CreatePost(newPost *domain.Post) (domain.Post, error)
	DeletePost(idPost int) (domain.Post, error)
	UpdatePost(idPost int, newPost domain.Post) (domain.Post, error)
	FindPostById(idPost int) (domain.Post, error)
}

type UsecaseForRepoPost struct {
	postRepo post.PostRepository
}

func New(postRepos post.PostRepository) UsecaseForRepoPost {
	return UsecaseForRepoPost{postRepo: postRepos}
}

func (u *UsecaseForRepoPost) CreatePost(newPost *domain.Post) (domain.Post, error) {
	post, err := u.postRepo.CreatePost(*newPost)
	if err != nil {
		return domain.Post{}, err
	}
	return post, err
}
func (u *UsecaseForRepoPost) DeletePost(idPost int) (domain.Post, error) {
	post, err := u.postRepo.DeletePost(idPost)
	if err != nil {
		return domain.Post{}, err
	}
	return post, err
}
func (u *UsecaseForRepoPost) FindPostById(idPost int) (domain.Post, error) {
	findPost, err := u.postRepo.FindPostById(idPost)
	if err != nil {
		return domain.Post{}, err
	}
	return findPost, err
}
func (u *UsecaseForRepoPost) UpdatePost(idPost int, newPost domain.Post) (domain.Post, error) {
	findPost, err := u.postRepo.UpdatePost(idPost, newPost)
	if err != nil {
		return domain.Post{}, err
	}
	return findPost, err
}
