package api

import (
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.signUp)
		auth.POST("/login", h.login)
		auth.GET("/refresh", h.refresh)
	}

	apiWifiMap := router.Group("/api/map") //, h.userIdentity)
	{
		apiWifiMap.POST("/calculation", h.calculationOfValues)
		apiWifiMap.POST("/save", h.saveData)
		apiWifiMap.POST("/load", h.loadData)
	}

	apiWifiUnit2 := router.Group("/api/unit2")
	{
		apiWifiUnit2.POST("/calculation")
	}

	apiWifiUnit3 := router.Group("/api/unit3")
	{
		apiWifiUnit3.Group("/calculation")
	}
	return router
}
