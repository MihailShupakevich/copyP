package post_usecase_test

import (
	"exp/internal/domain"
	"exp/internal/usecase/post_usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPostById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUsecase := mocks.NewMockUsecasePost(ctrl)
	postId := int(1)
	expectedPost := domain.Post{Id: postId, Title: "title for User"}
	mockUsecase.EXPECT().FindPostById(postId).Return(expectedPost, nil)

	result, err := mockUsecase.FindPostById(postId)
	assert.NoError(t, err)
	assert.Equal(t, expectedPost, result)
}

func TestCreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUsecase := mocks.NewMockUsecasePost(ctrl)
	post := domain.Post{Id: 3, Title: "title for User", Content: "content", IdUser: 23}
	mockUsecase.EXPECT().CreatePost(&post).Return(post, nil)

	result, err := mockUsecase.CreatePost(&post)
	assert.NoError(t, err)
	assert.Equal(t, post, result)

}

func TestDeletePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUsecase := mocks.NewMockUsecasePost(ctrl)

	postId := 1
	mockUser := domain.Post{Id: postId}
	mockUsecase.EXPECT().DeletePost(postId).Return(mockUser, nil)
	result, err := mockUsecase.DeletePost(postId)
	assert.NoError(t, err)
	assert.Equal(t, mockUser, result)
}

func TestUpdatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUsecase := mocks.NewMockUsecasePost(ctrl)
	postId := 1
	updatePost := domain.Post{Title: "Demian", Content: "Demon", IdUser: 20}
	mockUsecase.EXPECT().UpdatePost(postId, updatePost).Return(updatePost, nil)
	result, err := mockUsecase.UpdatePost(postId, updatePost)
	assert.NoError(t, err)
	assert.Equal(t, updatePost, result)
}
