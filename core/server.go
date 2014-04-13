// @filename core.go
// @author mownier

package core

import (
	"fmt"
	"net/http"
)

type Server struct {}

func (s *Server) Start(port int) {
	fmt.Printf("Start serving at port %d...", port)
	portString := fmt.Sprintf(":%d", port)
	http.ListenAndServe(portString, nil)
}

func (s *Server) Abort(rw http.ResponseWriter, statusCode int) {
    rw.WriteHeader(statusCode)
}

func (s *Server) DelegateModuleRequestHandler() {

}

func (s *Server) AddModule(m *Module) {

}




