package model

type User struct {
	Id       int    `json:"id"`
	Username string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
