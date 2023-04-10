package main

import (
	"fmt"
	"unsafe"
)

// detect endianness

func main() {
	var i int32 = 0x04030201
	p := unsafe.Pointer(&i)
	b := (*[4]byte)(p)
	if b[0] == 0x01 {
		fmt.Println("Little endian architecture")
	} else {
		fmt.Println("Big endian architecture")
	}
}
