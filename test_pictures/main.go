package main

import (
	"github.com/fogleman/gg"
	"github.com/sirupsen/logrus"
	"github.com/tfriedel6/canvas/sdlcanvas"
	"image/color"
	"math"
)

var (
	path2 = "./test_pictures/floor.png"
	n     = 320
)

func main() {
	im, err := gg.LoadPNG(path2)
	if err != nil {
		logrus.Fatal(err)
	}
	wnd, _, err := sdlcanvas.CreateWindow(1280, 720, "Hello")
	if err != nil {
		panic(err)
	}
	defer wnd.Destroy()

	var x, y, r float64 = 200, 200, 800
	var rotation float64 = 20
	angle := 2 * math.Pi / float64(n)
	rotation -= math.Pi / 2

	ctx := gg.NewContextForImage(im)
	var colorCircle color.RGBA
	logrus.Info(colorCircle)
	for j := 0; j < 1; j++ {
		nr := r / 5
		if j == 0 {
			colorCircle = color.RGBA{R: 255, G: 0, B: 0, A: 254}
		}
		if j == 1 {
			nr += nr
			colorCircle = color.RGBA{G: 255}
		}
		ctx.NewSubPath()
		for i := 0; i < n; i++ {
			a := angle * float64(i)
			cos := math.Cos(a)
			sin := math.Sin(a)
			if i == 0 {
				ctx.MoveTo(x+nr, y)
				continue
			}
			ctx.LineTo(x+nr*cos, y+nr*sin)
			ctx.SetColor(color.Alpha16{0})
		}
		ctx.Fill()
	}
	ctx.SavePNG("gradient-conic.png")
}

func gradient() {
	im, err := gg.LoadPNG(path2)
	if err != nil {
		logrus.Fatal(err)
	}
	c := gg.NewRadialGradient(100, 75, 5, 100, 75, 75)
	c.AddColorStop(0, color.White)
	c.AddColorStop(1, color.RGBA{B: 255, A: 255})
	ctx := gg.NewContextForImage(im)
	ctx.DrawArc(100, 75, 60, 0, 2*math.Pi)
	ctx.SetFillStyle(c)
	ctx.Fill()
	ctx.SavePNG("gradient-conic.png")
}
