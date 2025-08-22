package main

import (
	"fmt"

	"github.com/ppond454/race-to-n-core/core"
)

func main() {
	const n = 10
	test := core.New(n)
	fmt.Println("test", test)
}
