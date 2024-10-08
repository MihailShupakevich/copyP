package post_handler_test

import (
	"encoding/json"
	"exp/internal/domain"
	"exp/internal/handler/post_handler/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHandler := mocks.NewMockPostHandlerI(ctrl)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	mockHandler.EXPECT().CreatePost(c).DoAndReturn(
		func(c *gin.Context) {
			c.JSON(201, gin.H{"success": "true"})
		})

	mockHandler.CreatePost(c)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "{\"success\":\"true\"}", w.Body.String())
}
func TestGetPost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHandler := mocks.NewMockPostHandlerI(ctrl)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "IdUser", Value: "1"}}

	post := domain.Post{Title: "Burh", Content: "k498tgufgj5igh98gy4gjh4li", IdUser: 1}

	mockHandler.EXPECT().GetPost(c).DoAndReturn(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Post": post})
	})

	mockHandler.GetPost(c)

	var responseMap map[string]domain.Post
	err := json.Unmarshal(w.Body.Bytes(), &responseMap)
	assert.NoError(t, err)

	responsePost, exists := responseMap["Post"]
	assert.True(t, exists)
	assert.Equal(t, post.IdUser, responsePost.IdUser)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdatePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHandler := mocks.NewMockPostHandlerI(ctrl)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "Id", Value: "1"}}
	updatePost := domain.Post{Title: "fffff", Content: "2132131"}
	mockHandler.EXPECT().UpdatePost(c).DoAndReturn(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"updatePost": updatePost})
	})
	mockHandler.UpdatePost(c)
	assert.Equal(t, http.StatusOK, w.Code)
	var responseMap map[string]domain.Post
	err := json.Unmarshal(w.Body.Bytes(), &responseMap)
	assert.NoError(t, err)

	responsePost, exists := responseMap["updatePost"]
	assert.True(t, exists)
	assert.Equal(t, responsePost, updatePost)
}
func TestDeletePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHandler := mocks.NewMockPostHandlerI(ctrl)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "Id", Value: "1"}}
	mockHandler.EXPECT().DeletePost(c).DoAndReturn(func(c *gin.Context) {
		c.JSON(http.StatusNoContent, gin.H{"string": ""})
	})
	mockHandler.DeletePost(c)
	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "", w.Body.String())
}
