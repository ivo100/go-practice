package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// bin search of x in sorted slice a
// if found - returns index of first item == x
// if not found - returns index of first item > x
func find(a []int, x int) int {
	// bin search
	n := sort.SearchInts(a, x)
	// warning n == len(a) if x > max
	//log.Printf("bin search %v in %+v - result idx: %v", x, a, n)
	return n
}

// given slice, places next element to sorted array
// shifting array 1 pos to the left
func place(a []int, x int) {
	n := find(a, x)
	// first?
	// replace it
	if n == 0 {
		a[0] = x
		//log.Printf("first a: %+v", a)
		return
	}
	// last?
	l := len(a)
	if n == l {
		copy(a, a[1:l])
		a[l-1] = x
		//log.Printf("last a: %+v", a)
		return
	}
	// x found at pos n
	// shift left 1:n-1
	// replace a[n]
	//found := a[n] == x
	//log.Printf("x %v, n %v, found %v", x, n, found)
	copy(a, a[1:n])
	a[n-1] = x
	//log.Printf("** found %v, a: %+v", found, a)
	return
}

// Complete the activityNotifications function below.
func activityNotifications0(exp []int, td int) int {
	n := 0
	l := len(exp)
	if td >= l {
		return 0
	}
	even := td%2 == 0
	m := td / 2
	tr := make([]int, td)
	// first sort
	copy(tr, exp[0:td+1])
	sort.Ints(tr)
	med2 := 0
	for i := td; i < l; i++ {
		med2 = tr[m]
		if even {
			med2 += tr[m-1]
		} else {
			med2 += med2
		}
		x := exp[i]
		if x >= med2 {
			n++
		}
		// next day
		place(tr, x)
	}
	return n
}

func activityNotifications(exp []int, td int) int {
	n := 0
	l := len(exp)
	if td >= l {
		return 0
	}
	even := td%2 == 0
	m := td / 2
	//log.Printf("exp: %+v", exp)
	log.Printf("td: %v, l %v, m: %v", td, l, m)

	// td 4 -> m 2 med = (x[m-1] + x[m])/2
	// td 5 -> m 2 med = x[m]
	//tr := make([]int, td)
	// first sort
	//copy(tr, exp[0:td+1])
	//log.Printf("first slice %+v", tr)
	//log.Printf("tr len %v", len(tr))
	//sort.Ints(tr)
	//sort.SliceStable(tr, func(i, j int) bool { return tr[i] < tr[j] })
	//log.Printf("first sort %+v", tr)
	med2 := 0
	for i := td; i < l; i++ {
		tr := exp[i-td : i]
		sort.SliceStable(tr, func(i, j int) bool { return tr[i] < tr[j] })
		//log.Printf("exp: %+v", exp)
		med2 = tr[m]
		if even {
			med2 += tr[m-1]
		} else {
			med2 += med2
		}
		x := exp[i]
		//log.Printf("tr %+v med2 %v : x %v", tr, med2, x)
		//log.Printf("med2 %v : x %v", med2, x)
		if x >= med2 {
			n++
			log.Printf("== n %v med2 %v x %v", n, med2, x)
			//if n > 1691 {
			//log.Printf("== n %v tr %+v med2 %v x %v", n, tr, med2, x)
			//}
		}
		// next day
		//place(tr, x)
	}
	return n
}

func main() {
	file, err := os.Open("fraudulent-activity/big-test")
	//file, err := os.Open("fraudulent-activity/test")
	if err != nil {
		w, _ := os.Getwd()
		panic("Input file? " + err.Error() + " " + w)
	}
	reader := bufio.NewReaderSize(file, 1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	stdout := os.Stdout
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nd := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nd[0], 10, 64)
	checkError(err)
	n := int(nTemp)

	dTemp, err := strconv.ParseInt(nd[1], 10, 64)
	checkError(err)
	d := int(dTemp)

	expenditureTemp := strings.Split(readLine(reader), " ")

	var expenditure []int

	for i := 0; i < int(n); i++ {
		expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
		checkError(err)
		expenditureItem := int(expenditureItemTemp)
		expenditure = append(expenditure, expenditureItem)
	}

	result := activityNotifications(expenditure, d)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
