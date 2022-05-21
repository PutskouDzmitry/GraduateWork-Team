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
	thickness              = 10
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

//func ValidationValues(routers []model.RouterSettings) []model.RouterSettings {
//	newRouters := make([]model.RouterSettings, len(routers), len(routers)+1)
//	for i, value := range routers {
//		newRouters[i].CoordinatesOfRouter = value.CoordinatesOfRouter
//		if value.NumberOfChannels == -1 {
//			newRouters[i].NumberOfChannels = numberOfChannels
//		} else {
//			newRouters[i].NumberOfChannels = value.NumberOfChannels
//		}
//		if value.SignalLossTransmitting == -1 {
//			newRouters[i].SignalLossTransmitting = float64(signalLossTransmitting)
//		} else {
//			newRouters[i].SignalLossTransmitting = value.SignalLossTransmitting
//		}
//		if value.SignalLossReceiving == -1 {
//			newRouters[i].SignalLossReceiving = float64(signalLossReceiving)
//		} else {
//			newRouters[i].SignalLossReceiving = value.SignalLossReceiving
//		}
//		if value.Scale == -1 {
//			newRouters[i].Scale = float64(scale)
//		} else {
//			newRouters[i].Scale = value.Scale
//		}
//		if value.Thickness == -1 {
//			newRouters[i].Thickness = float64(thickness)
//		} else {
//			newRouters[i].Thickness = value.Thickness
//		}
//		if value.COM == -1 {
//			newRouters[i].COM = float64(com)
//		} else {
//			newRouters[i].COM = value.COM
//		}
//	}
//	return newRouters
//}
