package service

import "github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/model"

type statistics struct {
	routersSettings []model.RouterSettings
	point           model.CoordinatesPoints
}

func NewCalculationStatistics(routerSettings []model.RouterSettings, point model.CoordinatesPoints) statistics {
	return statistics{
		routersSettings: routerSettings,
		point:           point,
	}
}

func (s statistics) CalculateStatisticsInPoint() ([]model.ResponseOfGettingStatisticsOnPoint, error) {

	return nil, nil
}
