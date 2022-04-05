package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func checkMaxConsecutives(base2number [] string) int32{
    var max int32 = 0
    var max_temp int32 = 0
    for i:=0; i<len(base2number); i++{
        if base2number[i] == string('1') {
            max_temp += 1
        } else{
            max_temp = 0
        }
        if max_temp >= max{
            max = max_temp
        }
    }
    return max
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := strconv.FormatInt(nTemp, 2)
    n_array := strings.Split(n, "")
    fmt.Println(checkMaxConsecutives(n_array))
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
