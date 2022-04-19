package main

import (
	"github.com/fogleman/gg"
	"github.com/sirupsen/logrus"
	"image"
	"image/color"
	"math"
)

var (
	path1     = "1_example_ready.png"
	path2     = "1_example_ready (1).png"
	n         = 16
	koofStone = 0.4
)

type Coordinates struct {
	x0    float64
	y0    float64
	x1    float64
	y1    float64
	stage int64
}

func main() {
	im, err := gg.LoadPNG(path2)
	arrayXY := detectColor(im)
	if err != nil {
		logrus.Fatal(err)
	}

	var x, y, r float64 = 300, 150, 200
	var rotation float64 = 20
	angle := 2 * math.Pi / float64(n)
	rotation -= math.Pi / 2
	colorAndRangeShape := NewColorAndRadius(r)
	ctx := gg.NewContextForImage(im)
	var rNew float64
	var rPromej float64
	for j := 0; j < len(colorAndRangeShape); j++ {
		for a := 0; a < 1; a++ {
			ctx.NewSubPath()
			ctx.DrawCircle(x, y, 10)
			for i := 0; i <= n; i++ {
				r = colorAndRangeShape[j].Radius
				rNew = r
				rPromej = r
				colorPixels := colorAndRangeShape[j].Color
				a := angle * float64(i)
				for h := 0; float64(h) < r; h++ {
					xH := x + float64(h)*math.Cos(a)
					yH := y + float64(h)*math.Sin(a)
					for k := 0; k < len(arrayXY); k++ {
						if float64(int64(xH)) == arrayXY[k].x && float64(int64(yH)) == arrayXY[k].y {
							rT := getRadius(x, y, xH, yH)
							if rT < rNew {
								rNew = rT + (rPromej-rT)*koofStone
								rPromej = rNew
								h += 30
							}
						}
					}
					//colorCheck, _ := detectColorOfPixel(im, xH, yH)
					//if !colorCheck {
					//	rT := getRadius(x, y, xH, yH)
					//	if rT < rNew {
					//		rNew = rT + (rPromej-rT)*koofStone
					//		rPromej = rNew
					//		h += 5
					//	}
					//}
				}
				cosX := x + rNew*math.Cos(a)
				sinY := y + rNew*math.Sin(a)
				if i == 0 {
					ctx.MoveTo(x+rNew, y)
					continue
				}
				ctx.LineTo(cosX, sinY)
				ctx.SetRGBA255(int(colorPixels.R), int(colorPixels.G), int(colorPixels.B), int(colorPixels.A))
				as := gg.NewSolidPattern(color.Black)
				ctx.SetStrokeStyle(as)
			}
			ctx.SetLineWidth(1)
			ctx.FillPreserve()
			ctx.Stroke()
		}
	}
	ctx.SavePNG("gradient-conic.png1.png")
}

type XY struct {
	x float64
	y float64
}

func detectColor(im image.Image) []XY {
	getXYArray := make([]XY, 0, 10)
	getXY := XY{}
	for x := 0; x < 600; x++ {
		for y := 0; y < 400; y++ {
			pixel, err := getPixels(im, float64(x), float64(y))
			if err != nil {
				logrus.Fatal(err)
			}
			if pixel.B != 255 && pixel.R != 255 {
				getXY = XY{
					x: float64(x),
					y: float64(y),
				}
				getXYArray = append(getXYArray, getXY)
			}
		}
	}
	return getXYArray
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
	a := uint8(220)
	var kof2 float64 = 0.5
	var kof3 float64 = 0.4
	var kof4 float64 = 0.3
	var kof5 float64 = 0.25
	var kof6 float64 = 0.2
	var kof7 float64 = 0.15
	var kof8 float64 = 0.1
	colorArr := make([]ColorAndRadius, 9, 11)
	colorArr[0].Color = color.RGBA{A: a, R: 220, G: 0, B: 0}
	colorArr[0].Radius = radius
	colorArr[1].Color = color.RGBA{A: a, R: 223, G: 106, B: 78}
	colorArr[1].Radius = radius * kof2
	colorArr[2].Color = color.RGBA{A: a, R: 227, G: 138, B: 80}
	colorArr[2].Radius = radius * kof3
	colorArr[3].Color = color.RGBA{A: a, R: 234, G: 170, B: 82}
	colorArr[3].Radius = radius * kof4
	colorArr[4].Color = color.RGBA{A: a, R: 190, G: 255, B: 92}
	colorArr[4].Radius = radius * kof5
	colorArr[5].Color = color.RGBA{A: a, R: 140, G: 255, B: 91}
	colorArr[5].Radius = radius * kof6
	colorArr[6].Color = color.RGBA{A: a, R: 110, G: 255, B: 91}
	colorArr[6].Radius = radius * kof7
	colorArr[7].Color = color.RGBA{A: a, R: 92, G: 255, B: 90}
	colorArr[7].Radius = radius * kof8
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
