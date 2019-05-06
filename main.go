package main

import (
	"fmt"
	"os"
)

func main() {
	name := "Unknown"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	fmt.Println(fmt.Sprintf("Oh, hi %s!", name))
}
