// @filename core.go
// @author mownier

package goomoan

import (
	"fmt"
	"net/http"
)

type Goomoan struct {}

// Start server
func (goomoan *Goomoan) Start(port int) {
	portString := fmt.Sprintf(":%d", port)
	http.ListenAndServe(portString, nil)
	fmt.Println("Start serving...")
}

// Abort server
func(goomoan *Goomoan) Abort() {
	
}



