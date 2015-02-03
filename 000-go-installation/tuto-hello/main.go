// Command tuto-hello is a simple greeting command-line application.
package main

import (
	"fmt"
	"os"
)

func main() {
	switch len(os.Args) {
	case 0:
		panic("not possible")
	case 1:
		fmt.Printf("Hello from go-hep/tutos!\n")
	default:
		for _, arg := range os.Args {
			fmt.Printf("Hello %s!\n", arg)
		}
	}
}
