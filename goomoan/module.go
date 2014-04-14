// @filename module.go
// @author mownier

package goomoan

type Route struct {
	Action string
	Path   string
}

func NewRoute(action, path string) *Route {
	return &Route{
			Action: action,
			Path: path,
		}
}

type Module struct {
	Name   string
	Routes map[string]*Route
}
