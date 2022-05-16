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
	fileBytes, err := ioutil.ReadFile(service.GenerateFullPathOfFileToMap(outputPathFile, userId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	sEnc := b64.StdEncoding.EncodeToString(fileBytes)
	c.Writer.WriteString(sEnc)
}

func getValuesToFlux(c *gin.Context) ([]model.RoutersSettingForMigrator, error) {
	data := c.Request.FormValue("data")
	var settings []model.RequestRouters
	dataInByte := []byte(data)
	err := json.Unmarshal(dataInByte, &settings)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func migrateDataToFlux() {

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

func migrateDataToAcrylic() {

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

func migrateDataToPhone() {

}
