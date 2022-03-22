package api

import (
	"fmt"
	"github.com/PutskouDzmitry/GraduateWork/server/pkg/model"
	"github.com/PutskouDzmitry/GraduateWork/server/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

func (h Handler) calculationOfValues(c *gin.Context) {
	var routers []model.RouterSettings
	if err := c.BindJSON(&routers); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	logrus.Info("start data from front-end ", routers)

	file, header, err := c.Request.FormFile("upload")
	logrus.Info(file)
	filename := header.Filename
	fmt.Println(header.Filename)
	out, err := os.Create(filename + ".png")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	distances := make([]float64, 10, 10)
	for _, router := range routers {
		distance, err := service.CalculationOfValues(router)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		distances = append(distances, distance)
	}
	draw := service.NewDrawImage(routers, distances, filename)
	err = draw.DrawOnImage()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	logrus.Info("response date for front-end ", responseCoordinates)

	c.JSON(http.StatusOK, responseCoordinates)
}

func (h Handler) saveData(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	logrus.Info(header)
}

func (h Handler) loadData(c *gin.Context) {

}
