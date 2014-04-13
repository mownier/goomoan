// @filename main.go
// @author mownier

package main

import ( 
	"fmt"
	"goomoan/core"
	"goomoan/modules"
)

func main() {
	fmt.Println("Welcome to Goomoan v0.0.1")
	s := new(core.Server)	
	s.AddModule(modules.NewPostModule())
	s.Start(5454)
}
