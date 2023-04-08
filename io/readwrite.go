package main

import (
	//"golang.org/x/tour/reader"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
https://cs.opensource.google/search?q=Read%5C(%5Cw%2B%5Cs%5C%5B%5C%5Dbyte%5C)&ss=go%2Fgo

https://tour.golang.org/methods/21

The Go standard library contains many implementations of this interface, including files, network connections, compressors, ciphers, and others.

The io.Reader interface has a Read method:

func (T) Read(b []byte) (n int, err error)

*/

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (rdr MyReader) Read(a []byte) (int, error) {
	for i := 0; i < len(a); i++ {
		a[i] = 'A'
	}
	return len(a), nil
}

/*
Exercise: Readers
Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
*/

/*
https://tour.golang.org/methods/23
Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream
by applying the rot13 substitution cipher to all alphabetical characters.
*/

type rot13Reader struct {
	r io.Reader
}

// Input	ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
// Output	NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm
func rot13(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		d := c + 13
		if d > 'Z' {
			d = 'A' + d - 'Z' - 1
		}
		return d
	}
	if c >= 'a' && c <= 'z' {
		d := c + 13
		if d > 'z' {
			d = 'a' + d - 'z' - 1
		}
		return d
	}
	return c
}

func (rdr rot13Reader) Read(a []byte) (int, error) {
	n, err := rdr.r.Read(a)
	if err != nil {
		return n, err
	}
	// replacing each one by the letter 13 places further along in the alphabet, wrapping back 	   // to the beginning if necessary
	for i := 0; i < n; i++ {
		a[i] = rot13(a[i])
	}
	return n, nil
}

func main() {
	sr := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := sr.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	//reader.Validate(MyReader{})

	s := strings.NewReader("XYZ")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	println("")
	s2 := strings.NewReader("KLM")
	r2 := rot13Reader{s2}
	io.Copy(os.Stdout, &r2)
	println("")
}
