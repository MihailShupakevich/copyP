package post_repo_test

import (
	"exp/internal/domain"
	"exp/internal/repository/post_repo/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPostById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockPostRepositoryI(ctrl)
	postId := 1
	post := domain.Post{Id: 1, Title: "Alice", Content: "pwd", IdUser: 17}

	mockRepo.EXPECT().FindPostById(postId).Return(post, nil)

	result, err := mockRepo.FindPostById(postId)

	assert.NoError(t, err)
	assert.Equal(t, post, result)
}

func TestCreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockPostRepositoryI(ctrl)
	post := domain.Post{Id: 1, Title: "Alice", Content: "Alice", IdUser: 22}
	mockRepo.EXPECT().CreatePost(&post).Return(post, nil)
	result, err := mockRepo.CreatePost(&post)
	assert.NoError(t, err)
	assert.Equal(t, post, result)
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockPostRepositoryI(ctrl)
	updatedPost := domain.Post{Id: 1, Title: "AliceNewGod Updated", IdUser: 16, Content: "Alice"}
	mockRepo.EXPECT().UpdatePost(1, updatedPost).Return(updatedPost, nil)
	result, err := mockRepo.UpdatePost(1, updatedPost)
	assert.NoError(t, err)
	assert.Equal(t, updatedPost, result)
}

func TestDeletePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockPostRepositoryI(ctrl)
	postId := 1
	postDeleted := domain.Post{Id: 1, Title: "Alice", IdUser: 16, Content: "Alice"}
	mockRepo.EXPECT().DeletePost(postId).Return(postDeleted, nil)
	message, err := mockRepo.DeletePost(postId)
	assert.NoError(t, err)
	assert.Equal(t, postDeleted, message)
}
