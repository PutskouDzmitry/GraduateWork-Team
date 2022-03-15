package main

import (
	"github.com/fogleman/gg"
	"github.com/sirupsen/logrus"
	"image"
	_ "image/png"
	"math"
)

func main() {
	drawOnImage()
}

var path1 = "./test_pictures/test-1-removebg-preview.png"
var path2 = "./test_pictures/floor.png"

func drawOnImage() {
	im, err := gg.LoadPNG(path2)
	if err != nil {
		logrus.Fatal(err)
	}
	dc := gg.NewContextForImage(im)
	x, y := getPointOfRouter()
	drawShapes(im, dc, x, y, getRadius(), getRadius(), 0, 2*math.Pi)
	dc.Fill()
	dc.SavePNG("out.png")
}

func drawShapes(infile image.Image, dc *gg.Context, x, y, rx, ry, angle1, angle2 float64) {
	const n = 360
	var rotation float64 = 20
	var r float64 = getRadius()
	angle := 2 * math.Pi / float64(n)
	rotation -= math.Pi / 2
	if n%2 == 0 {
		rotation += angle / 2
	}
	dc.NewSubPath()
	var j int64
	for i := 0; i < n; i++ {
		a := rotation + angle*float64(i)
		color, err := detectColorOfPixel(infile, x+r*math.Cos(a), y+r*math.Sin(a))
		if err != nil {
			logrus.Fatal(err)
		}
		if !color && j == 0 {
			dc.SetRGBA255(255, 0, 0, 150)
			j++
		}
		if color && j == 0 {
			dc.SetRGBA255(0, 255, 0, 150)
		}
		dc.LineTo(x+r*math.Cos(a), y+r*math.Sin(a))
	}
	dc.ClosePath()
}

//func drawShapes(infile image.Image, dc *gg.Context, x, y, rx, ry, angle1, angle2 float64) {
//	const n = 360
//	var rotation float64 = 20
//	var r float64 = getRadius()
//	angle := 2 * math.Pi / float64(n)
//	rotation -= math.Pi / 2
//	if n%2 == 0 {
//		rotation += angle / 2
//	}
//	dc.NewSubPath()
//	for i := 0; i < n; i++ {
//		if i == 287 {
//			break
//		}
//		a := rotation + angle*float64(i)
//		color, err := detectColorOfPixel(infile, x+r*math.Cos(a), y+r*math.Sin(a))
//		if err != nil {
//			logrus.Fatal(err)
//		}
//		if !color {
//			j := 0.0
//			for {
//				color, err := detectColorOfPixel(infile, x - j+r*math.Cos(a), y - j+r*math.Sin(a))
//				if err != nil {
//					logrus.Fatal(err)
//				}
//				if color {
//					break
//				}
//				j++
//			}
//			dc.SetRGBA255(255, 0, 0, 255)
//			dc.LineTo(x - j+r*math.Cos(a), y- j+r*math.Sin(a))
//		} else {
//			dc.SetRGBA255(0, 255, 0, 150)
//		}
//		dc.LineTo(x+r*math.Cos(a), y+r*math.Sin(a))
//		dc.FillPreserve()
//	}
//	dc.ClosePath()
//}

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

func getPointOfRouter() (float64, float64) {
	return 550, 250
}

func getRadius() float64 {
	return 20
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
