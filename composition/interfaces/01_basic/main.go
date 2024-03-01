package _1_basic

import "fmt"

// Reader interface
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Writer interface
type Writer interface {
	Write(p []byte) (n int, err error)
}

// ReadWriter interface, composed of Reader and Writer
type ReadWriter interface {
	Reader
	Writer
}

// Concrete type
type File struct {
	// File has no fields in this example
}

// Implementing Reader interface for File
func (f File) Read(p []byte) (n int, err error) {
	// Simplified example: pretend to read into p
	fmt.Println("Reading from file")
	return len(p), nil // Assume full buffer is read
}

// Implementing Writer interface for File
func (f File) Write(p []byte) (n int, err error) {
	// Simplified example: pretend to write p
	fmt.Println("Writing to file")
	return len(p), nil // Assume full buffer is written
}

func main() {
	var f File

	// f is a ReadWriter because it implements both Read and Write
	var rw ReadWriter = f

	rw.Read(nil)
	rw.Write(nil)
}
