package main

import "fmt"

type Series struct {
	Data []float64
}

func NewSeries(data []float64) *Series {
	return &Series{Data: data}
}
func (s *Series) IsMonotonic() (inc, dec bool) {
	return IsMonotonic(s.Data)
}
func (s *Series) LocalMinMax() (iMin, iMax []int) {
	return LocalMinMax(s.Data)
}

func IsMonotonic(a []float64) (inc, dec bool) {
	inc = true
	dec = true

	for i := 1; i < len(a); i++ {
		d := a[i] - a[i-1]
		if d < 0 {
			inc = false
		} else {
			if d > 0 {
				dec = false
			}
		}
		if !inc && !dec {
			break
		}
	}
	return inc, dec
}

func Trend(a []float64) (inc, dec int) {

	for i := 1; i < len(a); i++ {
		d := a[i] - a[i-1]
		if d < 0 {
			dec++
		} else {
			if d > 0 {
				inc++
			}
		}
	}
	return inc, dec
}

func LocalMinMax(a []float64) (iMin, iMax []int) {
	//var iMin []int
	//var iMax []int

	for i := 1; i < len(a)-1; i++ {
		if a[i] < a[i-1] && a[i] < a[i+1] {
			iMin = append(iMin, i)
			continue
		}
		if a[i] > a[i-1] && a[i] > a[i+1] {
			iMax = append(iMax, i)
		}
	}
	return iMin, iMax
}

func computeFirstDerivative(arr []float64) []float64 {
	firstDerivative := make([]float64, len(arr)-1)
	for i := 1; i < len(arr); i++ {
		firstDerivative[i-1] = arr[i] - arr[i-1]
	}
	return firstDerivative
}

func computeSecondDerivative(arr []float64) []float64 {
	secondDerivative := make([]float64, len(arr)-2)
	for i := 1; i < len(arr)-1; i++ {
		secondDerivative[i-1] = arr[i+1] - 2*arr[i] + arr[i-1]
	}
	return secondDerivative
}

func SMA(a []float64, window int) []float64 {
	if window > len(a) {
		return nil
	}
	L := len(a)
	sma := make([]float64, L-window+1)
	for i := 0; i <= L-window; i++ {
		sum := 0.0
		for j := 0; j < window; j++ {
			sum += a[i+j]
		}
		sma[i] = sum / float64(window)
	}
	return sma
}

func main() {
	data := []float64{1.0, 2.0, 4.0, 7.0, 11.0}
	firstDerivative := computeFirstDerivative(data)
	secondDerivative := computeSecondDerivative(data)
	fmt.Printf("First Derivative: %v\n", firstDerivative)
	fmt.Printf("Second Derivative: %v\n", secondDerivative)

	win := 3
	sma3 := SMA(data, win)
	fmt.Printf("Moving Averages: %v\n", sma3)
	win = 4
	sma4 := SMA(data, win)
	fmt.Printf("Moving Averages: %v\n", sma4)

	data = []float64{5, 6, 8, 7, 9, 9}
	inc, dec := Trend(data)
	fmt.Printf("Trend: up: %d, down: %d\n", inc, dec)
	imin, imax := LocalMinMax(data)
	for _, x := range imin {
		fmt.Printf("local min: %v\n", x)
	}
	for _, x := range imin {
		fmt.Printf("local min - data[%d] = %f\n", x, data[x])
	}
	for _, x := range imax {
		fmt.Printf("local max - data[%d] = %f\n", x, data[x])
	}

}
