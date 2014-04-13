// @filename core.go
// @author mownier

package goomoan

import (
	"fmt"
	"net/http"
)

type Goomoan struct {}

func (goomoan *Goomoan) Start(port int) {
	portString := fmt.Sprintf(":%d", port)
	http.ListenAndServe(portString, nil)
	fmt.Println("Start serving...")
}

func(goomoan *Goomoan) Abort() {
	
}



