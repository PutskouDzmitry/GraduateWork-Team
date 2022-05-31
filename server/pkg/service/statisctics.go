package service

import "github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/model"

type statistics struct {
	routersSettings []model.RouterSettings
	point           model.CoordinatesPoints
}

//func NewCalculationStatistics(routerSettings []model.RouterSettings, point model.CoordinatesPoints) statistics {
//	return statistics{
//		routersSettings: routerSettings,
//		point:           point,
//	}
//}

func CalculateStatisticsInPoint() ([]model.ResponseOfGettingStatisticsOnPoint, error) {
	responses := make([]model.ResponseOfGettingStatisticsOnPoint, 0, 10)
	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
		Name:          "Kek1",
		MAC:           "kek1",
		SignalStrange: -79,
		SignalQuality: 20,
		Frequency:     2.431,
		MaxSpeed:      54,
	})
	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
		Name:          "Kek2",
		MAC:           "kek2",
		SignalStrange: -79,
		SignalQuality: 20,
		Frequency:     2.431,
		MaxSpeed:      54,
	})
	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
		Name:          "Kek3",
		MAC:           "kek3",
		SignalStrange: -79,
		SignalQuality: 20,
		Frequency:     2.431,
		MaxSpeed:      54,
	})
	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
		Name:          "Kek4",
		MAC:           "kek4",
		SignalStrange: -79,
		SignalQuality: 20,
		Frequency:     2.431,
		MaxSpeed:      54,
	})
	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
		Name:          "Kek5",
		MAC:           "kek5",
		SignalStrange: -79,
		SignalQuality: 20,
		Frequency:     2.431,
		MaxSpeed:      54,
	})
	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
		Name:          "Kek6",
		MAC:           "kek6",
		SignalStrange: -79,
		SignalQuality: 20,
		Frequency:     2.431,
		MaxSpeed:      54,
	})
	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
		Name:          "Kek7",
		MAC:           "kek7",
		SignalStrange: -79,
		SignalQuality: 20,
		Frequency:     2.431,
		MaxSpeed:      54,
	})
	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
		Name:          "Kek8",
		MAC:           "kek8",
		SignalStrange: -79,
		SignalQuality: 20,
		Frequency:     2.431,
		MaxSpeed:      54,
	})
	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
		Name:          "Kek9",
		MAC:           "kek9",
		SignalStrange: -79,
		SignalQuality: 20,
		Frequency:     2.431,
		MaxSpeed:      54,
	})
	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
		Name:          "Kek10",
		MAC:           "kek10",
		SignalStrange: -79,
		SignalQuality: 20,
		Frequency:     2.431,
		MaxSpeed:      54,
	})
	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
		Name:          "Kek11",
		MAC:           "kek11",
		SignalStrange: -79,
		SignalQuality: 20,
		Frequency:     2.431,
		MaxSpeed:      54,
	})
	return responses, nil
}

//func CalculateStatisticsInPoint(point model.CoordinatesPoints, routers model.RequestFlux) ([]model.ResponseOfGettingStatisticsOnPoint, error) {
//
//}
//
//func (s statistics) calculateRadius(routers []model.RequestFlux) {
//	powers := make([]float64, 0, 10)
//	powersMin := make([]float64, 0, 10)
//	minPowers := make([]valueOfPowerOnPoint, 0, 10)
//	maxPowers := make([]valueOfPowerOnPoint, 0, 10)
//	for _, value := range s.routersSettings {
//		for _, valueOfPoint := range value.RoutersSettingsMigration {
//			powers = append(powers, valueOfPoint.Power)
//		}
//		maxPowers = append(maxPowers, findMaxPower(powers, value))
//
//		maxPowerOnPoint := findMaxPower(powers, value)
//		for _, value := range d.coordinatesOfRouters {
//			for _, valueOfPoint := range value.RoutersSettingsMigration {
//				if valueOfPoint.MAC == maxPowerOnPoint.router.MAC {
//					powersMin = append(powersMin, valueOfPoint.Power)
//				}
//			}
//		}
//		minPowers = append(minPowers, findMinPower(powersMin, value))
//		powers = make([]float64, 0, 10)
//		powersMin = make([]float64, 0, 10)
//	}
//}
