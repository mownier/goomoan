// @filename main.go
// @author mownier

package goomoan

import ( 
	"fmt"
)

func main() {
	fmt.Println("Welcome to Goomoan v0.0.1")
	var api = new(Goomoan)
	api.Start(5353);
}
