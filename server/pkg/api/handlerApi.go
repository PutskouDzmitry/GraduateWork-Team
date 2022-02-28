package api

import (
	"github.com/PutskouDzmitry/GraduateWork/server/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService service.AuthService
	wifiService service.WifiService
}

func NewHandler(authService service.AuthService, wifiService service.WifiService) *Handler {
	return &Handler{
		authService: authService,
		wifiService: wifiService,
	}
}

func (h Handler) InitRoutes() *gin.Engine {

	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.signUp)
		auth.POST("/login", h.login)
		auth.GET("/refresh", h.refresh)
	}

	apiWifi := router.Group("/api/wifi", h.userIdentity)
	{
		apiWifi.GET("/", h.calculationOfValues)
		apiWifi.POST("/", h.saveData)
		apiWifi.POST("/", h.loadData)
	}
	return router
}
