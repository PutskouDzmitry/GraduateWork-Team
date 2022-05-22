package service

import (
	"fmt"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/data"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/model"
	"math"
)

var (
	COM float64 = 10
)

type wifi struct {
	transmitterPower          float64
	gainOfTransmittingAntenna float64
	gainOfReceivingAntenna    float64
	receiverSensitivity       float64
	signalLossTransmitting    float64
	signalLossReceiving       float64
	typeOfSignal              float64
}

func CalculationOfValues(coordinates model.RouterSettings) (float64, error) {
	sensitivity := getSensitivityVersusBaudRate(coordinates.Speed)
	if sensitivity == -1 {
		return -1, fmt.Errorf("this program doesn't supposed this speed: %v", coordinates.Speed)
	}

	numberOfChannel := getCenterFrequency(coordinates.NumberOfChannels, coordinates.TypeOfSignal)
	if numberOfChannel == -1 {
		return -1, fmt.Errorf("this program doesn't supposed this number of channel: %v", coordinates.NumberOfChannels)
	}
	numberOfChannel *= 1000

	wifiSignal := initValues(coordinates.TransmitterPower, coordinates.GainOfTransmittingAntenna, coordinates.GainOfReceivingAntenna, sensitivity,
		coordinates.SignalLossTransmitting, coordinates.SignalLossReceiving)
	FSL := wifiSignal.getTotalSystemGain(COM)
	distance := wifiSignal.getCommunicationRange(FSL, numberOfChannel)
	return distance * 1000, nil
}

func (w wifi) getTotalSystemGain(COM float64) float64 {
	//logrus.Info(w.transmitterPower, w.gainOfTransmittingAntenna, w.gainOfReceivingAntenna, w.receiverSensitivity, w.signalLossTransmitting, w.signalLossReceiving)
	return w.transmitterPower + w.gainOfTransmittingAntenna + w.gainOfReceivingAntenna - w.receiverSensitivity - w.signalLossTransmitting - w.signalLossReceiving - COM
}

func (w wifi) getCommunicationRange(FSL, F float64) float64 {
	FAfterLog := math.Log10(F)
	number := (FSL-33)/20 - FAfterLog
	return math.Pow(10, number)
}

func getSensitivityVersusBaudRate(speed int) float64 {
	switch speed {
	case 54:
		return -66
	case 48:
		return -71
	case 36:
		return -76
	case 24:
		return -80
	case 18:
		return -83
	case 12:
		return -85
	case 9:
		return -86
	case 6:
		return -87
	default:
		return float64(speed)
	}
}

func getCenterFrequency(number int, typeOfSignal float64) float64 {
	if typeOfSignal == 2.4 {
		switch number {
		case 1:
			return 2.412
		case 2:
			return 2.417
		case 3:
			return 2.422
		case 4:
			return 2.427
		case 5:
			return 2.432
		case 6:
			return 2.437
		case 7:
			return 2.447
		case 8:
			return 2.452
		case 9:
			return 2.422
		case 10:
			return 2.457
		case 11:
			return 2.462
		case 12:
			return 2.467
		case 13:
			return 2.472
		case 14:
			return 2.484
		default:
			return -1
		}
	} else {
		switch number {
		case 1:
			return 5.180
		case 2:
			return 5.200
		case 3:
			return 5.220
		case 4:
			return 5.240
		case 5:
			return 5.260
		case 6:
			return 5.280
		case 7:
			return 5.300
		case 8:
			return 5.320
		case 9:
			return 5.340
		case 10:
			return 5.360
		case 11:
			return 5.380
		case 12:
			return 5.400
		case 13:
			return 5.420
		case 14:
			return 5.440
		case 15:
			return 5.460
		case 16:
			return 5.480
		case 17:
			return 5.500
		case 18:
			return 5.520
		case 19:
			return 5.540
		case 20:
			return 5.560
		case 21:
			return 5.580
		case 22:
			return 5.600
		case 23:
			return 5.620
		case 24:
			return 5.640
		case 25:
			return 5.660
		default:
			return -1
		}
	}
}

func initValues(transmitterPower, gainOfTransmittingAntenna, gainOfReceivingAntenna, receiverSensitivity, signalLossTransmitting, signalLossReceiving float64) *wifi {
	return &wifi{
		transmitterPower:          transmitterPower,
		gainOfTransmittingAntenna: gainOfTransmittingAntenna,
		gainOfReceivingAntenna:    gainOfReceivingAntenna,
		receiverSensitivity:       receiverSensitivity,
		signalLossTransmitting:    signalLossTransmitting,
		signalLossReceiving:       signalLossReceiving,
	}
}

type wifiService struct {
	wifi data.WifiData
}

type WifiService interface {
	SaveData(routers []model.RouterSettings, userId int64, pathInput, pathOutput string) error
	GetData(userId int64) ([]model.Wifi, error)
	DeleteData(userId, routerId int64) error
}

func NewWifiService(wifi data.WifiData) WifiService {
	return &wifiService{wifi: wifi}
}

func (w wifiService) SaveData(routers []model.RouterSettings, userId int64, pathInput, pathOutput string) error {
	return w.wifi.SaveData(routers, userId, pathInput, pathOutput)
}

func (w wifiService) GetData(userId int64) ([]model.Wifi, error) {
	return w.wifi.GetData(userId)
}

func (w wifiService) DeleteData(userId, routerId int64) error {
	return w.wifi.DeleteData(userId, routerId)
}
