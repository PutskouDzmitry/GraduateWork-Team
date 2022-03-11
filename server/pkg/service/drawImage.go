package service

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/sirupsen/logrus"
	"image"
	"math"
	"os"
)

func DrawImage(id, fileName string, distance float64, x, y float64) string {
	im, err := gg.LoadPNG(fileName)
	if err != nil {
		logrus.Fatal(err)
	}
	dc := gg.NewContextForImage(im)
	infile, err := os.Open(fileName)
	img, _, err := image.Decode(infile)
	if err != nil {
		logrus.Fatal(err)
	}
	drawShapes(img, dc, x, y, distance)
	dc.Fill()
	fileNameResult := fmt.Sprint(id + ".png")
	dc.SavePNG(fileNameResult)
	return fileNameResult
}

func drawShapes(infile image.Image, dc *gg.Context, x, y, radius float64) {
	const n = 360
	var rotation float64 = 20
	var r float64 = 200
	angle := 2 * math.Pi / float64(n)
	rotation -= math.Pi / 2
	if n%2 == 0 {
		rotation += angle / 2
	}
	dc.NewSubPath()
	for i := 0; i < n; i++ {
		a := rotation + angle*float64(i)
		_, err := detectColorOfPixel(infile, x+r*math.Cos(a), y+r*math.Sin(a))
		if err != nil {
			logrus.Fatal(err)
		}
		//if !color {
		//	logrus.Info("1qwe")
		//	logrus.Info(x+r*math.Cos(a), y+r*math.Sin(a))
		//	logrus.Info("2qwe")
		//	//dc.SetRGBA255(255, 0, 0, 150)
		//	//dc.LineTo(x+r*math.Cos(a), y+r*math.Sin(a))
		//	continue
		//}
		logrus.Info(x+r*math.Cos(a), y+r*math.Sin(a))
		dc.LineTo(x+r*math.Cos(a), y+r*math.Sin(a))
		dc.SetRGBA255(0, 255, 0, 150)
	}
	dc.ClosePath()
}

func detectColorOfPixel(img image.Image, x, y float64) (bool, error) {
	pixel, err := getPixels(img, x, y)
	if err != nil {
		return false, err
	}
	if pixel.R == 0 && pixel.G == 0 && pixel.B == 0 && pixel.A == 0 {
		return false, nil
	}
	return true, nil
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
