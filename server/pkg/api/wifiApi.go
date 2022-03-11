package api

import (
	"fmt"
	"github.com/PutskouDzmitry/GraduateWork/server/pkg/model"
	"github.com/PutskouDzmitry/GraduateWork/server/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
	"os"
)

func (h Handler) calculationOfValues(c *gin.Context) {
	var coordinates model.CoordinatesAllSchemes
	var coordinatesOfRouter model.CoordinatesPoints
	if err := c.BindJSON(&coordinates); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	logrus.Info("start data from front-end ", coordinates)

	distance, err := service.CalculationOfValues(coordinates)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	file, header, err := c.Request.FormFile("upload")
	logrus.Info(file)
	filename := header.Filename
	fmt.Println(header.Filename)
	out, err := os.Create(filename + ".png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	name := service.DrawImage("2", filename, distance, coordinatesOfRouter.X, coordinatesOfRouter.Y)
	logrus.Info(name)
	//logrus.Info("response date for front-end ", responseCoordinates)
	//
	//c.JSON(http.StatusOK, responseCoordinates)
}

func (h Handler) saveData(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	logrus.Info(header)
}

func (h Handler) loadData(c *gin.Context) {

}
