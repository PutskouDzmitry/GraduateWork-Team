package api

import (
	"fmt"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/model"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
	"os"
)

func (h Handler) calculationOfValues(c *gin.Context) {
	var routers []model.RouterSettings
	if err := c.BindJSON(&routers); err != nil {
		//newErrorResponse(c, http.StatusBadRequest, err.Error())
		//return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	out, err := os.Create("public/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	draw := service.NewDrawImage(routers, filename)
	err = draw.DrawOnImage()
	if err != nil {
		//newErrorResponse(c, http.StatusBadRequest, err.Error())
		//return
	}
	filepath := "http://localhost:8080/file/" + filename
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}

func (h Handler) saveData(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	logrus.Info(header)
}

func (h Handler) loadData(c *gin.Context) {

}
