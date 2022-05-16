package api

import (
	b64 "encoding/base64"
	"encoding/json"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/model"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (h Handler) fluxMigrator(c *gin.Context) {
	userId := "2"
	dataOfRouters, err := getValuesToFlux(c)
	if err != nil {
		logrus.Error(err)
	}
	filePathInput, err := getImageFromContext(c, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	filePathOutput := service.GenerateFullPathOfFileToFlux(outputPathFile, userId)
	drawImage := service.NewDrawImageToMigrator(filePathInput, filePathOutput, dataOfRouters)
	err = drawImage.FluxMigrator()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	fileBytes, err := ioutil.ReadFile(service.GenerateFullPathOfFileToFlux(outputPathFile, userId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	sEnc := b64.StdEncoding.EncodeToString(fileBytes)
	c.Writer.WriteString(sEnc)
}

func getValuesToFlux(c *gin.Context) ([]model.RoutersSettingForMigrator, error) {
	data := c.Request.FormValue("data")
	var settings model.RequestFlux
	dataInByte := []byte(data)
	err := json.Unmarshal(dataInByte, &settings)
	routers := make([]model.RoutersSettingForMigrator, 0, 10)
	router := make([]model.RouterSettingForMigrator, 0, 10)
	if err != nil {
		return nil, err
	}
	for _, value := range settings.Steps {
		for _, value := range settings.AcsParsed {
			for _, valueOfPoint := range value.Signals {
				if s, err := strconv.ParseFloat(valueOfPoint.Obj.LastSignalStrength, 64); err == nil {
					router = append(router, model.RouterSettingForMigrator{
						Name:  valueOfPoint.Obj.AdId,
						Power: s,
						MAC:   valueOfPoint.Obj.MAC,
					})
				}
			}
		}
		routers = append(routers, model.RoutersSettingForMigrator{
			Coordinates: model.CoordinatesPoints{
				X: value.Coords.X,
				Y: value.Coords.Y,
			},
			RoutersSettingsMigration: router,
		})
	}
	logrus.Info(len(routers[0].RoutersSettingsMigration))
	return routers, nil
}

func (h Handler) acrylicMigrator(c *gin.Context) {
	userId := "2"
	dataOfRouters, err := getDataToAcrylic(c)
	if err != nil {
		logrus.Error(err)
	}
	filePathInput, err := getImageFromContext(c, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	filePathOutput := service.GenerateFullPathOfFileToAcrylic(outputPathFile, userId)
	drawImage := service.NewDrawImageToMigrator(filePathInput, filePathOutput, dataOfRouters)
	err = drawImage.AcrylicMigrator()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	fileBytes, err := ioutil.ReadFile(service.GenerateFullPathOfFileToMap(outputPathFile, userId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	sEnc := b64.StdEncoding.EncodeToString(fileBytes)
	c.Writer.WriteString(sEnc)
}

func getDataToAcrylic(c *gin.Context) ([]model.RoutersSettingForMigrator, error) {
	return nil, nil
}

func (h Handler) telephoneMigrator(c *gin.Context) {
	userId := "2"
	dataOfRouters, err := getDataToTelephone(c)
	if err != nil {
		logrus.Error(err)
	}
	filePathInput, err := getImageFromContext(c, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	filePathOutput := service.GenerateFullPathOfFileToTelephone(outputPathFile, userId)
	drawImage := service.NewDrawImageToMigrator(filePathInput, filePathOutput, dataOfRouters)
	err = drawImage.TelephoneMigrator()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	fileBytes, err := ioutil.ReadFile(service.GenerateFullPathOfFileToMap(outputPathFile, userId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	sEnc := b64.StdEncoding.EncodeToString(fileBytes)
	c.Writer.WriteString(sEnc)
}

func getDataToTelephone(c *gin.Context) ([]model.RoutersSettingForMigrator, error) {
	return nil, nil
}
