package api

import (
	"encoding/json"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/model"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h Handler) getStatisticsInPoint(c *gin.Context) {
	routers, err := getValuesOfRouters(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	pointForStatistics, err := getPointForStatistics(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	statistics := service.NewCalculationStatistics(routers, pointForStatistics)
	getStatistics, err := statistics.CalculateStatisticsInPoint()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	logrus.Info(getStatistics)
}

func getPointForStatistics(c *gin.Context) (model.CoordinatesPoints, error) {
	data := c.Request.FormValue("point")
	var settings model.CoordinatesPoints
	dataInByte := []byte(data)
	err := json.Unmarshal(dataInByte, &settings)
	if err != nil {
		return model.CoordinatesPoints{}, err
	}
	return settings, nil
}
