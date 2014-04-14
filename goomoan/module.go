// @filename moduel.go
// @author mownier

package goomoan

type Route struct {
	Action string
	Path   string
}

func NewRoute(m, p string) *Route {
	r := new(Route)
	r.Action = m
	r.Path = p
	return r
}

type Module struct {
	Name   string
	Routes map[string]*Route
}
