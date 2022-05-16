package api

import (
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/service"
	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "HEAD"},
		AllowHeaders: []string{"Content-Type"},
		MaxAge:       3600,
	}))

	apiNotAuth := router.Group("/")
	{
		apiNotAuth.POST("/getResult", h.calculationOfValues)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/login", h.login)
		auth.GET("/refresh", h.refresh)
		auth.GET("/google", h.loginTest)
		auth.GET("/callback", h.callback)
	}

	apiWifiMap := router.Group("/api/map") //, h.userIdentity)
	{
		apiWifiMap.POST("/calculation", h.calculationOfValues)
		//apiWifiMap.POST("/save", h.saveData)
		//apiWifiMap.POST("/load", h.loadData)
		//apiWifiMap.POST("/preload", h.preloadData)
		//apiWifiMap.POST("/getInfo", h.getInfo)
		apiWifiMap.POST("/fluxMigrator", h.fluxMigrator)
		apiWifiMap.POST("/acrylicMigrator", h.acrylicMigrator)
		apiWifiMap.POST("/telephoneMigrator", h.telephoneMigrator)
	}

	return router
}
