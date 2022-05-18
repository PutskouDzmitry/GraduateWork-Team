package service

import (
	"fmt"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/model"
	"github.com/fogleman/gg"
	"github.com/sirupsen/logrus"
	"math"
)

type drawImageToMigrator struct {
	filePathInput        string
	filePathOutput       string
	coordinatesOfRouters []model.RoutersSettingForMigrator
}

func NewDrawImageToMigrator(filePathInput, filePathOutput string, coordinatesOfRouters []model.RoutersSettingForMigrator) *drawImageToMigrator {
	return &drawImageToMigrator{
		filePathInput:        filePathInput,
		filePathOutput:       filePathOutput,
		coordinatesOfRouters: coordinatesOfRouters,
	}
}

type valueOfPowerOnPoint struct {
	coordinates model.CoordinatesPoints
	router      model.RouterSettingForMigrator
}

type radiusOfRouter struct {
	coordinates model.CoordinatesPoints
	radius      float64
}

func (d drawImageToMigrator) AcrylicMigrator() error {
	return nil
}

//find min value of powers on point
func findMinPower(powers []float64, routers model.RoutersSettingForMigrator) valueOfPowerOnPoint {
	var min float64 = 1000
	for _, power := range powers {
		if power < min {
			min = power
		}
	}
	coordinate := routers.Coordinates
	for _, value := range routers.RoutersSettingsMigration {
		if value.Power == min {
			return valueOfPowerOnPoint{
				coordinates: coordinate,
				router:      value,
			}
		}
	}
	return valueOfPowerOnPoint{}
}

//findMaxPower in one point
func findMaxPower(powers []float64, routers model.RoutersSettingForMigrator) valueOfPowerOnPoint {
	var max float64 = -1000
	for _, power := range powers {
		if power > max {
			max = power
		}
	}
	coordinate := routers.Coordinates
	for _, value := range routers.RoutersSettingsMigration {
		if value.Power == max {
			return valueOfPowerOnPoint{
				coordinates: coordinate,
				router:      value,
			}
		}
	}
	return valueOfPowerOnPoint{}
}

func (d drawImageToMigrator) FluxMigrator() error {
	powers := make([]float64, 0, 10)
	powersMin := make([]float64, 0, 10)
	minPowers := make([]valueOfPowerOnPoint, 0, 10)
	maxPowers := make([]valueOfPowerOnPoint, 0, 10)
	for _, value := range d.coordinatesOfRouters {
		for _, valueOfPoint := range value.RoutersSettingsMigration {
			powers = append(powers, valueOfPoint.Power)
		}
		maxPowers = append(maxPowers, findMaxPower(powers, value))

		maxPowerOnPoint := findMaxPower(powers, value)
		for _, value := range d.coordinatesOfRouters {
			for _, valueOfPoint := range value.RoutersSettingsMigration {
				if valueOfPoint.MAC == maxPowerOnPoint.router.MAC {
					powersMin = append(powersMin, valueOfPoint.Power)
				}
			}
		}
		minPowers = append(minPowers, findMinPower(powersMin, value))
		powers = make([]float64, 0, 10)
		powersMin = make([]float64, 0, 10)
	}

	if len(minPowers) == len(maxPowers) {
		logrus.Info("ok")
	} else {
		logrus.Info("not ko")
	}

	distance := make([]radiusOfRouter, 0, 10)
	for i, value := range maxPowers {
		distance = append(distance, radiusOfRouter{
			coordinates: value.coordinates,
			radius:      getRadius(value.coordinates.X, value.coordinates.Y, minPowers[i].coordinates.X, minPowers[i].coordinates.Y),
		})
	}
	err := d.drawWifiOnMap(distance)
	if err != nil {
		logrus.Error(err)
	}
	return nil
}

func (d drawImageToMigrator) TelephoneMigrator() error {
	return nil
}

func (d drawImageToMigrator) drawWifiOnMap(data []radiusOfRouter) error {
	im, err := gg.LoadPNG(d.filePathInput)
	arrayCoordinatesOfPoint := detectColor(im)
	if err != nil {
		return fmt.Errorf("error with load png file: %w", err)
	}
	radii := make([]float64, 0, len(data)+1)
	for _, value := range data {
		radii = append(radii, value.radius)
	}
	//for i, _ := range data {
	//	radii = append(radii, float64(i + value.radius))
	//}
	rotation -= math.Pi / 2
	ctx := gg.NewContextForImage(im)
	var rNew float64
	var checkSignal float64
	//draw all rings
	for j := 0; j < 7; j++ {
		//отрисовка по одному кругу покрытия каждого роутера
		for line := 0; line < len(data); line++ {
			ctx.NewSubPath()
			x, y, r := data[line].coordinates.X, data[line].coordinates.Y, radii[line]
			colorAndRangeShape := NewColorAndRadius(r)
			chooseColor := colorAndRangeShape[j].Color
			//отрисовка по линиям
			for i := 0; i <= n; i++ {
				r = colorAndRangeShape[j].Radius
				rNew = r
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
					checkSignal = -1
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
	logrus.Info(d.filePathOutput)
	ctx.SavePNG(d.filePathOutput)
	return nil
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
