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
	COM                       float64 `gorm:"com"`
}

type wifiData struct {
	postgres *gorm.DB
}

type WifiData interface {
	SaveData(wifiSettings []model.RouterSettings, userId int64, filePath string) error
	GetData(userId int64) (model.Wifi, error)
	DeleteData(userId, routerId int64) error
}

func NewWifiData(postgres *gorm.DB) WifiData {
	return &wifiData{postgres: postgres}
}

func (w wifiData) SaveData(wifiSettings []model.RouterSettings, userId int64, filePath string) error {
	var newUser model.User
	newUser1 := model.User{
		Id:       int(userId),
		Username: "dima",
		Password: "dima",
	}
	result := w.postgres.Where("id=?", userId).Find(&newUser)
	if newUser.Id == 0 {
		result = w.postgres.Create(&newUser1)
		if result.Error != nil {
			return fmt.Errorf("user doesn't find: %w", result.Error)
		}
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
		result := w.postgres.Create(&coord)
		if result.Error != nil {
			return result.Error
		}
		var coordPoints CoordinatesPoints
		result = w.postgres.Where("x=? AND y=?", coord.X, coord.Y).Find(&coordPoints)
		if result.Error != nil {
			return result.Error
		}

		routerSettings := convertRouterSettingsToRouterDataModel(value, coordPoints)
		result = w.postgres.Table("router_data_models").Create(&routerSettings)
		if result.Error != nil {
			return result.Error
		}
		router := make([]RouterDataModel, 0, 10)
		result = w.postgres.Table("router_data_models").Find(&router)
		if result.Error != nil {
			return result.Error
		}

		wifiModel := createWifiModels(router[i].IdRouter, userId, filePath)
		result = w.postgres.Create(&wifiModel)
		if result.Error != nil {
			return result.Error
		}
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

func (w wifiData) GetData(userId int64) (model.Wifi, error) {
	var newUser model.User
	result := w.postgres.Where("id=?", userId).Find(&newUser)
	if result.Error != nil && newUser.Id == 0 {
		return model.Wifi{}, fmt.Errorf("user doesn't find: %w", result.Error)
	}
	dataResult, err := w.getDataFromDb(userId)
	if err != nil {
		return model.Wifi{}, err
	}
	return dataResult, nil
}

func (w wifiData) getDataFromDb(userId int64) (model.Wifi, error) {
	var wifiDataModel []WifiDataModel
	result := w.postgres.Table("wifi_data_models").Where("id_user_data=?", userId).Find(&wifiDataModel)
	if result.Error != nil {
		return model.Wifi{}, result.Error
	}
	routers := make([]RouterDataModel, 0, 10)
	var router RouterDataModel
	for _, value := range wifiDataModel {
		result = w.postgres.Table("router_data_models").Where("id_router=?", value.IdRouterWifi).Find(&router)
		if result.Error != nil {
			return model.Wifi{}, result.Error
		}
		routers = append(routers, router)
	}
	routerSettings := make([]model.RouterSettings, 0, 10)
	for _, value := range routers {
		var coordPoint CoordinatesPoints
		result = w.postgres.Table("coordinates_points").Where("id=?", value.CoordinatesOfRouterID).Find(&coordPoint)
		if result.Error != nil {
			return model.Wifi{}, result.Error
		}
		routerSetting := convertRouterDataModelToRouterSettings(value, coordPoint)
		routerSettings = append(routerSettings, routerSetting)
	}
	return model.Wifi{
		User:   userId,
		Router: routerSettings,
		Path:   wifiDataModel[0].Path,
	}, nil
}

func convertRouterDataModelToRouterSettings(router RouterDataModel, point CoordinatesPoints) model.RouterSettings {
	return model.RouterSettings{
		CoordinatesOfRouter: model.CoordinatesPoints{
			X: point.X,
			Y: point.Y,
		},
		TransmitterPower:          router.TransmitterPower,
		GainOfTransmittingAntenna: router.GainOfTransmittingAntenna,
		GainOfReceivingAntenna:    router.GainOfReceivingAntenna,
		Speed:                     router.Speed,
		SignalLossTransmitting:    router.SignalLossTransmitting,
		SignalLossReceiving:       router.SignalLossReceiving,
		NumberOfChannels:          router.NumberOfChannels,
		Scale:                     router.Scale,
		COM:                       router.COM,
	}
}

func (w wifiData) DeleteData(userId, routerId int64) error {
	return nil
}
