package user

import (
	"exp/internal/domain"
	"exp/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Handler interface {
	FindUsers(c *gin.Context)
	FindUser(c *gin.Context)
	CreateUsers(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	Registration(c *gin.Context)
	Login(c *gin.Context)
}

type userHandler struct {
	userUC user.UsecaseForRepo
}

func New(uc user.UsecaseForRepo) *userHandler {
	return &userHandler{userUC: uc}
}

func (h *userHandler) FindUsers(ctx *gin.Context) {
	allUsers, err := h.userUC.FindAll()
	if err != nil {
		log.Fatal("Error finding users")
	}
	ctx.JSON(http.StatusOK, gin.H{"users": &allUsers})

}
func (h *userHandler) FindUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(404, "User not found")
	}
	idInt, _ := strconv.Atoi(id)
	user, err := h.userUC.FindUserById(idInt)
	if err != nil {
		log.Fatal("Error finding user")
	}
	ctx.JSON(http.StatusOK, gin.H{"user": &user})
}

func (h *userHandler) CreateUsers(ctx *gin.Context) {
	body := new([]domain.User)
	err := ctx.BindJSON(body)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	users, err := h.userUC.CreateUsers(body)
	if err != nil {
		log.Fatal("Error creating users")
	}
	ctx.JSON(http.StatusOK, gin.H{"users": &users})
}

func (h *userHandler) UpdateUser(ctx *gin.Context) {
	updateUser := new(domain.User)
	err := ctx.BindJSON(updateUser)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(404, "User not found")
	}
	idInt, _ := strconv.Atoi(id)
	user, err := h.userUC.UpdateUser(idInt, *updateUser)
	if err != nil {
		log.Fatal("Error update user")
	}
	ctx.JSON(http.StatusOK, gin.H{"user - Updating": &user})
}
func (h *userHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(404, "User not found")
	}
	idInt, _ := strconv.Atoi(id)
	user, err := h.userUC.DeleteUser(idInt)
	if err != nil {
		log.Fatal("Error delete user")
	}
	ctx.JSON(http.StatusOK, gin.H{user: "successfully deleted"})
}

func (h *userHandler) Registration(ctx *gin.Context) {
	body := new(domain.User)
	err := ctx.BindJSON(body)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	token, err := h.userUC.Registration(*body)
	if err != nil {
		log.Fatal("Error registration user")
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token, "message": "Успешно сформирован"})

}
func (h *userHandler) Login(ctx *gin.Context) {
	body := new(domain.User)
	err := ctx.BindJSON(body)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	login, err := h.userUC.Login(*body)
	if err != nil {
		log.Fatal("Error login user")
	}
	ctx.JSON(http.StatusOK, gin.H{"login": login})
}
