package main

import (
	"github.com/fogleman/gg"
	"github.com/sirupsen/logrus"
	"image"
	"image/color"
	"math"
)

var (
	path2     = "./test_pictures/example_floorplan.png"
	n         = 32
	koofStone = 0.85
)

func main() {
	im, err := gg.LoadPNG(path2)
	if err != nil {
		logrus.Fatal(err)
	}

	var x, y, r float64 = 900, 900, 700
	var x1, y1, r1 float64 = 1200, 1200, 700
	var rotation float64 = 20
	angle := 2 * math.Pi / float64(n)
	rotation -= math.Pi / 2
	colorAndRangeShape := NewColorAndRadius(r)
	ctx := gg.NewContextForImage(im)
	var rNew float64
	var rPromej float64
	for a := 0; a < 2; a++ {
		if a == 1 {
			x = x1
			y = y1
			r = r1
		}
		for j := 0; j < len(colorAndRangeShape); j++ {
			ctx.NewSubPath()
			for i := 0; i < n; i++ {
				r = colorAndRangeShape[j].Radius
				rNew = r
				rPromej = r
				colorPixels := colorAndRangeShape[j].Color
				a := angle * float64(i)
				for h := 0; float64(h) < r; h++ {
					xH := x + float64(h)*math.Cos(a)
					yH := y + float64(h)*math.Sin(a)
					colorCheck, err := detectColorOfPixel(im, xH, yH)
					if err != nil {
						logrus.Error(err)
					}
					if !colorCheck {
						rT := getRadius(x, y, xH, yH)
						if rT < rNew {
							rNew = rT + (rPromej-rT)*koofStone
							rPromej = rNew
							h += 40
						}

					}
				}
				cosX := x + rNew*math.Cos(a)
				sinY := y + rNew*math.Sin(a)
				if i == 0 {
					ctx.MoveTo(x+rNew, y)
					continue
				}
				ctx.LineTo(cosX, sinY)
				ctx.SetRGBA255(int(colorPixels.R), int(colorPixels.G), int(colorPixels.B), int(colorPixels.A))
			}
			ctx.Fill()
		}
	}
	ctx.SavePNG("gradient-conic.png")
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
	a := uint8(150)
	var kof1 float64 = 0.5
	var kof2 float64 = 0.2
	var kof3 float64 = 0.15
	var kof4 float64 = 0.1
	var kof5 float64 = 0.06
	var kof6 float64 = 0.03
	var kof7 float64 = 0.02
	var kof8 float64 = 0.01
	colorArr := make([]ColorAndRadius, 9, 11)
	colorArr[0].Color = color.RGBA{R: 255, A: a}
	colorArr[0].Radius = radius
	colorArr[1].Color = color.RGBA{A: a, R: 254, G: 100, B: 1}
	colorArr[1].Radius = radius * kof1
	colorArr[2].Color = color.RGBA{A: a, R: 253, G: 190, B: 11}
	colorArr[2].Radius = radius * kof2
	colorArr[3].Color = color.RGBA{A: a, R: 234, G: 253, B: 20}
	colorArr[3].Radius = radius * kof3
	colorArr[4].Color = color.RGBA{A: a, R: 140, G: 252, B: 20}
	colorArr[4].Radius = radius * kof4
	colorArr[5].Color = color.RGBA{A: a, G: 252, B: 30}
	colorArr[5].Radius = radius * kof5
	colorArr[6].Color = color.RGBA{A: a, G: 250, B: 143}
	colorArr[6].Radius = radius * kof6
	colorArr[7].Color = color.RGBA{B: a, A: 120}
	colorArr[7].Radius = radius * kof7
	colorArr[8].Color = color.RGBA{B: a, A: 120}
	colorArr[8].Radius = radius * kof8
	return colorArr
}

func detectColorOfPixel(img image.Image, x, y float64) (bool, error) {
	pixel, err := getPixels(img, x, y)
	if pixel.R == 0 && pixel.G == 0 && pixel.B == 0 {
		return false, err
	}
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
