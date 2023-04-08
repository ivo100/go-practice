// max int sizes
package main

import "fmt"

const (
	UintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
	MaxInt   = 1<<(UintSize-1) - 1        // 1<<31 - 1 or 1<<63 - 1
	MinInt   = -MaxInt - 1                // -1 << 31 or -1 << 63
	MaxUint  = 1<<UintSize - 1            // 1<<32 - 1 or 1<<64 - 1
)

func main() {
	fmt.Printf("UintSize \t%24d \n", UintSize)
	fmt.Printf("  MinInt \t%24d \n", MinInt)
	fmt.Printf("  MaxInt \t%24d \n", MaxInt)
	fmt.Printf(" MaxUint \t%24d \n", uint64(MaxUint))
	// this will overflow int
	//fmt.Printf("MaxUint %d \n", MaxUint)
}

/*
UintSize                              64
  MinInt            -9223372036854775808
  MaxInt             9223372036854775807
 MaxUint            18446744073709551615

9,223,372,036,854,775,807
9 mil bil
*/
