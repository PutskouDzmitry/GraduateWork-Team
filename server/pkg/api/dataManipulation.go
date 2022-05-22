package api

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/model"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
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

	//filePathInput, err := getImageFromContextForSave(c, "2", "myFile") //fmt.Sprint("./users_images/input/", userId , "-map.png")
	//if err != nil {
	//	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	//}
	//filePathOutput, err := getImageFromContextForSave(c, "2", "myFileOutput") //fmt.Sprint("./users_images/input/", userId , "-map.png")
	//if err != nil {
	//	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	//}
	//routers, err := getValuesOfRouters(c)
	//if err != nil {
	//	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	//}
	filePathOutput := fmt.Sprint("./users_images/input/", userId, "-map.png")
	filePathInput := fmt.Sprint("./users_images/input/", userId, "-map.png")
	routers := []model.RouterSettings{
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
	c.JSON(http.StatusOK, convertToResponseData(c, data))
}

func convertToResponseData(c *gin.Context, wifi []model.Wifi) model.Response {
	var data model.Response
	var dataWifi []model.WifiResponseForManipulation
	for _, value := range wifi {
		fileBytesInput, err := ioutil.ReadFile(value.PathInput)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return model.Response{}
		}
		sEncInput := b64.StdEncoding.EncodeToString(fileBytesInput)

		fileBytesOutput, err := ioutil.ReadFile(value.PathOutput)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return model.Response{}
		}
		sEncOutPut := b64.StdEncoding.EncodeToString(fileBytesOutput)
		dataWifi = append(dataWifi, model.WifiResponseForManipulation{
			Router:     value.Router,
			PathInput:  sEncInput,
			PathOutput: sEncOutPut,
		})
	}
	data.User = wifi[0].User
	data.Data = dataWifi
	return data
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
