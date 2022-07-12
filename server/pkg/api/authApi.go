package api

import (
	"fmt"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/model"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

func (h Handler) signUp(c *gin.Context) {
	var input model.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.authService.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h Handler) login(c *gin.Context) {
	input, err := getLoginAndPassword(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	accessToken, err := h.authService.GenerateTokenAccessToken(input.Id, input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	refreshToken, err := h.authService.GenerateTokenRefreshToken(input.Id, input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.SetCookie("refreshToken", refreshToken, 20, "/", "localhost", true, true)
	redirectResponse(c, "http://localhost:3000/api/map/home", accessToken)
}

func getLoginAndPassword(c *gin.Context) (model.User, error) {
	username := c.Request.FormValue("login")
	if username == "" {
		return model.User{}, fmt.Errorf("login is empty")
	}
	password := c.Request.FormValue("password")
	if password == "" {
		return model.User{}, fmt.Errorf("password is empty")
	}
	return model.User{
		Id:       generateId(),
		Username: username,
		Password: password,
	}, nil
}

func generateId() int {
	return rand.Intn(13000)
}

func (h Handler) refresh(c *gin.Context) {
	var input model.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	accessToken, err := h.authService.GenerateTokenAccessToken(input.Id, input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	refreshToken, err := h.authService.GenerateTokenRefreshToken(input.Id, input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.SetCookie("refreshToken", refreshToken, 20, "/", "localhost", true, true)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": accessToken,
	})
}
