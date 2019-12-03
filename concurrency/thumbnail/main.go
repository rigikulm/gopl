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

// Make Thumbnails for the specified files in parallel.
// Have the inner go routine report its completion to the outer
// go routine by sending an event.
func makeThumbnails(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			fmt.Println("Processing ", f)
			// Ignoring errors
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}

	// Wait for go routines to complete
	for range filenames {
		<-ch
	}
}

// Original Version
//func makeThumbnails(filenames []string) {
//	for _, f := range filenames {
//		thumb, err := thumbnail.ImageFile(f)
//		if err != nil {
//			log.Print(err)
//			continue
//		}
//		fmt.Println(thumb)
//	}
//}

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
