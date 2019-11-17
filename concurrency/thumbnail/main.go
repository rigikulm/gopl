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

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		thumb, err := thumbnail.ImageFile(f)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
	}
}

func main() {
	var filenames []string

	// Get a list of filenames from stdin
	fmt.Println("Enter a list of .jpg filenames (one per line). Enter Ctrl-D when finished.")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fname := input.Text()
		filenames = append(filenames, fname)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}

	makeThumbnails(filenames)
}
