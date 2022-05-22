package api

import (
	"encoding/json"
	"fmt"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strconv"
)

func (h Handler) saveData(c *gin.Context) {
	_, err := h.GetUserFromToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		logrus.Info(err)
		//newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}
	userId := 2

	filePathInput, err := getImageFromContextForSave(c, "2", "myFile") //fmt.Sprint("./users_images/input/", userId , "-map.png")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	filePathOutput, err := getImageFromContextForSave(c, "2", "myFileOutput") //fmt.Sprint("./users_images/input/", userId , "-map.png")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	routers, err := getValuesOfRouters(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	//[]model.RouterSettings{
	//	{
	//		CoordinatesOfRouter: model.CoordinatesPoints{
	//			X: 200,
	//			Y: 300,
	//		},
	//		//мощность передатчика P
	//		TransmitterPower: 18,
	//		//коэффициент усиления передающей антенны Gt
	//		GainOfTransmittingAntenna: 5,
	//		//коэффициент усиления приемной антенны GT
	//		GainOfReceivingAntenna: 4,
	//		//чувствительность приемника на данной скорости Pmin
	//		Speed: 54,
	//		//потери сигнала в коаксиальном кабеле и разъемах передающего тракта Lt
	//		SignalLossTransmitting: -1,
	//		//потери сигнала в коаксиальном кабеле и разъемах приемного тракта LT
	//		SignalLossReceiving: -1,
	//		NumberOfChannels:    13,
	//		Scale:               1,
	//	},
	//	{
	//		CoordinatesOfRouter: model.CoordinatesPoints{
	//			X: 200,
	//			Y: 600,
	//		},
	//		//мощность передатчика P
	//		TransmitterPower: 180,
	//		//коэффициент усиления передающей антенны Gt
	//		GainOfTransmittingAntenna: 50,
	//		//коэффициент усиления приемной антенны GT
	//		GainOfReceivingAntenna: 40,
	//		//чувствительность приемника на данной скорости Pmin
	//		Speed: 540,
	//		//потери сигнала в коаксиальном кабеле и разъемах передающего тракта Lt
	//		SignalLossTransmitting: -1,
	//		//потери сигнала в коаксиальном кабеле и разъемах приемного тракта LT
	//		SignalLossReceiving: -1,
	//		NumberOfChannels:    13,
	//		Scale:               1,
	//	},
	//}
	err = h.wifiService.SaveData(routers, int64(userId), filePathInput, filePathOutput)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "data is saved")
}

func getImageFromContextForSave(c *gin.Context, userId string, typeOfFile string) (string, error) {
	c.Request.ParseMultipartForm(10 * 1024 * 1024)
	file, _, err := c.Request.FormFile(typeOfFile)
	if err != nil {
		return "", fmt.Errorf("error with get file from form: %w", err)
	}
	filename := service.GenerateFullPathOfFileForSave(inputPathFile, userId)
	out, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return "", fmt.Errorf("error with copy file: %w", err)
	}
	return filename, nil
}

func (h Handler) loadData(c *gin.Context) {
	_, err := h.GetUserFromToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		logrus.Info(err)
		//newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}
	var userId = 2
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

func (h Handler) deleteData(c *gin.Context) {
	_, err := h.GetUserFromToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}
	userId := 2
	routerId := 1
	err = h.wifiService.DeleteData(int64(userId), int64(routerId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "data is deleted")
}

func (h Handler) GetUserFromToken(token string) (int, error) {
	idStr, err := h.authService.ParseAccessToken(token)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(idStr)
}
