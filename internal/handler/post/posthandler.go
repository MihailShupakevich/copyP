package post

import (
	"exp/internal/domain"
	"exp/internal/usecase/post"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PostHandler struct {
	UC post.UsecasePost
}

type PostHandlerI interface {
	CreatePost(ctx *gin.Context)
	GetPost(ctx *gin.Context)
	DeletePost(ctx *gin.Context)
	UpdatePost(ctx *gin.Context)
}

func New(ucp post.UsecaseForRepoPost) PostHandler {
	return PostHandler{UC: &ucp}
}

func (p *PostHandler) CreatePost(c *gin.Context) {
	var newPost domain.Post
	if err := c.ShouldBind(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(newPost)
	createPost, err := p.UC.CreatePost(&newPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusCreated, gin.H{"post": createPost})
}
func (p *PostHandler) DeletePost(c *gin.Context) {
	id := c.Param("id")
	idPost, _ := strconv.Atoi(id)
	deletePost, err := p.UC.DeletePost(idPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"deletePost": deletePost})
}
func (p *PostHandler) UpdatePost(c *gin.Context) {
	id := c.Param("id")
	idPost, _ := strconv.Atoi(id)
	var updatePost domain.Post
	if err := c.ShouldBind(&updatePost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	updatePost, err := p.UC.UpdatePost(idPost, updatePost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"updatePost": updatePost})
}
func (p *PostHandler) GetPost(c *gin.Context) {
	id := c.Param("id")
	idPost, _ := strconv.Atoi(id)
	getPost, err := p.UC.FindPostById(idPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"getPost": getPost})
}
