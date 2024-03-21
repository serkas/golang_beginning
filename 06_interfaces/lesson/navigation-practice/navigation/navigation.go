package navigation

import (
	"navigation-practice/primitives"
)

type Router interface {
	GetRouteInstructions(s, f primitives.Point) []string
}

type Navigator struct {
	routers map[string]Router
}

func NewNavigator() *Navigator {
	n := &Navigator{
		routers: make(map[string]Router),
	}

	return n
}

func (n *Navigator) AddRoutingType(routingType string, r Router) {
	n.routers[routingType] = r
}

func (n *Navigator) Navigate(s, f primitives.Point, terrain string) []string {
	return n.routers[terrain].GetRouteInstructions(s, f)
}
