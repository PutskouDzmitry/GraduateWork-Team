package api

import (
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/model"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
	"os"
)

type T struct {
	file string `json:"file"`
}

func (h Handler) calculationOfValues(c *gin.Context) {
	var routers []model.RouterSettings
	var t T
	if err := c.BindJSON(&t); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	logrus.Info("start data from front-end ", routers)

	response, e := http.Get(t.file)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()
	out, err := os.Create("filename" + ".png")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	defer out.Close()
	_, err = io.Copy(out, response.Body)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	draw := service.NewDrawImage(routers, "filename"+".png")
	err = draw.DrawOnImage()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	filepath := "http://localhost:8080/file/" + "qwe"
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}

func (h Handler) saveData(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	logrus.Info(header)
}

func (h Handler) loadData(c *gin.Context) {

}
