package service

import (
	"fmt"
	"github.com/PutskouDzmitry/GraduateWork/server/pkg/data"
	"github.com/PutskouDzmitry/GraduateWork/server/pkg/model"
	"github.com/sirupsen/logrus"
	"math"
)

type wifi struct {
	transmitterPower          float64
	gainOfTransmittingAntenna float64
	gainOfReceivingAntenna    float64
	receiverSensitivity       float64
	signalLossTransmitting    float64
	signalLossReceiving       float64
}

func CalculationOfValues(coordinates model.RouterSettings) (float64, error) {
	var COM float64 = 10

	sensitivity := getSensitivityVersusBaudRate(coordinates.Speed)
	if sensitivity == -1 {
		return -1, fmt.Errorf("this program doesn't supposed this speed: %v", coordinates.Speed)
	}

	numberOfChannel := getCenterFrequency(coordinates.NumberOfChannels)
	if numberOfChannel == -1 {
		return -1, fmt.Errorf("this program doesn't supposed this number of channel: %v", coordinates.NumberOfChannels)
	}
	numberOfChannel *= 1000

	wifiSignal := initValues(coordinates.TransmitterPower, coordinates.GainOfTransmittingAntenna, coordinates.GainOfReceivingAntenna, sensitivity,
		coordinates.SignalLossTransmitting, coordinates.SignalLossReceiving)
	FSL := wifiSignal.getTotalSystemGain(COM)
	distance := wifiSignal.getCommunicationRange(FSL, numberOfChannel)
	logrus.Info("distance: ", distance)
	return distance, nil
}

func (w wifi) getTotalSystemGain(COM float64) float64 {
	return w.transmitterPower + w.gainOfTransmittingAntenna + w.gainOfReceivingAntenna - w.receiverSensitivity - w.signalLossTransmitting - w.signalLossReceiving - COM
}

func (w wifi) getCommunicationRange(FSL, F float64) float64 {
	FAfterLog := math.Log10(F)
	logrus.Info("F ", FAfterLog)
	number := (FSL-33)/20 - FAfterLog
	logrus.Info("Number", number)
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
		return -1
	}
}

func getCenterFrequency(number int) float64 {
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
}

func NewWifiService(wifi data.WifiData) WifiService {
	return &wifiService{wifi: wifi}
}
