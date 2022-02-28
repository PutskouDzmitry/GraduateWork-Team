package data

import (
	"gorm.io/gorm"
)

type wifiData struct {
	postgres *gorm.DB
}

type WifiData interface {
}

func NewWifiData(postgres *gorm.DB) WifiData {
	return &wifiData{postgres: postgres}
}
