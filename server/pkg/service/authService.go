package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/PutskouDzmitry/GraduateWork/server/pkg/data"
	"github.com/PutskouDzmitry/GraduateWork/server/pkg/model"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

const (
	salt            = "hjqrhjqw124617ajfhajs"
	signingKey      = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenAccessTTL  = 30 * time.Minute
	tokenRefreshTTL = 12 * time.Hour
)

type authService struct {
	user data.UserData
}

type AuthService interface {
	GenerateTokenAccessToken(id int, username string, password string) (string, error)
	GenerateTokenRefreshToken(id int, username string, password string) (string, error)
	ParseAccessToken(token string) (string, error)
	ParseRefreshToken(token string) (string, error)
	CreateUser(user model.User) (string, error)
	GeneratePasswordHash(password string) string
}

func NewAuthService(user data.UserData) AuthService {
	return &authService{user: user}
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

func (a authService) GenerateTokenAccessToken(id int, username string, password string) (string, error) {
	user, err := a.user.GetUser(id, username, a.GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenAccessTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		strconv.Itoa(user.Id),
	})
	tokenStr, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}
	err = a.user.SetAccessToken(tokenStr, strconv.Itoa(id))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (a authService) GenerateTokenRefreshToken(id int, username string, password string) (string, error) {
	_, err := a.user.GetUser(id, username, a.GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenRefreshTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		a.GeneratePasswordHash(password),
	})
	tokenStr, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}
	err = a.user.SetRefreshToken(tokenStr, strconv.Itoa(id))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (a authService) ParseAccessToken(token string) (string, error) {
	tokens, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(accessToken *jwt.Token) (interface{}, error) {
		if _, ok := accessToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid singing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := tokens.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type tokenClaims")
	}
	_, err = a.user.GetAccessToken(token, claims.UserId)
	if err != nil {
		return "", err
	}
	return claims.UserId, nil
}

func (a authService) ParseRefreshToken(token string) (string, error) {
	tokens, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid singing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := tokens.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type tokenClaims")
	}
	_, err = a.user.GetRefreshToken(token, claims.UserId)
	if err != nil {
		return "", err
	}
	return claims.UserId, nil
}

func (a authService) CreateUser(user model.User) (string, error) {
	user.Password = a.GeneratePasswordHash(user.Password)
	return a.user.CreateUser(user)
}

func (a authService) GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
