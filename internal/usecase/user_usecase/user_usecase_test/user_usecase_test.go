package user_usecase_test

import (
	"exp/internal/domain"
	"exp/internal/usecase/user_usecase/mocks"
	_ "exp/internal/usecase/user_usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/mock"
	"testing"
)

func TestFindAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUsecase := mocks.NewMockUsecase(ctrl)
	users := []domain.User{{Id: 1, Name: "Valera", Age: 111, UserName: "chalka"}, {Id: 2, Name: "Semen", Age: 67, UserName: "ghyhgg"}}
	mockUsecase.EXPECT().FindAllUsers().Return(users, nil)

	result, err := mockUsecase.FindAllUsers()
	assert.NoError(t, err)
	assert.Equal(t, users, result)

}
func TestFindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUsecase := mocks.NewMockUsecase(ctrl)
	userId := int(1)
	expectedUser := domain.User{Id: userId, Name: "Test User"}
	mockUsecase.EXPECT().FindUserById(userId).Return(expectedUser, nil)

	result, err := mockUsecase.FindUserById(userId)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
}
func TestCreateUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUsecase := mocks.NewMockUsecase(ctrl)
	users := []domain.User{{Id: 1, Name: "Test User", Age: 11, UserName: "vjg34uij"}, {Id: 2, Name: "Test Swerq", Age: 56, UserName: "Timosha"}}
	mockUsecase.EXPECT().CreateUsers(users).Return(users, nil)

	result, err := mockUsecase.CreateUsers(users)
	assert.NoError(t, err)
	assert.Equal(t, users, result)

}
func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUsecase := mocks.NewMockUsecase(ctrl)

	userId := 1
	mockUser := domain.User{Id: userId}

	mockUsecase.EXPECT().DeleteUser(userId).Return(mockUser, nil)
	result, err := mockUsecase.DeleteUser(userId)
	assert.NoError(t, err)
	assert.Equal(t, mockUser, result)
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUsecase := mocks.NewMockUsecase(ctrl)
	userId := 1
	updateUser := domain.User{Name: "Demian", UserName: "Demon", Age: 20}
	mockUsecase.EXPECT().UpdateUser(userId, updateUser).Return(updateUser, nil)
	result, err := mockUsecase.UpdateUser(userId, updateUser)
	assert.NoError(t, err)
	assert.Equal(t, updateUser, result)
}
func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUsecase := mocks.NewMockUsecase(ctrl)

	body := domain.User{UserName: "Chalka", Password: "231djk8yhfg795y"}
	user := domain.User{Name: "Chalka", Age: 14, UserName: "Chalka", Password: "231djk8yhfg795y"}
	mockUsecase.EXPECT().Login(body).Return(user, nil)

	result, err := mockUsecase.Login(body)
	assert.NoError(t, err)
	assert.Equal(t, user, result)

}
func TestRegistration(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUsecase := mocks.NewMockUsecase(ctrl)

	body := domain.User{Name: "Mila", Age: 25, UserName: "Namie", Password: "gjh794yh4yi4g"}
	mockUsecase.EXPECT().Registration(body).Return("success", nil)
	result, err := mockUsecase.Registration(body)
	assert.NoError(t, err)
	assert.Equal(t, "success", result)

}
