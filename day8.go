package main
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "strconv"
)

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

func checkPhonebook(data map[string]int, reader *bufio.Reader){
    name := strings.Fields(readLine(reader))[0]
    if _, exists := data[name]; exists{
            fmt.Printf("%s=%d\n", name, data[name])
        }else{
            fmt.Printf("Not found\n")
        }
}

func buildPhonebook(data *map[string]int, reader *bufio.Reader){
    arr := strings.Fields(readLine(reader))
    name := arr[0]
    tempData := *data
    if len(arr) == 2{
        number, err := strconv.Atoi(arr[1])
        tempData[name] = number
        checkError(err)
    }
    *data = tempData
}


func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
    testCases, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    data := make(map[string]int)
    for i:=0; i<int(testCases); i++{
        buildPhonebook(&data, reader)
    }
    for i:=0; i<int(testCases); i++{
        checkPhonebook(data, reader)
    }
}
