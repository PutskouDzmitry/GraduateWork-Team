package service

import (
	"github.com/PutskouDzmitry/GraduateWork/server/pkg/model"
	"github.com/fogleman/gg"
	"image"
	"image/color"
	"math"
)

type drawImage struct {
	fileName             string
	coordinatesOfRouters []model.RouterSettings
}

func NewDrawImage(coordinatesOfRouters []model.RouterSettings, fileName string) *drawImage {
	return &drawImage{
		coordinatesOfRouters: coordinatesOfRouters,
		fileName:             fileName,
	}
}

var (
	path2     = "./test_pictures/floor.png"
	n         = 16
	koofStone = 0.6
)

func (d drawImage) DrawOnImage() error {
	im, err := gg.LoadPNG(d.fileName)
	if err != nil {
		return err
	}

	var rotation float64 = 20
	angle := 2 * math.Pi / float64(n)
	rotation -= math.Pi / 2
	ctx := gg.NewContextForImage(im)
	var rNew float64
	var rPromej float64
	for j := 0; j < 8; j++ {
		for a := 0; a < len(d.coordinatesOfRouters); a++ {
			ctx.NewSubPath()
			x, y, r, err := d.getXYR(a)
			colorAndRangeShape := NewColorAndRadius(r)
			if err != nil {
				return err
			}
			for i := 0; i <= n; i++ {
				r = colorAndRangeShape[j].Radius
				rNew = r
				rPromej = r
				colorPixels := colorAndRangeShape[j].Color
				a := angle * float64(i)
				for h := 0; float64(h) < r; h++ {
					xH := x + float64(h)*math.Cos(a)
					yH := y + float64(h)*math.Sin(a)
					colorCheck, _ := detectColorOfPixel(im, xH, yH)
					if !colorCheck {
						rT := getRadius(x, y, xH, yH)
						if rT < rNew {
							rNew = rT + (rPromej-rT)*koofStone
							rPromej = rNew
							h += 5
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
				as := gg.NewSolidPattern(color.Black)
				ctx.SetStrokeStyle(as)
			}
			ctx.SetLineWidth(5)
			ctx.FillPreserve()
			ctx.Stroke()
		}
	}
	ctx.SavePNG("gradient-conic.png")
	return nil
}

func (d drawImage) getXYR(i int) (float64, float64, float64, error) {
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
