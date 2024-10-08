package user_repo_test

import (
	"exp/internal/domain"
	"exp/internal/repository/user_repo/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// Создание мока
	mockRepo := mocks.NewMockUserRepository(ctrl)

	users := []domain.User{
		{Id: 1, UserName: "Alice", Password: "pwd", Age: 17, Name: "AAAA"},
		{Id: 2, UserName: "Bob", Password: "2132132212vfvgtgh4tg4", Age: 27, Name: "Bob"},
	}

	// Установка ожидания для метода FindAllUsers
	mockRepo.EXPECT().FindAllUsers().Return(users, nil)

	// Вызов метода
	result, err := mockRepo.FindAllUsers()

	// Проверка результатов
	assert.NoError(t, err)
	assert.Equal(t, users, result)
}

func TestCreateUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	users := []domain.User{
		{Id: 1, UserName: "Alice", Name: "Alice", Age: 22, Password: "123456"},
	}

	// Установка ожидания для метода CreateUsers
	mockRepo.EXPECT().CreateUsers(users).Return(users, nil)

	// Вызов метода
	result, err := mockRepo.CreateUsers(users)

	// Проверка результатов
	assert.NoError(t, err)
	assert.Equal(t, users, result)
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	user := domain.User{Id: 1, UserName: "Alice"}

	// Установка ожидания для метода Login
	mockRepo.EXPECT().Login(user).Return(user, nil)

	// Вызов метода
	result, err := mockRepo.Login(user)

	// Проверка результатов
	assert.NoError(t, err)
	assert.Equal(t, user, result)
}

func TestRegistration(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	user := domain.User{Id: 1, UserName: "AliceSuperGirl", Name: "Alice", Age: 14, Password: "werowekwewerowerop2131"}

	// Установка ожидания для метода Registration
	mockRepo.EXPECT().Registration(user).Return(user.Id, nil)

	// Вызов метода
	id, err := mockRepo.Registration(user)

	// Проверка результатов
	assert.NoError(t, err)
	assert.Equal(t, user.Id, id)
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	updatedUser := domain.User{Id: 1, UserName: "AliceNewGod Updated", Age: 16, Name: "Alice", Password: "werowekwewerowerop2131"}

	// Установка ожидания для метода UpdateUser
	mockRepo.EXPECT().UpdateUser(1, updatedUser).Return(updatedUser, nil)

	// Вызов метода
	result, err := mockRepo.UpdateUser(1, updatedUser)

	// Проверка результатов
	assert.NoError(t, err)
	assert.Equal(t, updatedUser, result)
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	// Установка ожидания для метода DeleteUser
	mockRepo.EXPECT().DeleteUser(1).Return("Successfully deleted", nil)

	// Вызов метода
	message, err := mockRepo.DeleteUser(1)

	// Проверка результатов
	assert.NoError(t, err)
	assert.Equal(t, "Successfully deleted", message)
}
