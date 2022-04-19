package unit

import (
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/data"
	"gorm.io/gorm"
	"testing"
)

func createGormConnection() (*gorm.DB, error) {
	connPostgres, err := data.GetConnectionPostgres("localhost", "5432", "postgres", "postgres", "password", "disable")
	if err != nil {
		return nil, err
	}
	return connPostgres, nil
}

func TestSaveData(t *testing.T) {
	gormConnection, err := createGormConnection()
	if err != nil {
		t.Fatal(err)
	}

	//testModel := data.User{
	//	//Model:    gorm.Model{},/
	//	UserName: "test",
	//	RouterID: 0,
	//	Router:   data.Router{
	//		IdRouter:                  1,
	//		TransmitterPower:          1,
	//		GainOfTransmittingAntenna: 1,
	//		GainOfReceivingAntenna:    1,
	//		Speed:                     1,
	//		SignalLossTransmitting:    1,
	//		SignalLossReceiving:       1,
	//		NumberOfChannels:          1,
	//		Scale:                     1,
	//		Thickness:                 1,
	//		COM:                       1,
	//	},
	//}
	//
	testModel := data.WifiDataModel{
		IdUserData:   1,
		UserData:     data.WifiUserModels{
			UserModelId: 2,
			Username:    "test",
			Password:    "test",
		},
		IdRouterWifi: 1,
		RouterWifi: data.RouterDataModel{
			IdRouter:                  1,
			CoordinatesOfRouterID:     1,
			CoordinatesOfRouter:       data.CoordinatesPoints{
				IdCoordinates: 1,
				X:             1,
				Y:             1,
			},
			TransmitterPower:          1,
			GainOfTransmittingAntenna: 1,
			GainOfReceivingAntenna:    1,
			Speed:                     1,
			SignalLossTransmitting:    1,
			SignalLossReceiving:       1,
			NumberOfChannels:          1,
			Scale:                     1,
			Thickness:                 1,
			COM:                       1,
		},
		Path:     "test",
	}
	
	wifiData := data.NewWifiData(gormConnection)
	err = wifiData.SaveData(testModel)
	if err != nil {
		t.Fatal(err)
	}
}
