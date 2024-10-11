package user_handler

import (
	"exp/internal/domain"
	"exp/internal/middlewares/tokens"
	"exp/internal/usecase/user_usecase"
	"github.com/gin-gonic/gin"
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
	RefreshToken(c *gin.Context)
}

type userHandler struct {
	userUC user_usecase.UsecaseForRepo
}

func New(uc user_usecase.UsecaseForRepo) *userHandler {
	return &userHandler{userUC: uc}
}

func (h *userHandler) FindUsers(ctx *gin.Context) {
	allUsers, err := h.userUC.FindAll()
	if err != nil {
		ctx.JSON(400, "Error finding users")
	}
	ctx.JSON(http.StatusOK, gin.H{"users": &allUsers})

}
func (h *userHandler) FindUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(404, "User not found")
	}
	idInt, _ := strconv.Atoi(id)
	//user, err := h.userUC.FindUserById(idInt)
	//if err != nil {
	//	ctx.JSON(400, "Error finding user_repo")
	//}
	//ctx.JSON(http.StatusOK, gin.H{"useR": &user})

	// Попробуйте получить пользователя из Redis
	userRedis, err := h.userUC.GetUserById(idInt)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"user_repo": &userRedis})
		return
	}
	// Если не нашли в Redis, получаем из основной БД
	user, err := h.userUC.FindUserById(idInt)
	if err != nil {
		ctx.JSON(404, "User not found in the database")
		return
	}
	// Сохраняем пользователя в Redis для последующих запросов
	_, err = h.userUC.SetUser(user)
	if err != nil {
		ctx.JSON(400, "Error setting user")
	}

	ctx.JSON(http.StatusOK, gin.H{"user_repo": &user})
}

func (h *userHandler) CreateUsers(ctx *gin.Context) {
	body := new([]domain.User)
	err := ctx.BindJSON(body)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	users, err := h.userUC.CreateUsers(body)
	if err != nil {
		ctx.JSON(400, "Error creating users")
		return
	}
	for _, user := range users {
		_, setErr := h.userUC.SetUser(user)
		if setErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving user to Redis"})
			return
		}
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
		ctx.JSON(400, "Error update user_repo")
	}
	ctx.JSON(http.StatusOK, gin.H{"user_repo - Updating": &user})
}
func (h *userHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(404, "User not found")
	}
	idInt, _ := strconv.Atoi(id)
	user, err := h.userUC.DeleteUser(idInt)
	if err != nil {
		ctx.JSON(400, "Error delete user_repo")
	}
	ctx.JSON(http.StatusOK, gin.H{user: "successfully deleted"})
}

func (h *userHandler) Registration(ctx *gin.Context) {
	body := new(domain.User)
	err := ctx.BindJSON(body)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	newBody, err := h.userUC.Registration(*body)
	if err != nil {
		ctx.JSON(400, "Error registration user")
	}
	ctx.JSON(http.StatusOK, gin.H{"newBody": newBody, "message": "Успешно создан"})

}
func (h *userHandler) Login(ctx *gin.Context) {
	body := new(domain.User)
	err := ctx.BindJSON(body)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}

	login, err := h.userUC.Login(*body)
	if err != nil {
		ctx.JSON(400, "Error login user")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"login": login})

	userId := login.Id
	tokenPair, err := tokens.GenerateTokenPair(userId)
	if err != nil {
		ctx.JSON(400, "Error generating token pair")
	}
	ctx.JSON(200, gin.H{"login": login, "tokens": tokenPair})
}

func (h *userHandler) RefreshToken(ctx *gin.Context) {
	refreshToken := ctx.PostForm("refresh_token")
	if refreshToken == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "refresh_token is required"})
		return
	}
	newTokenPair, err := tokens.RefreshToken(refreshToken)
	if err != nil {
		ctx.JSON(403, gin.H{"error": "Error refresh token"})
		return
	}
	ctx.JSON(200, newTokenPair)
}
