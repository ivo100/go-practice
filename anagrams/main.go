package main

import (
	"bufio"
	"io"
	"strings"
)

func main1() {
	//s := "mom"
	//n := anagrams.SherlockAndAnagrams(s)
	//fmt.Printf("n %v", n)

	//reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	//
	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)
	//
	//defer stdout.Close()
	//
	//writer := bufio.NewWriterSize(stdout, 1024 * 1024)
	//
	//qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	//checkError(err)
	//q := int32(qTemp)
	//for qItr := 0; qItr < int(q); qItr++ {
	//	s := readLine(reader)
	//
	//	result := sherlockAndAnagrams(s)
	//
	//	fmt.Fprintf(writer, "%d\n", result)
	//}
	//writer.Flush()

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
