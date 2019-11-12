// netcat3 - listens on port 8000 and displays anything
// that is sent
package main

import (
	"net"
	"log"
	"io"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	// Anonymous function to handle copying the received data
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} // signal the 'main' goroutine
	}()

	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
