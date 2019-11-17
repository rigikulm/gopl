// main.go
//
// Generates thumbnail images for list of JPEG filenames provided on
// the command-line.
// Example:
//		go run main.go
//		foo.jpg
//		bar.jpg
//		CTRL-D (ends the program)
//
// This code is based on the 'thumbnail' example in chapter 8 of
// "The Go Programming Language" by A. Donovan and B. Kernighan.
//
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"lemler.org/thumbnail/internal/pkg/thumbnail"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fname := input.Text()
		thumb, err := thumbnail.ImageFile(fname)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
		//log.Print("Got: ", fname)
		//fmt.Println("Input: ", fname)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}
