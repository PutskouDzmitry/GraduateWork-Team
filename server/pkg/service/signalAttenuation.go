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
		name:        "",
		pixel:       Pixel{},
		coefficient: 0,
	})
	coeff = append(coeff, attenuation{
		name:        "",
		pixel:       Pixel{},
		coefficient: 0,
	})
	coeff = append(coeff, attenuation{
		name:        "",
		pixel:       Pixel{},
		coefficient: 0,
	})
	coeff = append(coeff, attenuation{
		name:        "",
		pixel:       Pixel{},
		coefficient: 0,
	})
	coeff = append(coeff, attenuation{
		name:        "",
		pixel:       Pixel{},
		coefficient: 0,
	})
	coeff = append(coeff, attenuation{
		name:        "",
		pixel:       Pixel{},
		coefficient: 0,
	})
	coeff = append(coeff, attenuation{
		name:        "",
		pixel:       Pixel{},
		coefficient: 0,
	})
	coeff = append(coeff, attenuation{
		name:        "",
		pixel:       Pixel{},
		coefficient: 0,
	})
	coeff = append(coeff, attenuation{
		name:        "",
		pixel:       Pixel{},
		coefficient: 0,
	})
	coeff = append(coeff, attenuation{
		name:        "",
		pixel:       Pixel{},
		coefficient: 0,
	})
	coeff = append(coeff, attenuation{
		name:        "",
		pixel:       Pixel{},
		coefficient: 0,
	})
	coeff = append(coeff, attenuation{
		name:        "",
		pixel:       Pixel{},
		coefficient: 0,
	})
	coeff = append(coeff, attenuation{
		name:        "",
		pixel:       Pixel{},
		coefficient: 0,
	})
	coeff = append(coeff, attenuation{
		name:        "",
		pixel:       Pixel{},
		coefficient: 0,
	})
	coeff = append(coeff, attenuation{
		name:        "",
		pixel:       Pixel{},
		coefficient: 0,
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
		return 0.6
	}
	return coeff
}

func detectSizeOfAttenuation() {

}
