package service

import (
	"fmt"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/model"
	"github.com/fogleman/gg"
)

var (
	// basic values for router
	numberOfChannels       = 1
	signalLossTransmitting = 0
	signalLossReceiving    = 0
	scale                  = 9
	com                    = 10
)

func ValidationOfPlaceRouter(filePath string, routers []model.RouterSettings) error {
	im, err := gg.LoadPNG(filePath)
	if err != nil {
		return err
	}
	for _, value := range routers {
		for j := 0; j < 600; j++ {
			if value.CoordinatesOfRouter.X+float64(j) >= 600 {
				return fmt.Errorf("out of border")
			}
			pixel, err := getPixels(im, value.CoordinatesOfRouter.X+float64(j), value.CoordinatesOfRouter.Y)
			if err != nil {
				return nil
			}
			if pixel.G != 255 && pixel.B != 255 {
				break
			}
		}
		for j := 0; j < 600; j++ {
			if value.CoordinatesOfRouter.X-float64(j) <= 0 {
				return fmt.Errorf("out of border")
			}
			pixel, err := getPixels(im, value.CoordinatesOfRouter.X-float64(j), value.CoordinatesOfRouter.Y)
			if err != nil {
				return nil
			}
			if pixel.G != 255 && pixel.B != 255 {
				break
			}
		}
		for j := 0; j < 400; j++ {
			if value.CoordinatesOfRouter.Y+float64(j) >= 400 {
				return fmt.Errorf("out of border")
			}
			pixel, err := getPixels(im, value.CoordinatesOfRouter.X, value.CoordinatesOfRouter.Y+float64(j))
			if err != nil {
				return nil
			}
			if pixel.G != 255 && pixel.B != 255 {
				break
			}
		}
		for j := 0; j < 400; j++ {
			if value.CoordinatesOfRouter.Y-float64(j) <= 0 {
				return fmt.Errorf("out of border")
			}
			pixel, err := getPixels(im, value.CoordinatesOfRouter.X+float64(j), value.CoordinatesOfRouter.Y-float64(j))
			if err != nil {
				return nil
			}
			if pixel.G != 255 && pixel.B != 255 {
				break
			}
		}
	}
	return nil
}

func GenerateFullPathOfFileToMap(path, userId string) string {
	return path + userId + "-map.png"
}

func GenerateFullPathOfFileToFlux(path, userId string) string {
	return path + userId + "-flux.png"
}

func GenerateFullPathOfFileToAcrylic(path, userId string) string {
	return path + userId + "-acrylic.png"
}

func GenerateFullPathOfFileToMobile(path, userId string) string {
	return path + userId + "-mobile.png"
}

func GenerateFullPathOfFileForSaveOrigin(path, userId string) string {
	return path + userId + "-saveOrigin.png"
}

func GenerateFullPathOfFileForSaveNotOrigin(path, userId string) string {
	return path + userId + "-saveNotOrigin.png"
}
