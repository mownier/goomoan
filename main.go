// @filename main.go
// @author mownier

package main

import ( 
	"fmt"
	"goomoan/goomoan"
)

func main() {
	fmt.Println("Welcome to Goomoan v0.0.1")
	app := new(goomoan.Application)
	app.Start(5454)
}
