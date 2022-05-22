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
	Id           int64 `gorm:"id"`
	IdUserData   int   `gorm:"id_user_data"`
	IdRouterWifi int   `gorm:"id_router_wifi"`
	IdFilePath   int   `gorm:"file_path"`
}

type FilePath struct {
	Id         int64  `gorm:"id"`
	PathInput  string `gorm:"path_input"`
	PathOutput string `gorm:"path_output"`
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
	SaveData(wifiSettings []model.RouterSettings, userId int64, pathInput, pathOutput string) error
	GetData(userId int64) ([]model.Wifi, error)
	DeleteData(userId, routerId int64) error
}

func NewWifiData(postgres *gorm.DB) WifiData {
	return &wifiData{postgres: postgres}
}

func (w wifiData) SaveData(wifiSettings []model.RouterSettings, userId int64, pathInput, pathOutput string) error {
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
	err := w.addDataIntoDb(wifiSettings, userId, pathInput, pathOutput)
	if err != nil {
		return err
	}
	return nil
}

func generateObjectFilePath(input, output string) *FilePath {
	return &FilePath{
		Id:         0,
		PathInput:  input,
		PathOutput: output,
	}
}

func (w wifiData) addDataIntoDb(wifiSettings []model.RouterSettings, userId int64, pathInput, pathOutput string) error {
	var coord []CoordinatesPoints
	w.postgres.Find(&coord)
	for _, value := range coord {
		w.postgres.Where("id=?", value.Id).Delete(&coord)
	}

	var routers []RouterDataModel
	w.postgres.Find(&routers)
	for _, value := range routers {
		w.postgres.Where("id_router=?", value.IdRouter).Delete(&routers)
	}

	var user []WifiUserModels
	w.postgres.Find(&user)
	for _, value := range user {
		w.postgres.Where("id=?", value.UserModelId).Delete(&user)
	}

	var path []FilePath
	w.postgres.Find(&path)
	for _, value := range path {
		w.postgres.Table("file_path").Where("id=?", value.Id).Delete(&path)
	}

	var wifi []WifiDataModel
	w.postgres.Find(&wifi)
	for _, value := range wifi {
		w.postgres.Table("wifi_data_models").Where("id=?", value.Id).Delete(&wifi)
	}
	for _, value := range wifiSettings {
		coord := value.CoordinatesOfRouter

		var coordPoints CoordinatesPoints
		var coordPointsCheck CoordinatesPoints
		result := w.postgres.Where("x=? AND y=?", coord.X, coord.Y).Find(&coordPointsCheck)
		if result.Error != nil {
			return result.Error
		}
		if coordPointsCheck.Id == 0 {
			result = w.postgres.Create(&coord)
			if result.Error != nil {
				return result.Error
			}
			var getCoordPoints CoordinatesPoints
			result = w.postgres.Where("x=? AND y=?", coord.X, coord.Y).Find(&getCoordPoints)
			if result.Error != nil {
				return result.Error
			}
			coordPoints = getCoordPoints
		} else {
			coordPoints = coordPointsCheck
		}

		var routerCheck RouterDataModel
		var router RouterDataModel
		result = w.postgres.Table("router_data_models").Where("transmitter_power=? AND gain_of_transmitting_antenna=?", value.TransmitterPower, value.GainOfTransmittingAntenna).Find(&routerCheck)
		if result.Error != nil {
			return result.Error
		}
		if routerCheck.IdRouter == 0 {
			routerSettings := convertRouterSettingsToRouterDataModel(value, coordPoints)
			result = w.postgres.Table("router_data_models").Create(&routerSettings)
			if result.Error != nil {
				return result.Error
			}
			result = w.postgres.Table("router_data_models").Where("transmitter_power=? AND gain_of_transmitting_antenna=? AND gain_of_receiving_antenna=?", routerSettings.TransmitterPower, routerSettings.GainOfTransmittingAntenna, routerSettings.GainOfReceivingAntenna).Find(&router)
			if result.Error != nil {
				return result.Error
			}
		} else {
			router = routerCheck
		}

		var pathCheck FilePath
		result = w.postgres.Table("file_path").Where("path_input=?", pathInput).Find(&pathCheck)
		if pathCheck.Id == 0 {
			result = w.postgres.Table("file_path").Create(generateObjectFilePath(pathInput, pathOutput))
			if result.Error != nil {
				return result.Error
			}
			pathCheck = FilePath{}
			result = w.postgres.Table("file_path").Where("path_input=?", pathInput).Find(&pathCheck)
		}

		var wifiCheck WifiDataModel
		wifiModel := createWifiModels(router.IdRouter, userId, pathCheck.Id)

		result = w.postgres.Where("id_router_wifi=?", routerCheck.IdRouter).Find(&wifiCheck)
		if result.Error != nil {
			return result.Error
		}
		if wifiCheck.IdUserData == 0 {
			result = w.postgres.Create(&wifiModel)
			if result.Error != nil {
				return result.Error
			}
		}
	}
	//var routerCheck []RouterDataModel
	//w.postgres.Table("router_data_models").Find(&routerCheck)
	//logrus.Info("---------------------")
	//for _, value := range routerCheck {
	//	logrus.Info(value)
	//}
	//logrus.Info("---------------------")

	//var routerCheck []WifiDataModel
	//w.postgres.Table("wifi_data_models").Find(&routerCheck)
	//logrus.Info("---------------------")
	//for _, value := range routerCheck {
	//	logrus.Info(value)
	//}
	//logrus.Info("---------------------")
	return nil
}

