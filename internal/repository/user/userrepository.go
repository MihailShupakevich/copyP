package user

import (
	"exp/internal/domain"
	"gorm.io/gorm"
	"log"
)

type UserRepository interface {
	FindAllUsers() ([]domain.User, error)
	FindUser(userId int) (domain.User, error)
	CreateUsers(users []domain.User) ([]domain.User, error)
	UpdateUser(userId int, updateUser domain.User) (domain.User, error)
	DeleteUser(userId int) (string, error)
	Login(body domain.User) (domain.User, error)
	Registration(body domain.User) (int, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func New(DaB *gorm.DB) UserRepo {
	return UserRepo{DB: DaB}
}

func (u UserRepo) FindAllUsers() ([]domain.User, error) {
	var users []domain.User
	err := u.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u UserRepo) FindUser(userId int) (domain.User, error) {
	var user domain.User
	u.DB.Preload("Posts").Model(&user).Where("id = ?", userId).Omit("password").Find(&user)
	//.Omit("Posts.id_user").First(&user, "id = ?", userId).Omit("password", "user_id")
	return user, nil
}

func (u UserRepo) CreateUsers(users []domain.User) ([]domain.User, error) {
	err := u.DB.Create(&users).Error
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

func (u UserRepo) UpdateUser(userId int, updateUser domain.User) (domain.User, error) {
	var user domain.User
	result := u.DB.First(&user, "id = ?", userId)
	u.DB.Model(&user).Updates(updateUser)
	if result.Error != nil {
		return domain.User{}, result.Error
	}
	return user, nil
}

func (u UserRepo) DeleteUser(userId int) (string, error) {
	var user domain.User
	u.DB.Delete(&user, "id=?", userId)
	return "Successfully deleted", nil
}

func (u UserRepo) Login(body domain.User) (domain.User, error) {
	var user domain.User
	u.DB.First(&user, "user_name = ? ", body.UserName)
	return user, nil
}
func (u UserRepo) Registration(body domain.User) (int, error) {
	err := u.DB.Create(&body).Error
	if err != nil {
		log.Fatal("Registration is failed")
	}
	return body.Id, nil
}
