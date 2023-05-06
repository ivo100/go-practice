package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*

In this challenge, you will determine whether a string is funny or not. To determine whether a string is funny, create a copy of the string in reverse e.g. . Iterating through each string, compare the absolute difference in the ascii values of the characters at positions 0 and 1, 1 and 2 and so on to the end. If the list of absolute differences is the same for both strings, they are funny.
Determine whether a give string is funny. If it is, return Funny, otherwise return Not Funny.
Example

The ordinal values of the charcters are .  and the ordinals are . The absolute differences of the adjacent elements for both strings are , so the answer is Funny.
Function Description
Complete the funnyString function in the editor below.
funnyString has the following parameter(s):
string s: a string to test
Returns
string: either Funny or Not Funny
Input Format
The first line contains an integer , the number of queries.
The next  lines each contain a string, .
Constraints


Sample Input
STDIN   Function
-----   --------
2       q = 2
acxz    s = 'acxz'
bcxz    s = 'bcxz'
Sample Output
Funny
Not Funny
Explanation
Let  be the reverse of .
Test Case 0:
,
Corresponding ASCII values of characters of the strings:
 and
For both the strings the adjacent difference list is [2, 21, 2].
Test Case 1:
,
Corresponding ASCII values of characters of the strings:
 and
The difference list for string  is [1, 21, 2] and for string  is [2, 21, 1].


 * Complete the 'funnyString' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
*/

func funnyString(s string) string {
	// Write your code here
	//var d byte
	l := len(s)
	a := make([]byte, l)
	b := make([]byte, l)

	for i, j := 0, l-1; i < l; {
		a[i] = s[i]
		b[i] = s[j]
		i++
		j--
	}
	diff := func(x, y byte) int {
		d := int(x) - int(y)
		if d < 0 {
			return -d
		} else {
			return d
		}
	}
	for i := 0; i < l-1; i++ {
		d1 := diff(a[i], a[i+1])
		d2 := diff(b[i], b[i+1])
		if d1 != d2 {
			return "Not Funny"
		}
	}
	return "Funny"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := funnyString(s)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
