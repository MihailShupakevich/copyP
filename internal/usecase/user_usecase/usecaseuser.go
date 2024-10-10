package user_usecase

import (
	"exp/internal/domain"
	"exp/internal/repository/user_repo"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gtank/crypto/bcrypt"
	"time"
)

type Usecase interface {
	FindAllUsers() ([]domain.User, error)
	FindUserById(userId int) (domain.User, error)
	CreateUsers(users []domain.User) ([]domain.User, error)
	DeleteUser(userId int) (domain.User, error)
	UpdateUser(userId int, user domain.User) (domain.User, error)
	Login(body domain.User) (domain.User, error)
	Registration(body domain.User) (string, error)
	SetUser(user domain.User) (domain.User, error)
	GetUserById(userId int) (domain.User, error)
}

type UsecaseForRepo struct {
	userRepo  user_repo.UserRepository
	redisRepo user_repo.RedisRepository
}

func New(userRepos user_repo.UserRepository, redisRepos user_repo.RedisRepository) UsecaseForRepo {
	return UsecaseForRepo{userRepo: userRepos, redisRepo: redisRepos}
}

func (u *UsecaseForRepo) FindAll() ([]domain.User, error) {
	users, err := u.userRepo.FindAllUsers()
	return users, err
}

func (u *UsecaseForRepo) FindUserById(userId int) (domain.User, error) {
	user, err := u.userRepo.FindUser(userId)
	if err != nil {
		err.Error()
	}
	return user, err
}

func (u *UsecaseForRepo) CreateUsers(users *[]domain.User) ([]domain.User, error) {
	newUsers, err := u.userRepo.CreateUsers(*users)
	if err != nil {
		err.Error()
	}
	return newUsers, err
}
func (u *UsecaseForRepo) DeleteUser(userId int) (string, error) {
	deleteUser, err := u.userRepo.DeleteUser(userId)
	if err != nil {
		err.Error()
	}
	return deleteUser, err
}
func (u *UsecaseForRepo) UpdateUser(userId int, updateUser domain.User) (domain.User, error) {
	updateUserVar, err := u.userRepo.UpdateUser(userId, updateUser)
	if err != nil {
		err.Error()
	}
	return updateUserVar, err
}
func (u *UsecaseForRepo) Login(body domain.User) (domain.User, error) {
	login, err := u.userRepo.Login(body)
	if err != nil {
		return domain.User{}, err
	}
	return login, err
}

func (u *UsecaseForRepo) Registration(body domain.User) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
	body.Password = string(hashedPassword)
	fmt.Println(body)
	id, err := u.userRepo.Registration(body)
	if err != nil {
		return "nil", err
	}
	token, _ := GenerateToken(id)
	return token, err
}

// Ge To
func GenerateToken(userId int) (string, error) {
	secretKey := []byte("your-256-bit-secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id_user": userId,
		"exp":     time.Now().Add(time.Minute * 30).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	return tokenString, err
}
func (u *UsecaseForRepo) SetUser(user domain.User) (domain.User, error) {
	newRedisUser, err := u.redisRepo.SetUser(user)
	fmt.Println("UC1")
	if err != nil {
		return domain.User{}, err
	}
	return newRedisUser, nil // Возвращаем пользователя, если все прошло успешно
}

func (u *UsecaseForRepo) GetUserById(userId int) (domain.User, error) {
	fmt.Println("UC2#1")
	user, err := u.redisRepo.GetUserById(userId)
	fmt.Println("UC2#2")
	if err != nil {
		return domain.User{}, err // Возвращаем ошибку, если не удалось получить пользователя
	}
	return user, nil // Возвращаем пользователя, если все прошло успешно
}
