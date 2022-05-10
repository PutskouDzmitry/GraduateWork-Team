package api

import (
	"encoding/json"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h Handler) saveData(c *gin.Context) {
	_, err := h.GetUserFromToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}
	userId := 1
	filePath := "kek"
	routerss := []model.RouterSettings{
		{
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
		},
		{
			CoordinatesOfRouter: model.CoordinatesPoints{
				X: 200,
				Y: 600,
			},
			//мощность передатчика P
			TransmitterPower: 180,
			//коэффициент усиления передающей антенны Gt
			GainOfTransmittingAntenna: 50,
			//коэффициент усиления приемной антенны GT
			GainOfReceivingAntenna: 40,
			//чувствительность приемника на данной скорости Pmin
			Speed: 540,
			//потери сигнала в коаксиальном кабеле и разъемах передающего тракта Lt
			SignalLossTransmitting: -1,
			//потери сигнала в коаксиальном кабеле и разъемах приемного тракта LT
			SignalLossReceiving: -1,
			NumberOfChannels:    13,
			Scale:               1,
		},
	}
	err = h.wifiService.SaveData(routerss, int64(userId), filePath)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "data is saved")
}

func (h Handler) loadData(c *gin.Context) {
	_, err := h.GetUserFromToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}
	var userId = 1
	data, err := h.wifiService.GetData(int64(userId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	logrus.Info(data)
	jsonData, err := json.Marshal(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, jsonData)
}

func (h Handler) preloadData(c *gin.Context) {
	userId, err := h.GetUserFromToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}
	preloadData, err := h.wifiService.PreloadData(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	logrus.Info(preloadData)
	jsonData, err := json.Marshal(preloadData)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, jsonData)
}

func (h Handler) GetUserFromToken(token string) (int, error) {
	idStr, err := h.authService.ParseAccessToken(token)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(idStr)
}