func findId(ids []FilePath) int64 {
	var max int64
	for _, value := range ids {
		if value.Id > max {
			max = value.Id
		}
	}
	return max
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

func createWifiModels(routersID int64, userID, filepathId int64) WifiDataModel {
	return WifiDataModel{
		IdUserData:   int(userID),
		IdRouterWifi: int(routersID),
		IdFilePath:   int(filepathId),
	}
}

func (w wifiData) GetData(userId int64) ([]model.Wifi, error) {
	var newUser model.User
	result := w.postgres.Where("id=?", userId).Find(&newUser)
	if result.Error != nil && newUser.Id == 0 {
		return nil, fmt.Errorf("user doesn't find: %w", result.Error)
	}
	dataResult, err := w.getDataFromDb(userId)
	if err != nil {
		return nil, err
	}
	return dataResult, nil
}

func (w wifiData) getDataFromDb(userId int64) ([]model.Wifi, error) {
	wifiDataModel := make([]WifiDataModel, 0, 10)
	result := w.postgres.Table("wifi_data_models").Where("id_user_data=?", userId).Find(&wifiDataModel)
	if result.Error != nil {
		return nil, result.Error
	}

	wifi := make([]model.Wifi, len(wifiDataModel), len(wifiDataModel)+1)

	routers := make([]RouterDataModel, 0, 10)
	var router RouterDataModel
	for _, value := range wifiDataModel {
		result = w.postgres.Table("router_data_models").Where("id_router=?", value.IdRouterWifi).Find(&router)
		if result.Error != nil {
			return nil, result.Error
		}
		routers = append(routers, router)
	}

	var file FilePath
	result = w.postgres.Table("file_path").Find(&file)
	if result.Error != nil {
		return nil, result.Error
	}

	routerSettings := make([]model.RouterSettings, 0, 10)
	for _, value := range routers {
		var coordPoint CoordinatesPoints
		result = w.postgres.Table("coordinates_points").Where("id=?", value.CoordinatesOfRouterID).Find(&coordPoint)
		if result.Error != nil {
			return nil, result.Error
		}
		routerSetting := convertRouterDataModelToRouterSettings(value, coordPoint)
		routerSettings = append(routerSettings, routerSetting)
	}
	for i, value := range wifiDataModel {
		wifi[i].User = int64(value.IdUserData)
		wifi[i].PathInput = file.PathInput
		wifi[i].PathOutput = file.PathOutput
		wifi[i].Router = append(wifi[i].Router, routerSettings[i])
	}
	return wifi, nil
}

func (w wifiData) findIdFilePath(id int) FilePath {
	var filepathCheck FilePath
	w.postgres.Table("file_path").Where("id=?", id).Find(&filepathCheck)
	return filepathCheck
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
	var wifiCheck WifiDataModel

	result := w.postgres.Where("id_user_data=? AND id_router_wifi=?", userId, routerId).Delete(&wifiCheck)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
