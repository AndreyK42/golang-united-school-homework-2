package square

import (
	"math"
)

type FigureSidesNumber int

const (
	SidesCircle FigureSidesNumber = iota
	_
	_
	SidesTriangle = iota
	SidesSquare   = iota
)

func CalcSquare(sideLen float64, sidesNum FigureSidesNumber) float64 {
	switch sidesNum {
	case SidesCircle:
		return math.Pi * sideLen * sideLen
	case SidesTriangle:
		return 0.25 * math.Sqrt(3) * sideLen * sideLen
	case SidesSquare:
		return sideLen * sideLen
	default:
		return 0
	}
}
