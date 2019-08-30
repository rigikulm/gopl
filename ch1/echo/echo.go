// Prints the list of arguments passed on the command-line.
package main

import (
	"fmt"
	"os"
)

func echoArgs(args []string) string {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}

	return s
}

func main() {
	s := echoArgs(os.Args[1:])
	fmt.Println(s)
}
