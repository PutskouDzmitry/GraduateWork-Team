package api

import (
	"fmt"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/model"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/service"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	// basic values for router
	numberOfChannels       = 1
	signalLossTransmitting = 0
	signalLossReceiving    = 0
	scale                  = 5
	thickness              = 10
	com                    = 10
)

func testValue() []model.RouterSettings {
	return []model.RouterSettings{
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
	}
}

var (
	testFileName   = ""
	pathOfOutImage = "./pictures/"
)

func (h Handler) calculationOfValues(c *gin.Context) {
	var routers []model.RouterSettings
	//if err := c.BindJSON(&routers); err != nil {
	//	//newErrorResponse(c, http.StatusBadRequest, err.Error())
	//	//return
	//}
	routers = testValue()
	filePathInput, err := getImageFromContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// validation
	validationStartValue(&routers)
	// username
	var userId string = "Dima"
	filePathOutput, err := generateFilePathOutput(userId, pathOfOutImage)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	drawImage := service.NewDrawImage(routers, filePathInput, filePathOutput)
	err = drawImage.DrawOnImage()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	fileBytes, err := ioutil.ReadFile("gradient-conic.png")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.Writer.Write(fileBytes)
}

func validationStartValue(routers *[]model.RouterSettings) {
	for _, value := range *routers {
		if value.NumberOfChannels == -1 {
			value.NumberOfChannels = numberOfChannels
		}
		if value.SignalLossTransmitting == -1 {
			value.SignalLossTransmitting = float64(signalLossTransmitting)
		}
		if value.SignalLossReceiving == -1 {
			value.SignalLossReceiving = float64(signalLossReceiving)
		}
		if value.Scale == -1 {
			value.Scale = float64(scale)
		}
		if value.Thickness == -1 {
			value.Thickness = float64(thickness)
		}
		if value.COM == -1 {
			value.COM = float64(com)
		}
	}
}

func generateFilePathOutput(userId, pathOfOutImage string) (string, error) {
	return userId, nil
}

func getImageFromContext(c *gin.Context) (string, error) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return "", fmt.Errorf("error with get file from form: %w", err)
	}
	filename := header.Filename
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

func (h Handler) saveData(c *gin.Context) {
	var routers []model.RouterSettings
	if err := c.BindJSON(&routers); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.wifiService.SaveData()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "data is saved")
}

func (h Handler) loadData(c *gin.Context) {

}
