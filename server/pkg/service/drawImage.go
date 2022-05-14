package service

import (
	"fmt"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/model"
	"github.com/fogleman/gg"
	"github.com/sirupsen/logrus"
	"image"
	"image/color"
	"math"
)

type drawImage struct {
	filePathInput        string
	filePathOutput       string
	coordinatesOfRouters []model.RouterSettings
}

func NewDrawImage(coordinatesOfRouters []model.RouterSettings, filePathInput string, filePathOutput string) *drawImage {
	return &drawImage{
		coordinatesOfRouters: coordinatesOfRouters,
		filePathInput:        filePathInput,
		filePathOutput:       filePathOutput,
	}
}

var (
	n                = 16
	rotation float64 = 20
	angle            = 2 * math.Pi / float64(n)
)

func (d drawImage) DrawOnImage() error {
	im, err := gg.LoadPNG(d.filePathInput)
	arrayCoordinatesOfPoint := detectColor(im)
	if err != nil {
		return fmt.Errorf("error with load png file: %w", err)
	}
	radii := make([]float64, 0, len(d.coordinatesOfRouters))
	for _, value := range d.coordinatesOfRouters {
		radius, err := CalculationOfValues(value)
		if err != nil {
			return err
		}
		radius /= value.Scale
		radii = append(radii, radius)
	}

	rotation -= math.Pi / 2
	ctx := gg.NewContextForImage(im)
	var rNew float64
	var checkSignal float64
	//draw all rings
	for j := 0; j < 7; j++ {
		//отрисовка по одному кругу покрытия каждого роутера
		for a := 0; a < len(d.coordinatesOfRouters); a++ {
			ctx.NewSubPath()
			x, y, r := d.coordinatesOfRouters[a].CoordinatesOfRouter.X, d.coordinatesOfRouters[a].CoordinatesOfRouter.Y, radii[a]
			colorAndRangeShape := NewColorAndRadius(r)
			chooseColor := colorAndRangeShape[j].Color
			//отрисовка по линиям
			for i := 0; i <= n; i++ {
				r = colorAndRangeShape[j].Radius
				rNew = r
				//rPromej = r
				a := angle * float64(i)
				// расчет длины сигнала(поиск препядствий)
				for h := 0; float64(h) < r; h++ {
					xH := x + float64(h)*math.Cos(a)
					yH := y + float64(h)*math.Sin(a)
					for k := 0; k < len(arrayCoordinatesOfPoint); k++ {
						if float64(int64(xH)) == arrayCoordinatesOfPoint[k].x && float64(int64(yH)) == arrayCoordinatesOfPoint[k].y {
							attenuationOfSignal := signalAttenuation(im, xH, yH)
							if attenuationOfSignal == 2 {
								continue
							}
							if checkSignal != attenuationOfSignal {
								rNew = float64(h) + (r-float64(h))*attenuationOfSignal
								checkSignal = attenuationOfSignal
							}
						}
					}
					checkSignal = 3
				}
				cosX := x + rNew*math.Cos(a)
				sinY := y + rNew*math.Sin(a)
				detectOutPositionOfSignal(im, cosX, sinY)
				if i == 0 {
					ctx.MoveTo(x+rNew, y)
					continue
				}
				ctx.LineTo(cosX, sinY)
			}
			ctx.SetRGBA255(int(chooseColor.R), int(chooseColor.G), int(chooseColor.B), int(chooseColor.A))
		}
		ctx.SetLineWidth(0)
		ctx.FillPreserve()
		ctx.Stroke()
	}
	ctx.SavePNG(d.filePathOutput)
	return nil
}

func (d drawImage) getCoordinatesOfPointR(i int) (float64, float64, float64, error) {
	r, err := CalculationOfValues(d.coordinatesOfRouters[i])
	if err != nil {
		return -1, -1, -1, err
	}
	return d.coordinatesOfRouters[i].CoordinatesOfRouter.X, d.coordinatesOfRouters[i].CoordinatesOfRouter.Y, r, nil
}

func getRadius(x0, y0, x1, y1 float64) float64 {
	var x0x1 float64
	var y0y1 float64
	if x1-x0 >= 0 {
		x0x1 = x1 - x0
	} else {
		x0x1 = x0 - x1
	}
	if y1-y0 >= 0 {
		y0y1 = y1 - y0
	} else {
		y0y1 = y0 - y1
	}
	return math.Sqrt(math.Pow(x0x1, 2) + math.Pow(y0y1, 2))
}

