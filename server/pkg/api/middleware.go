package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userId, err := h.authService.ParseAccessToken(headerParts[1])
	if err != nil {
		refreshToken, err := c.Cookie("refreshToken")
		if err != nil {
			newErrorResponse(c, http.StatusUnauthorized, err.Error())
		}
		_, err = h.authService.ParseRefreshToken(refreshToken)
		if err != nil {
			newErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}
		id, err := strconv.Atoi(userId)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		accessToken, err := h.authService.GenerateTokenAccessToken(id, "test", "test")
		if err != nil {
			newErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"Warn": "Please, choose login for update your tokens",
		})
		c.JSON(http.StatusOK, map[string]interface{}{
			"id": accessToken,
		})
	}
	c.Set("userId", userId)
}
