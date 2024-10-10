package user_repo

import (
	"context"
	"encoding/json"
	"exp/internal/domain"
	"fmt"
	"github.com/go-redis/redis/v8"
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
	Registration(body domain.User) (int, error)
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
	//ctx := context.Background()
	//
	//// Попытка получить пользователя из кеша
	//cachedUser, err := u.RedisClient.Get(ctx, strconv.Itoa(userId)).Result()
	//if err == nil {
	//	// Если пользователь найден в кэше, десериализуем его
	//	if err := json.Unmarshal([]byte(cachedUser), &user); err == nil {
	//		return user, nil
	//	}
	//}

	// Если пользователь не найден в кэше, получаем из базы данных
	if err := u.DB.Preload("Posts").First(&user, userId).Error; err != nil {
		return domain.User{}, err
	}

	//// Кэшируем пользователя в Redis
	//userJSON, err := json.Marshal(user)
	//if err == nil {
	//	u.RedisClient.Set(ctx, strconv.Itoa(userId), userJSON, time.Hour)
	//}

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

	// Обновление пользователя в базе данных
	if err := u.DB.Model(&user).Updates(updateUser).Error; err != nil {
		return domain.User{}, err
	}
	//
	//// Обновление кеша
	//if err := u.SetUser(user); err != nil {
	//	return domain.User{}, err
	//}

	return user, nil
}

func (u *UserRepo) DeleteUser(userId int) (string, error) {
	var user domain.User
	err := u.DB.Delete(&user, userId).Error
	if err != nil {
		return "", err
	}

	// Удаление пользователя из кеша
	ctx := context.Background()
	u.RedisClient.Del(ctx, strconv.Itoa(userId))

	return "Successfully deleted", nil
}

func (u *UserRepo) Login(body domain.User) (domain.User, error) {
	var user domain.User
	if err := u.DB.First(&user, "user_name = ?", body.UserName).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (u *UserRepo) Registration(body domain.User) (int, error) {
	err := u.DB.Create(&body).Error
	if err != nil {
		log.Fatal("Registration is failed:", err)
		return 0, err
	}
	return body.Id, nil
}

// Реализация интерфейса RedisRepository
func (u *UserRepo) GetUserById(id int) (domain.User, error) {
	fmt.Println("REpo1")
	ctx := context.Background()

	// Получение пользователя из кэша Redis
	cachedUser, err := u.RedisClient.Get(ctx, strconv.Itoa(id)).Result()
	fmt.Println("REpo2")
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
	fmt.Println("Repo3")

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
