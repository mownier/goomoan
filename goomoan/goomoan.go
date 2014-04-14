// @filename goomoan.go
// @author mownier

package goomoan

import (
	"fmt"
	"net/http"
)

type Application struct {}

func (a *Application) registerModules() {

}

func (a *Application) Start(port int) {
	a.registerModules()
	fmt.Printf("Start serving at port %d...", port)
	portString := fmt.Sprintf(":%d", port)
	http.ListenAndServe(portString, nil)
}

func (a *Application) Abort(rw http.ResponseWriter, statusCode int) {
    rw.WriteHeader(statusCode)
}
