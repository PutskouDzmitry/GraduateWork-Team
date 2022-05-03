package data

import (
	"fmt"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/model"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type userData struct {
	postgres *gorm.DB
	redis    *redis.Client
}

type UserData interface {
	SetAccessToken(token string, id string) error
	GetAccessToken(token string, id string) (string, error)
	SetRefreshToken(token string, id string) error
	GetRefreshToken(token string, id string) (string, error)
	CreateUser(dataUser model.User) (string, error)
	GetUser(id int, username string, password string) (model.User, error)
}

func NewUserData(postgres *gorm.DB, redis *redis.Client) UserData {
	return &userData{
		postgres: postgres,
		redis:    redis,
	}
}

func (a userData) SetAccessToken(token string, id string) error {
	err := a.redis.Set(fmt.Sprintf("Access_token-%v", id), token, time.Minute*30).Err()
	if err != nil {
		return err
	}
	return nil
}

func (a userData) GetAccessToken(token string, id string) (string, error) {
	val, err := a.redis.Get(fmt.Sprintf("Access_token-%v", id)).Result()
	if err != nil {
		return "", err
	}
	if val != token {
		return "", fmt.Errorf("your token doesn't equal to original token")
	}
	return token, err
}

func (a userData) SetRefreshToken(token string, id string) error {
	err := a.redis.Set(fmt.Sprintf("Refresh_token-%v", id), token, time.Hour*12).Err()
	if err != nil {
		return err
	}
	return nil
}

func (a userData) GetRefreshToken(token string, id string) (string, error) {
	val, err := a.redis.Get(fmt.Sprintf("Refresh_token-%v", id)).Result()
	if err != nil {
		return "", err
	}
	if val != token {
		return "", fmt.Errorf("your token doesn't equal to original token")
	}
	return token, err
}

func (a userData) CreateUser(dataUser model.User) (string, error) {
	result := a.postgres.Create(&dataUser)
	if result.Error != nil {
		return "", result.Error
	}
	var newUser model.User
	result = a.postgres.Where("username=? and password=?", dataUser.Username, dataUser.Password).Find(&newUser)
	if result.Error != nil {
		return "", result.Error
	}
	logrus.Info("id: ", newUser.Id)
	return strconv.Itoa(newUser.Id), nil
}

func (a userData) GetUser(id int, username string, password string) (model.User, error) {
	var newUser model.User
	result := a.postgres.Where("id=?", id).Find(&newUser)
	if result.Error != nil {
		return newUser, result.Error
	}
	return newUser, nil
}
