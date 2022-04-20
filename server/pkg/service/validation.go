package service

import (
	"fmt"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/model"
	"github.com/fogleman/gg"
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

func GenerateFullPathOfFile(path, userId string) string {
	return path + userId + ".png"
}
