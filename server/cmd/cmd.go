package main

import (
	"github.com/PutskouDzmitry/GraduateWork/server/pkg/api"
	"github.com/PutskouDzmitry/GraduateWork/server/pkg/data"
	"github.com/PutskouDzmitry/GraduateWork/server/pkg/service"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

var (
	passwordRedis    = os.Getenv("REDIS_PASSWORD")
	portRedis        = os.Getenv("REDIS_PORT")
	hostRedis        = os.Getenv("REDIS_HOST")
	hostPostgres     = os.Getenv("POSTGRES_HOST_SERVER")
	portPostgres     = os.Getenv("POSTGRES_PORT_SERVER")
	userPostgres     = os.Getenv("POSTGRES_USER_SERVER")
	dbnamePostgres   = os.Getenv("POSTGRES_DB_NAME_SERVER")
	passwordPostgres = os.Getenv("POSTGRES_PASSWORD_SERVER")
	sslmodePostgres  = os.Getenv("POSTGRES_SSLMODE_SERVER")
	portServer       = os.Getenv("SERVER_OUT_PORT")
)

func init() {
	if passwordRedis == "none" {
		passwordRedis = ""
	}
	if portRedis == "" {
		portRedis = "6379"
	}
	if hostRedis == "" {
		hostRedis = "localhost"
	}
	if hostPostgres == "" {
		hostPostgres = "localhost"
	}
	if portPostgres == "" {
		portPostgres = "5432"
	}
	if userPostgres == "" {
		userPostgres = "postgres"
	}
	if dbnamePostgres == "" {
		dbnamePostgres = "postgres"
	}
	if passwordPostgres == "" {
		passwordPostgres = "password"
	}
	if sslmodePostgres == "" {
		sslmodePostgres = "disable"
	}
	if portServer == "" {
		portServer = "8080"
	}

}

func main() {
	connPostgres, err := data.GetConnectionPostgres(hostPostgres, portPostgres, userPostgres, dbnamePostgres, passwordPostgres, sslmodePostgres)
	if err != nil {
		logrus.Fatal(err)
	}
	connRedis, err := data.GetConnectionRedis(hostRedis, portRedis, passwordRedis)
	if err != nil {
		logrus.Fatal(err)
	}
	user := data.NewUserData(connPostgres, connRedis)
	wifi := data.NewWifiData(connPostgres)
	userService := service.NewAuthService(user)
	wifiService := service.NewWifiService(wifi)
	handler := api.NewHandler(userService, wifiService)
	srv := new(Server)
	logrus.Info("Server works")
	if err := srv.Run(portServer, handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}
