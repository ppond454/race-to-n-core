package main

import (
	"fmt"

	"github.com/ppond454/race-to-n-core/core"
)

func main() {
	foo := core.New()
	if err := foo.SetN(20); err != nil {
		fmt.Println(err)
		return
	}
	// can, err := foo.CanStart()

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	choice, err := foo.GetChoice()
	fmt.Println("test", choice, err)

}
