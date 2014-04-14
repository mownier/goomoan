// @filename main.go
// @author mownier

package main

import ( 
	"fmt"
	// "goomoan/goomoan"
	"goomoan/modules"
	"encoding/json"
)

func main() {
	fmt.Println("Welcome to Goomoan v0.0.1")
	// app := new(goomoan.Application)
	// app.Start(5454)
	p := new(modules.PostModule)
	status, data := p.SendPost(nil)
	js, err := json.Marshal(data)
	fmt.Println(status)
	if (err != nil) {
		fmt.Println("Something went wrong while marshalling")
	} else {
		fmt.Printf("json: %s\n",js)
	}
}
