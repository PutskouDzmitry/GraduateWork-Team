package main

import (
	"github.com/sirupsen/logrus"
	"math"
)

type point struct {
	x float64
	y float64
}

func main() {
	x1 := point{x: 5, y: 10}
	x2 := point{
		x: 10,
		y: 15,
	}
	// x side
	xDef := x1.x - x2.x
	if xDef < 0 {
		f := xDef
		xDef = -f
	}
	// y side
	yDef := x1.y - x2.y
	if yDef < 0 {
		f := xDef
		yDef = -f
	}
	// theory Pifagora
	mathS := math.Sqrt(math.Pow(xDef, 2) + math.Pow(yDef, 2))
	logrus.Info(mathS)
	// get Angle: cos / katet
	angle := xDef / mathS
	logrus.Info(angle)
}
