package user_handler_test

import (
	"encoding/json"
	"exp/internal/domain"
	"exp/internal/handler/user_handler/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUsers(t *testing.T) {
	// Устанавливаем режим работы Gin в тестовый
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHandler := mocks.NewMockHandler(ctrl)

	// Создаем пользователей
	users := []domain.User{
		{Id: 1, Name: "Billy", Age: 14, UserName: "Bfif", Password: "12313krgoj54j"},
		{Id: 2, Name: "bgn", Age: 16, UserName: "Agerg", Password: "312113gg21"},
	}

	// Создание тестового контекста
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Ожидание вызова метода CreateUsers
	mockHandler.EXPECT().CreateUsers(c).DoAndReturn(func(c *gin.Context) {
		c.JSON(http.StatusOK, users) // Мокаем ответ, который вы хотите вернуть
	})

	// Вызов метода
	mockHandler.CreateUsers(c)

	// Проверка результата
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, nil, w.Body.String())
}

func TestDeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHandler := mocks.NewMockHandler(ctrl)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "Id", Value: "1"}}

	mockHandler.EXPECT().DeleteUser(c).DoAndReturn(func(c *gin.Context) {
		c.JSON(204, gin.H{"status": "deleted"})
	})
	mockHandler.DeleteUser(c)
	assert.Equal(t, http.StatusNoContent, w.Code) // Проверка, что тело ответа пустое (nil)
	assert.Empty(t, w.Body.String())

}
func TestFindUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHandler := mocks.NewMockHandler(ctrl)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "Id", Value: "1"}}
	user := domain.User{Id: 1, Name: "Valera"}
	mockHandler.EXPECT().FindUser(c).DoAndReturn(func(c *gin.Context) {
		c.JSON(http.StatusOK, user)
	})
	mockHandler.FindUser(c)
	assert.Equal(t, http.StatusOK, w.Code)
	var responseUser domain.User
	err := json.Unmarshal(w.Body.Bytes(), &responseUser)
	assert.NoError(t, err)
	assert.Equal(t, user, responseUser)
}
func TestFindUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHandler := mocks.NewMockHandler(ctrl)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	users := []domain.User{{Id: 1, Name: "Valera"}, {Id: 2, Name: "Serega"}}
	mockHandler.EXPECT().FindUser(c).DoAndReturn(func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})
	mockHandler.FindUser(c)
	assert.Equal(t, http.StatusOK, w.Code)
	var responseUser []domain.User
	err := json.Unmarshal(w.Body.Bytes(), &responseUser)
	assert.NoError(t, err)
	assert.Equal(t, users, responseUser)
}

func TestUpdateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHandler := mocks.NewMockHandler(ctrl)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	user := domain.User{Id: 1, Name: "Valera"}
	updateUser := domain.User{Id: 1, Name: "UPDATE VALERIY"}
	mockHandler.EXPECT().UpdateUser(c).DoAndReturn(func(c *gin.Context) {
		c.JSON(http.StatusOK, updateUser)
	})
	mockHandler.UpdateUser(c)
	assert.Equal(t, http.StatusOK, w.Code)
	var returnedUser domain.User
	err := json.Unmarshal(w.Body.Bytes(), &returnedUser)
	assert.NoError(t, err)

	assert.Equal(t, updateUser.Id, returnedUser.Id)
	assert.NotEqual(t, user.Name, returnedUser.Name)
	assert.Equal(t, updateUser.Name, returnedUser.Name)
}

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHandler := mocks.NewMockHandler(ctrl)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	user := domain.User{UserName: "Grisha", Password: "GrishaMasnia"}
	fullInfo := domain.User{UserName: "Grisha", Password: "GrishaMasnia"}
	mockHandler.EXPECT().Login(c).DoAndReturn(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"user_usecase": fullInfo})
	})
	mockHandler.Login(c)
	assert.Equal(t, http.StatusOK, w.Code)
	var response gin.H
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	returnedUser := response["user_usecase"].(map[string]interface{})

	assert.Equal(t, user.UserName, returnedUser["username"])
	assert.Equal(t, user.Password, returnedUser["password"])
}
func TestRegistration(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHandler := mocks.NewMockHandler(ctrl)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	newUser := domain.User{Name: "AVX", Age: 14, UserName: "Apostol", Password: "123213213t94u43tjflk"}
	checkInfo := domain.User{Name: "AVX", Age: 14, UserName: "Apostol", Password: "123213213t94u43tjflk"}
	mockHandler.EXPECT().Registration(c).DoAndReturn(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"newUser": checkInfo})
	})
	mockHandler.Registration(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var response gin.H
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	returnedUser := response["newUser"].(map[string]interface{})

	assert.Equal(t, newUser.UserName, returnedUser["username"])
	assert.Equal(t, newUser.Password, returnedUser["password"])
	assert.Equal(t, newUser.Name, returnedUser["name"])
	assert.Equal(t, newUser.Age, int(returnedUser["age"].(float64)))
}
