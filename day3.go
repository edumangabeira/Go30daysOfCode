package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)
	printIfWeird(N)
}

func printIfWeird(N int32) {

	var even bool = N%2 == 0
	var odd bool = N%2 != 0
	if odd {
		fmt.Printf("Weird")
	} else {
		if N < 1 || N > 100 {
			return
		} else {
			if even && N <= 5 {
				fmt.Printf("Not Weird")
			} else {
				if even && N <= 20 {
					fmt.Printf("Weird")
				} else {
					if even && N > 20 {
						fmt.Printf("Not Weird")
					}
				}
			}
		}
	}
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
