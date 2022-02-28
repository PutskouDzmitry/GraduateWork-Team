package api

import (
	"github.com/PutskouDzmitry/GraduateWork/server/pkg/model"
	"github.com/gin-gonic/gin"
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
