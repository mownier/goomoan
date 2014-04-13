// @filename controller.go
// @author mownier

package goomoan

type Action struct {
	RequestMethod string
}

type Module struct {
	Name string
	Actions[] Action
}

func (m Module) RegisterAction() {

}
