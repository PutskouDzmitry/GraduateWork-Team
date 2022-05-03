package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, ErrorMessage{message})
}

func redirectResponse(c *gin.Context, message string, url string, token string) {
	c.Writer.Header().Set("Authorization", "Bearer "+token)
	c.Redirect(http.StatusMovedPermanently, url)
}
