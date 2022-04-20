package data

import (
	"fmt"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/model"
	"gorm.io/gorm"
)

type WifiUserModels struct {
	UserModelId int
	Username    string `gorm:"username"`
	Password    string `gorm:"password"`
}

type WifiDataModel struct {
	IdUserData   int    `gorm:"id_user_data"`
	IdRouterWifi int    `gorm:"id_router_wifi"`
	Path         string `gorm:"path"`
}

type CoordinatesPoints struct {
	Id int64   `gorm:"id"`
	X  float64 `gorm:"x"`
	Y  float64 `gorm:"y"`
}

type RouterDataModel struct {
	IdRouter                  int64   `gorm:"id_router"`
	CoordinatesOfRouterID     int64   `gorm:"id_coordinates"`
	TransmitterPower          float64 `gorm:"transmitter_power"`
	GainOfTransmittingAntenna float64 `gorm:"gain_of_transmitting_antenna"`
	GainOfReceivingAntenna    float64 `gorm:"gain_of_receiving_antenna"`
	Speed                     int     `gorm:"speed"`
	SignalLossTransmitting    float64 `gorm:"signal_loss_transmitting"`
	SignalLossReceiving       float64 `gorm:"signal_loss_receiving"`
	NumberOfChannels          int     `gorm:"number_of_channels"`
	Scale                     float64 `gorm:"scale"`
	Thickness                 float64 `gorm:"thickness"`
	COM                       float64 `gorm:"com"`
}

type RouterDataModelWithOutID struct {
	CoordinatesOfRouterID     int64   `gorm:"id_coordinates"`
	TransmitterPower          float64 `gorm:"transmitter_power"`
	GainOfTransmittingAntenna float64 `gorm:"gain_of_transmitting_antenna"`
	GainOfReceivingAntenna    float64 `gorm:"gain_of_receiving_antenna"`
	Speed                     int     `gorm:"speed"`
	SignalLossTransmitting    float64 `gorm:"signal_loss_transmitting"`
	SignalLossReceiving       float64 `gorm:"signal_loss_receiving"`
	NumberOfChannels          int     `gorm:"number_of_channels"`
	Scale                     float64 `gorm:"scale"`
	Thickness                 float64 `gorm:"thickness"`
	COM                       float64 `gorm:"com"`
}

type wifiData struct {
	postgres *gorm.DB
}

type WifiData interface {
	SaveData(wifiSettings []model.RouterSettings, userId int64, filePath string) error
	GetData() error
}

func NewWifiData(postgres *gorm.DB) WifiData {
	return &wifiData{postgres: postgres}
}

func (w wifiData) SaveData(wifiSettings []model.RouterSettings, userId int64, filePath string) error {
	var newUser model.User
	result := w.postgres.Where("id=?", userId).Find(&newUser)
	if result.Error != nil && newUser.Id != 0 {
		return fmt.Errorf("user doesn't find: %w", result.Error)
	}
	err := w.addDataIntoDb(wifiSettings, userId, filePath)
	if err != nil {
		return err
	}
	return nil
}

func (w wifiData) addDataIntoDb(wifiSettings []model.RouterSettings, userId int64, filePath string) error {
	for i, value := range wifiSettings {
		coord := value.CoordinatesOfRouter
		result := w.postgres.Create(coord)
		if result.Error != nil {
			return result.Error
		}
		var coordPoints []CoordinatesPoints
		result = w.postgres.Table("coordinates_points").Where("x=? AND y=?", coord.X, coord.Y).Find(&coordPoints)
		if result.Error != nil {
			return result.Error
		}
		routerSettings := convertRouterSettingsToRouterDataModel(value, coordPoints[i])
		result = w.postgres.Table("router_data_models").Create(routerSettings)
		if result.Error != nil {
			return result.Error
		}
		var router []RouterDataModel
		result = w.postgres.Table("router_data_models").Find(&router)
		if result.Error != nil {
			return result.Error
		}

		wifiModel := createWifiModels(router[i].IdRouter, userId, filePath)
		result = w.postgres.Create(wifiModel)
		if result.Error != nil {
			return result.Error
		}
	}
	var wifiDataModel WifiDataModel
	result := w.postgres.Table("wifi_data_models").Find(&wifiDataModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func convertRouterSettingsToRouterDataModel(routers model.RouterSettings, point CoordinatesPoints) RouterDataModelWithOutID {
	return RouterDataModelWithOutID{
		CoordinatesOfRouterID:     point.Id,
		TransmitterPower:          routers.TransmitterPower,
		GainOfTransmittingAntenna: routers.GainOfTransmittingAntenna,
		GainOfReceivingAntenna:    routers.GainOfReceivingAntenna,
		Speed:                     routers.Speed,
		SignalLossTransmitting:    routers.SignalLossTransmitting,
		SignalLossReceiving:       routers.SignalLossReceiving,
		NumberOfChannels:          routers.NumberOfChannels,
		Scale:                     routers.Scale,
		Thickness:                 routers.Thickness,
		COM:                       routers.COM,
	}
}

func createWifiModels(routersID int64, userID int64, path string) WifiDataModel {
	return WifiDataModel{
		IdUserData:   int(userID),
		IdRouterWifi: int(routersID),
		Path:         path,
	}
}

func (w wifiData) GetData() error {

	return nil
}
