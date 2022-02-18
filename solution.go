package square

import "math"

type customInt int

const (
	SidesTriangle customInt = 3
	SidesSquare   customInt = 4
	SidesCircle   customInt = 0
)

func CalcSquare(sideLen float64, sidesNum customInt) float64 {
	var result float64

	if sidesNum == SidesCircle {
		result = math.Pi * sideLen * sideLen
	} else if sidesNum == SidesSquare {
		result = sideLen * sideLen
	} else if sidesNum == SidesTriangle {
		result = math.Sqrt(3) * sideLen * sideLen / 4
	}

	return result
}
