package user_repo

import (
	"context"
	"encoding/json"
	"exp/internal/domain"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gtank/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

type UserRepository interface {
	FindAllUsers() ([]domain.User, error)
	FindUser(userId int) (domain.User, error)
	CreateUsers(users []domain.User) ([]domain.User, error)
	UpdateUser(userId int, updateUser domain.User) (domain.User, error)
	DeleteUser(userId int) (string, error)
	Login(body domain.User) (domain.User, error)
	Registration(body domain.User) (domain.User, error)
}

type RedisRepository interface {
	GetUserById(id int) (domain.User, error)
	SetUser(user domain.User) (domain.User, error)
}

type UserRepo struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

func New(db *gorm.DB, redisClient *redis.Client) *UserRepo {
	return &UserRepo{DB: db, RedisClient: redisClient}
}

func (u *UserRepo) FindAllUsers() ([]domain.User, error) {
	var users []domain.User
	err := u.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepo) FindUser(userId int) (domain.User, error) {
	var user domain.User
	if err := u.DB.Preload("Posts").First(&user, userId).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (u *UserRepo) CreateUsers(users []domain.User) ([]domain.User, error) {
	err := u.DB.Create(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepo) UpdateUser(userId int, updateUser domain.User) (domain.User, error) {
	var user domain.User
	if err := u.DB.First(&user, "id = ?", userId).Error; err != nil {
		return domain.User{}, err
	}
	if err := u.DB.Model(&user).Updates(updateUser).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (u *UserRepo) DeleteUser(userId int) (string, error) {
	var user domain.User
	err := u.DB.Delete(&user, userId).Error
	if err != nil {
		return "", err
	}
	ctx := context.Background()
	u.RedisClient.Del(ctx, strconv.Itoa(userId))

	return "Successfully deleted", nil
}

func (u *UserRepo) Login(body domain.User) (domain.User, error) {
	var user domain.User
	if err := u.DB.First(&user, "user_name = ?", body.UserName).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, nil // Пользователь не найден
		}
		return domain.User{}, err // Возврат ошибки, если ошибка не связана с отсутствием записи
	}
	// Проверка пароля
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) // Проверяем хэш пароля
	if err != nil {
		return domain.User{}, nil // Неверный пароль
	}
	return user, nil // Возвращаем найденного пользователя
}

func (u *UserRepo) Registration(body domain.User) (domain.User, error) {
	err := u.DB.Create(&body).Error
	if err != nil {
		log.Fatal("Registration is failed:", err)
		return domain.User{}, err
	}
	return body, nil
}

// Реализация интерфейса RedisRepository
func (u *UserRepo) GetUserById(id int) (domain.User, error) {
	ctx := context.Background()
	// Получение пользователя из кэша Redis
	cachedUser, err := u.RedisClient.Get(ctx, strconv.Itoa(id)).Result()
	if err == redis.Nil {
		// Если ключ не найден, возвращаем ошибку или nil
		return domain.User{}, nil // или return domain.User{}, errors.New("user not found in cache")
	} else if err != nil {
		// Если произошла другая ошибка при доступе к Redis
		return domain.User{}, err
	}

	// Десериализация данных из JSON в структуру domain.User
	var user domain.User
	if err := json.Unmarshal([]byte(cachedUser), &user); err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (u *UserRepo) SetUser(user domain.User) (domain.User, error) {
	ctx := context.Background()
	userJSON, err := json.Marshal(user)
	if err != nil {
		return domain.User{}, err
	}
	u.RedisClient.Set(ctx, strconv.Itoa(user.Id), userJSON, time.Hour)
	fmt.Println("R2")
	return user, nil
}
