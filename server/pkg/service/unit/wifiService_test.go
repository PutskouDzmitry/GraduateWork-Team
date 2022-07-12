package unit

import (
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/model"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/service"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestCalculationOfValues(t *testing.T) {
	wifi := model.RouterSettings{
		CoordinatesOfRouter: model.CoordinatesPoints{
			X: 200,
			Y: 300,
		},
		//мощность передатчика P
		TransmitterPower: 18,
		//коэффициент усиления передающей антенны Gt
		GainOfTransmittingAntenna: 5,
		//коэффициент усиления приемной антенны GT
		GainOfReceivingAntenna: 4,
		//чувствительность приемника на данной скорости Pmin
		Speed: 54,
		//потери сигнала в коаксиальном кабеле и разъемах передающего тракта Lt
		SignalLossTransmitting: -1,
		//потери сигнала в коаксиальном кабеле и разъемах приемного тракта LT
		SignalLossReceiving: -1,
		NumberOfChannels:    13,
		Scale:               1,
	}
	// ответ в километрах
	value, err := service.CalculationOfValues(wifi)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(value)
}
