package data

import (
	"gorm.io/gorm"
)

type WifiUserModels struct {
	UserModelId int
	Username    string `gorm:"username"`
	Password    string `gorm:"password"`
}

type WifiDataModel struct {
	IdUserData   int
	UserData     WifiUserModels `gorm:"foreignKey:UserModelId;references:IdUserData"`
	IdRouterWifi int
	RouterWifi   RouterDataModel `gorm:"foreignKey:IdRouter;references:IdRouterWifi"`
	Path         string
}

type CoordinatesPoints struct {
	IdCoordinates int64
	X             float64 `json:"x"`
	Y             float64 `json:"y"`
}

type RouterDataModel struct {
	IdRouter                  int64             `gorm:"id_router"`
	CoordinatesOfRouterID     int64             `gorm:"id_coordinates"`
	CoordinatesOfRouter       CoordinatesPoints `gorm:"foreignKey:IdCoordinates;references:CoordinatesOfRouterID"`
	TransmitterPower          float64           `gorm:"transmitter_power"`
	GainOfTransmittingAntenna float64           `gorm:"gain_of_transmitting_antenna"`
	GainOfReceivingAntenna    float64           `gorm:"gain_of_receiving_antenna"`
	Speed                     int               `gorm:"speed"`
	SignalLossTransmitting    float64           `gorm:"signal_loss_transmitting"`
	SignalLossReceiving       float64           `gorm:"signal_loss_receiving"`
	NumberOfChannels          int               `gorm:"number_of_channels"`
	Scale                     float64           `gorm:"scale"`
	Thickness                 float64           `gorm:"thickness"`
	COM                       float64           `gorm:"com"`
}

type wifiData struct {
	postgres *gorm.DB
}

type WifiData interface {
	SaveData(userAndWifiSetting WifiDataModel) error
	GetData() error
}

func NewWifiData(postgres *gorm.DB) WifiData {
	return &wifiData{postgres: postgres}
}

func (w wifiData) SaveData(wifi WifiDataModel) error {
	result := w.postgres.Create(&wifi)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (w wifiData) GetData() error {

	return nil
}
