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
			TransmitterPower:          200,
			GainOfTransmittingAntenna: 1,
			GainOfReceivingAntenna:    2,
			Speed:                     48,
			SignalLossTransmitting:    0,
			SignalLossReceiving:       0,
			NumberOfChannels:          2,
			Scale:                     7,
			Thickness:                 10,
		},
		{
			CoordinatesOfRouter: model.CoordinatesPoints{
				X: 100,
				Y: 200,
			},
			TransmitterPower:          0,
			GainOfTransmittingAntenna: 0,
			GainOfReceivingAntenna:    0,
			Speed:                     48,
			SignalLossTransmitting:    0,
			SignalLossReceiving:       0,
			NumberOfChannels:          3,
			Scale:                     7,
			Thickness:                 5,
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
	fileBytes, err := ioutil.ReadFile(pathOfOutImage + filePathOutput)
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
	filename := pathOfOutImage + header.Filename
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
