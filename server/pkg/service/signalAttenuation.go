package service

import (
	"github.com/sirupsen/logrus"
	"image"
)

type attenuation struct {
	name        string
	pixel       Pixel
	coefficient float64
}

func initCoefficient() []attenuation {
	coeff := make([]attenuation, 1, 10)
	coeff = append(coeff, attenuation{
		name:        "window without",
		pixel:       Pixel{R: 204, G: 242, B: 255},
		coefficient: 0.7,
	})
	coeff = append(coeff, attenuation{
		name:        "window with",
		pixel:       Pixel{R: 0, G: 172, B: 230},
		coefficient: 0.5,
	})
	coeff = append(coeff, attenuation{
		name:        "wood",
		pixel:       Pixel{R: 138, G: 73, B: 40},
		coefficient: 0.3,
	})
	coeff = append(coeff, attenuation{
		name:        "between",
		pixel:       Pixel{R: 248, G: 191, B: 0},
		coefficient: 0.15,
	})
	coeff = append(coeff, attenuation{
		name:        "main",
		pixel:       Pixel{R: 0, G: 153, B: 51},
		coefficient: 0.1,
	})
	coeff = append(coeff, attenuation{
		name:        "fish",
		pixel:       Pixel{R: 0, G: 230, B: 230},
		coefficient: 0.001,
	})
	return coeff
}

func getCoefficient(pixel Pixel) float64 {
	for _, value := range initCoefficient() {
		if pixel.R == value.pixel.R && pixel.G == value.pixel.G && pixel.B == value.pixel.B {
			return value.coefficient
		}
	}
	return 0
}

func signalAttenuation(img image.Image, x, y float64) float64 {
	pixel, err := getPixels(img, float64(x), float64(y))
	if err != nil {
		logrus.Fatal(err)
	}

	coeff := getCoefficient(pixel)

	if coeff == -1 {
		logrus.Error("something wrong in getCoefficient")
	}
	if coeff == 0 || coeff == -1 {
		return 2
	}
	return coeff
}
