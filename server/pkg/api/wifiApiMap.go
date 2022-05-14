package api

import (
	b64 "encoding/base64"
	"encoding/json"
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
	"strings"
)

//detect path of pictures
var (
	inputPathFile  = "./users_images/input/"
	outputPathFile = "./users_images/output/"
)

func testValue() []model.RouterSettings {
	return []model.RouterSettings{
		{
			CoordinatesOfRouter: model.CoordinatesPoints{
				X: 300,
				Y: 200,
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
				X: 350,
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

type us struct {
	Name string
	Age  int64
}

func (h Handler) calculationOfValues(c *gin.Context) {
	//_, err := h.GetUserFromToken(c.GetHeader(authorizationHeader))
	//if err != nil {
	//	newErrorResponse(c, http.StatusUnauthorized, err.Error())
	//}

	//var u us
	//if err := c.BindJSON(&u); err != nil {
	//newErrorResponse(c, http.StatusBadRequest, err.Error())
	//return
	//}
	//header := c.GetHeader(authorizationHeader)
	//userId, err := h.getUserId(header)
	//if err != nil {
	//	//newErrorResponse(c, http.StatusInternalServerError, err.Error())
	//	//return
	//}
	userId := "2"
	routersOld, err := getValues(c)
	if err != nil {
		logrus.Error(err)
	}
	filePathInput, err := getImageFromContext(c, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// validation
	err = service.ValidationOfPlaceRouter(filePathInput, routersOld)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	routers := routersOld //service.ValidationValues(routersOld)
	filePathOutput := service.GenerateFullPathOfFile(outputPathFile, userId)
	drawImage := service.NewDrawImage(routers, filePathInput, filePathOutput)
	err = drawImage.DrawOnImage()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	fileBytes, err := ioutil.ReadFile(service.GenerateFullPathOfFile(outputPathFile, userId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	sEnc := b64.StdEncoding.EncodeToString(fileBytes)
	c.Writer.WriteString(sEnc)
}

func (h Handler) getUserId(c *gin.Context, header string) (string, error) {
	if c.Request.URL.String() != "/api/map/calculation" {
		return "200", nil
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return "", fmt.Errorf("invalid auth header")
	}
	userId, err := h.authService.ParseAccessToken(headerParts[1])
	if err != nil {
		return "", err
	}
	return userId, nil
}

func getValues(c *gin.Context) ([]model.RouterSettings, error) {
	data := c.Request.FormValue("data")
	var settings []model.RequestRouters
	dataInByte := []byte(data)
	err := json.Unmarshal(dataInByte, &settings)
	if err != nil {
		return nil, err
	}
	routerSettings := make([]model.RouterSettings, len(settings), len(settings)+1)
	for i, value := range settings {
		routerSettings[i].CoordinatesOfRouter.X = value.Coords.X
		routerSettings[i].CoordinatesOfRouter.Y = value.Coords.Y
		transmitterPower, _ := strconv.ParseFloat(value.Settings.TransmitterPower, 8)
		routerSettings[i].TransmitterPower = transmitterPower
		gainOfTransmittingAntenna, _ := strconv.ParseFloat(value.Settings.GainOfTransmittingAntenna, 8)
		routerSettings[i].GainOfTransmittingAntenna = gainOfTransmittingAntenna
		gainOfReceivingAntenna, _ := strconv.ParseFloat(value.Settings.GainOfReceivingAntenna, 8)
		routerSettings[i].GainOfReceivingAntenna = gainOfReceivingAntenna
		speed, _ := strconv.Atoi(value.Settings.Speed)
		routerSettings[i].Speed = speed
		signalLossTransmitting, _ := strconv.ParseFloat(value.Settings.SignalLossTransmitting, 8)
		routerSettings[i].SignalLossTransmitting = signalLossTransmitting
		signalLossReceiving, _ := strconv.ParseFloat(value.Settings.SignalLossReceiving, 8)
		routerSettings[i].SignalLossReceiving = signalLossReceiving
		numberOfChannels, _ := strconv.Atoi(value.Settings.NumberOfChannels)
		routerSettings[i].NumberOfChannels = numberOfChannels
		routerSettings[i].Scale = 1
		routerSettings[i].Thickness = 4
		routerSettings[i].COM = 10
	}
	return routerSettings, nil
}

func getImageFromContext(c *gin.Context, userId string) (string, error) {
	c.Request.ParseMultipartForm(10 * 1024 * 1024)
	file, _, err := c.Request.FormFile("myFile")
	if err != nil {
		return "", fmt.Errorf("error with get file from form: %w", err)
	}
	filename := service.GenerateFullPathOfFile(inputPathFile, userId)
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

func (h Handler) getInfo(c *gin.Context) {
	jsonData, err := json.Marshal(mockDataGetInfo())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, jsonData)
}

func mockDataGetInfo() model.ResponseInfoPoint {
	infoOfPoint := make([]model.InfoOfPoint, 2, 2)
	infoOfPoint = append(infoOfPoint, model.InfoOfPoint{
		NameOfRouter:   "test 1",
		CurrentSpeed:   24,
		MaxSpeed:       54,
		SignalStrength: -50,
		SignalQuality:  40,
		Channel:        2,
	})
	infoOfPoint = append(infoOfPoint, model.InfoOfPoint{
		NameOfRouter:   "test 2",
		CurrentSpeed:   30,
		MaxSpeed:       64,
		SignalStrength: -20,
		SignalQuality:  80,
		Channel:        5,
	})
	return model.ResponseInfoPoint{InfoPoint: infoOfPoint}
}
