package api

import (
	"github.com/PutskouDzmitry/GraduateWork/server/pkg/model"
	"github.com/PutskouDzmitry/GraduateWork/server/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h Handler) calculationOfValues(c *gin.Context) {
	var coordinates model.CoordinatesAllSchemes
	if err := c.BindJSON(&coordinates); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	logrus.Info("start data from front-end ", coordinates)

	responseCoordinates, err := service.CalculationOfValues(coordinates)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	logrus.Info("response date for front-end ", responseCoordinates)

	c.JSON(http.StatusOK, responseCoordinates)
}

func (h Handler) saveData(c *gin.Context) {

}

func (h Handler) loadData(c *gin.Context) {

}
