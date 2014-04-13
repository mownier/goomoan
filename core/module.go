// @filename controller.go
// @author mownier

package core

type Route struct {
	Method string
	Path   string
}

func NewRoute(m, p string) *Route {
	r := new(Route)
	r.Method = m
	r.Path = p
	return r
}

type Module struct {
	Routes map[string]*Route
}

