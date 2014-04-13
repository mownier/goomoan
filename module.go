// @filename controller.go
// @author mownier

package goomoan

import (
	"strings"
)

type Action struct {
	Name string
	RequestMethod string
	RequestRoute string
}

type Module struct {
	Name string
	Actions map[string]*Action
}

func (m *Module) RegisterAction(a *Action) {
	name := strings.ToLower(a.Name)
	m.Actions[name] = a
}
