package routing

import (
	"fmt"
	"math"
	"navigation-practice/primitives"
)

type LandRouter struct {
}

func (r LandRouter) GetRouteInstructions(s, f primitives.Point) (instructions []string) {
	mag, angle := cartesianToPolar(f.X-s.X, f.Y-s.Y)
	return []string{
		fmt.Sprintf("move for %.2f with angle %.2f", mag, angle*180/math.Pi),
	}
}

func cartesianToPolar(x, y int) (magnitude, angle float64) {
	magnitude = math.Sqrt(float64(x*x) + float64(y*y))
	angle = math.Atan2(float64(y), float64(x))
	if angle < 0 {
		angle += 2 * math.Pi
	}
	return magnitude, angle
}
