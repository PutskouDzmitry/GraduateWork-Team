package service

import (
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/model"
)

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

//func CalculateStatisticsInPoint() ([]model.ResponseOfGettingStatisticsOnPoint, error) {
//	responses := make([]model.ResponseOfGettingStatisticsOnPoint, 0, 10)
//	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
//		Name:           "Kek1",
//		MAC:            "kek1",
//		SignalStrength: -79,
//		SignalQuality:  20,
//		Frequency:      2.431,
//		MaxSpeed:       54,
//	})
//	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
//		Name:           "Kek2",
//		MAC:            "kek2",
//		SignalStrength: -79,
//		SignalQuality:  20,
//		Frequency:      2.431,
//		MaxSpeed:       54,
//	})
//	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
//		Name:           "Kek3",
//		MAC:            "kek3",
//		SignalStrength: -79,
//		SignalQuality:  20,
//		Frequency:      2.431,
//		MaxSpeed:       54,
//	})
//	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
//		Name:           "Kek4",
//		MAC:            "kek4",
//		SignalStrength: -79,
//		SignalQuality:  20,
//		Frequency:      2.431,
//		MaxSpeed:       54,
//	})
//	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
//		Name:           "Kek5",
//		MAC:            "kek5",
//		SignalStrength: -79,
//		SignalQuality:  20,
//		Frequency:      2.431,
//		MaxSpeed:       54,
//	})
//	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
//		Name:           "Kek6",
//		MAC:            "kek6",
//		SignalStrength: -79,
//		SignalQuality:  20,
//		Frequency:      2.431,
//		MaxSpeed:       54,
//	})
//	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
//		Name:           "Kek7",
//		MAC:            "kek7",
//		SignalStrength: -79,
//		SignalQuality:  20,
//		Frequency:      2.431,
//		MaxSpeed:       54,
//	})
//	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
//		Name:           "Kek8",
//		MAC:            "kek8",
//		SignalStrength: -79,
//		SignalQuality:  20,
//		Frequency:      2.431,
//		MaxSpeed:       54,
//	})
//	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
//		Name:           "Kek9",
//		MAC:            "kek9",
//		SignalStrength: -79,
//		SignalQuality:  20,
//		Frequency:      2.431,
//		MaxSpeed:       54,
//	})
//	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
//		Name:           "Kek10",
//		MAC:            "kek10",
//		SignalStrength: -79,
//		SignalQuality:  20,
//		Frequency:      2.431,
//		MaxSpeed:       54,
//	})
//	responses = append(responses, model.ResponseOfGettingStatisticsOnPoint{
//		Name:           "Kek11",
//		MAC:            "kek11",
//		SignalStrength: -79,
//		SignalQuality:  20,
//		Frequency:      2.431,
//		MaxSpeed:       54,
//	})
//	return responses, nil
//}

type routerSettingForMigrator struct {
	coord  model.CoordinatesPoints
	Name   string
	Power  float64
	MAC    string
	radius float64
}

func CalculateStatisticsInPoint(point model.CoordinatesPoints, routers []model.RoutersSettingForMigrator) ([]model.ResponseOfGettingStatisticsOnPoint, error) {
	powers := make([]float64, 0, 10)
	powersMin := make([]float64, 0, 10)
	minPowers := make([]valueOfPowerOnPoint, 0, 10)
	maxPowers := make([]valueOfPowerOnPoint, 0, 10)
	for _, value := range routers {
		for _, valueOfPoint := range value.RoutersSettingsMigration {
			powers = append(powers, valueOfPoint.Power)
		}
		maxPowers = append(maxPowers, findMaxPower(powers, value))

		maxPowerOnPoint := findMaxPower(powers, value)
		for _, valueOfPoint := range value.RoutersSettingsMigration {
			if valueOfPoint.MAC == maxPowerOnPoint.router.MAC {
				powersMin = append(powersMin, valueOfPoint.Power)
			}
		}
		minPowers = append(minPowers, findMinPower(powersMin, value))
		powers = make([]float64, 0, 10)
		powersMin = make([]float64, 0, 10)
	}

	distance := make([]routerSettingForMigrator, 0, 10)
	for i, value := range maxPowers {
		distance = append(distance, routerSettingForMigrator{
			coord:  value.coordinates,
			radius: getRadius(value.coordinates.X, value.coordinates.Y, minPowers[i].coordinates.X, minPowers[i].coordinates.Y, true),
			MAC:    minPowers[i].router.MAC,
			Name:   minPowers[i].router.Name,
			Power:  minPowers[i].router.Power,
		})
	}
	statRouters := make([]model.ResponseOfGettingStatisticsOnPoint, 0, 10)
	for _, value := range distance {
		radiusOfPoint := getRadius(point.X, point.Y, value.coord.X, value.coord.Y, false)
		diff := radiusOfPoint / value.radius
		statRouters = append(statRouters, model.ResponseOfGettingStatisticsOnPoint{
			Name:           value.Name,
			MAC:            value.MAC,
			SignalStrength: value.Power * diff,
			SignalQuality:  float64(int(1 / diff * 100)),
			Frequency:      -1,
			MaxSpeed:       -1,
		})
	}
	return statRouters, nil
}

func checkInShapes(point model.CoordinatesPoints, routers []routerSettingForMigrator) []routerSettingForMigrator {
	validRouters := make([]routerSettingForMigrator, 0, 10)
	var check bool = false
	for _, value := range routers {
		if value.coord.X+value.radius < point.X {
			check = true
		}

		if value.coord.X-value.radius > point.X {
			check = true
		}

		if value.coord.Y+value.radius < point.X {
			check = true
		}

		if value.coord.Y-value.radius > point.X {
			check = true
		}

		if check == false {
			validRouters = append(validRouters, value)
		}
	}
	return validRouters
}
