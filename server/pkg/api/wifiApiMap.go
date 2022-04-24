package api

import (
	"fmt"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/model"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	b64 "encoding/base64"
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
	var routers []model.RouterSettings
	//var u us
	//if err := c.BindJSON(&u); err != nil {
	//newErrorResponse(c, http.StatusBadRequest, err.Error())
	//return
	//}
	header := c.GetHeader(authorizationHeader)
	userId, err := h.getUserId(header)
	if err != nil {
		//newErrorResponse(c, http.StatusInternalServerError, err.Error())
		//return
	}
	routers = testValue()
	filePathInput, err := getImageFromContext(c, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// validation
	err = service.ValidationOfPlaceRouter(filePathInput, routers)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	validationStartValue(&routers)
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
s := string(fileBytes)
sEnc := b64.StdEncoding.EncodeToString([]byte(s))
c.Writer.WriteString(sEnc)
}

func (h Handler) getUserId(header string) (string, error) {
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

func getFile(c *gin.Context) {
	path := "C://Users/Dzmitry_Putskou/go/src/github.com/PutskouDzmitry/GraduateWork-Team/users_images/"
	mr, err := c.Request.MultipartReader()
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fileName := "dima"

	filename := service.GenerateFullPathOfFile(inputPathFile, "1")
	out, err := os.Create(filename)
	if err != nil {
		return
	}
	defer out.Close()
	for {
		part, err := mr.NextPart()

		//Break, if no more file
		if err == io.EOF {
			break
		}
		//Get File Name attribute
		fileName = part.FileName()

		//Open the file, Create file if it's not existing
		dst, err := os.OpenFile(path+fileName, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return
		}

		for {
			//Create and read the file into temparory byte array
			buffer := make([]byte, 1)
			cBytes, err := part.Read(buffer)

			//break when the reading process is finished
			if err == io.EOF {
				break
			}

			//Write into the file from byte array
			dst.Write(buffer[0:cBytes])
		}
		defer dst.Close()
	}

}

func getImageFromContext(c *gin.Context, userId string) (string, error) {
	c.Request.ParseMultipartForm(10 * 1024 * 1024)
	file, _, err := c.Request.FormFile("myFile")
	//t := c.Request.FormValue("testInput")
	//logrus.Info(t)
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

func getDataFromForm() {

}

func (h Handler) saveData(c *gin.Context) {
	var routers model.RouterSettings
	if err := c.BindJSON(&routers); err != nil {
		//newErrorResponse(c, http.StatusBadRequest, err.Error())
		//return
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
	err := h.wifiService.SaveData(routerss, int64(userId), filePath)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "data is saved")
}

func (h Handler) loadData(c *gin.Context) {
	var routers model.User
	if err := c.BindJSON(&routers); err != nil {
		//newErrorResponse(c, http.StatusBadRequest, err.Error())
		//return
	}
	var userId = 1
	data, err := h.wifiService.GetData(int64(userId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	logrus.Info(data)
	c.JSON(http.StatusOK, "data is loaded")
}