type ColorAndRadius struct {
	Color  color.RGBA
	Radius float64
}

func NewColorAndRadius(radius float64) []ColorAndRadius {
	a := uint8(180)
	var kof2 float64 = 0.5
	var kof3 float64 = 0.4
	var kof4 float64 = 0.3
	var kof5 float64 = 0.25
	var kof6 float64 = 0.2
	var kof7 float64 = 0.15
	colorArr := make([]ColorAndRadius, 9, 11)
	colorArr[0].Color = color.RGBA{A: a, R: 100, G: 123, B: 251}
	colorArr[0].Radius = radius
	colorArr[1].Color = color.RGBA{A: a, R: 107, G: 184, B: 240}
	colorArr[1].Radius = radius * kof2
	colorArr[2].Color = color.RGBA{A: a, R: 120, G: 245, B: 242}
	colorArr[2].Radius = radius * kof3
	colorArr[3].Color = color.RGBA{A: a, R: 115, G: 246, B: 105}
	colorArr[3].Radius = radius * kof4
	colorArr[4].Color = color.RGBA{A: a, R: 237, G: 247, B: 123}
	colorArr[4].Radius = radius * kof5
	colorArr[5].Color = color.RGBA{A: a, R: 243, G: 187, B: 115}
	colorArr[5].Radius = radius * kof6
	colorArr[6].Color = color.RGBA{A: a, R: 239, G: 117, B: 109}
	colorArr[6].Radius = radius * kof7
	return colorArr
}

func getPixels(img image.Image, x, y float64) (Pixel, error) {
	return rgbaToPixel(img.At(int(x), int(y)).RGBA()), nil
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

type Pixel struct {
	R int
	G int
	B int
	A int
}

type CoordinatesOfPoint struct {
	x float64
	y float64
}

func detectColor(im image.Image) []CoordinatesOfPoint {
	getCoordinatesOfPointArray := make([]CoordinatesOfPoint, 0, 10)
	getCoordinatesOfPoint := CoordinatesOfPoint{}
	for x := 0; x < 600; x++ {
		for y := 0; y < 400; y++ {
			pixel, err := getPixels(im, float64(x), float64(y))
			if err != nil {
				logrus.Fatal(err)
			}
			if pixel.B != 255 && pixel.R != 255 {
				getCoordinatesOfPoint = CoordinatesOfPoint{
					x: float64(x),
					y: float64(y),
				}
				getCoordinatesOfPointArray = append(getCoordinatesOfPointArray, getCoordinatesOfPoint)
			}
		}
	}
	return getCoordinatesOfPointArray
}

func detectOutPositionOfSignal(im image.Image, x, y float64) {
	for j := 0; j < 600; j++ {
		if x+float64(j) >= 600 {
			//logrus.Info("out of border")
			break
		}
		pixel, err := getPixels(im, x+float64(j), y)
		if err != nil {
			//logrus.Info("out of border")
			break
		}
		if pixel.G != 255 && pixel.B != 255 {
			break
		}
	}
	for j := 0; j < 600; j++ {
		if x-float64(j) <= 0 {
			//logrus.Info("out of border")
			break
		}
		pixel, err := getPixels(im, x-float64(j), y)
		if err != nil {
			//logrus.Info("out of border")
			break
		}
		if pixel.G != 255 && pixel.B != 255 {
			break
		}
	}
	for j := 0; j < 400; j++ {
		if y+float64(j) >= 400 {
			//logrus.Info("out of border")
			break
		}
		pixel, err := getPixels(im, x, y+float64(j))
		if err != nil {
			//logrus.Info("out of border")
			break
		}
		if pixel.G != 255 && pixel.B != 255 {
			break
		}
	}
	for j := 0; j < 400; j++ {
		if y-float64(j) <= 0 {
			//logrus.Info("out of border")
			break
		}
		pixel, err := getPixels(im, x+float64(j), y-float64(j))
		if err != nil {
			//logrus.Info("out of border")
			break
		}
		if pixel.G != 255 && pixel.B != 255 {
			break
		}
	}
}
