package routing

import (
	"fmt"
	"navigation-practice/primitives"
)

type CityRouter struct {
}

func (r CityRouter) GetRouteInstructions(s, f primitives.Point) (instructions []string) {
	if f.X != s.X {
		instructions = append(instructions, fmt.Sprintf("move along X for %d", f.X-s.X))
	}

	if f.Y != s.Y {
		instructions = append(instructions, fmt.Sprintf("move along Y for %d", f.Y-s.Y))
	}
	return instructions
}
