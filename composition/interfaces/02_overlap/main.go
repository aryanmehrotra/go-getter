package main

import "fmt"

// Closer interface
type Closer interface {
	Close() error
}

// Reader interface
type Reader interface {
	Read(p []byte) (n int, err error)
	Close() error // Closer method is repeated
}

// ReadCloser interface, composed of Reader and Closer (redundant but illustrative)
type ReadCloser interface {
	Reader
	Closer
}

// Concrete type
type NetworkStream struct {
	// NetworkStream has no fields in this example
}

// Implementing Reader interface for NetworkStream
func (ns NetworkStream) Read(p []byte) (n int, err error) {
	fmt.Println("Reading from network stream")
	return 0, nil // Simplified
}

// Implementing Closer interface for NetworkStream (also part of Reader)
func (ns NetworkStream) Close() error {
	fmt.Println("Closing network stream")
	return nil
}

func main() {
	var ns NetworkStream

	// ns is a ReadCloser because it implements both Read and Close
	var rc ReadCloser = ns

	rc.Read(nil)
	rc.Close()
}
