package main
import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "io"
    "strconv"
)


func SplitByIndex(s []string, size int) ([]string, []string){
    evenSlice := make([]string, size)
    oddSlice := make([]string, size)
    
    for i:=0; i<size; i++{
        if isEven(i){
            evenSlice = append(evenSlice, s[i])
        }else{
            oddSlice = append(oddSlice, s[i])
            
        }
    }
    return evenSlice, oddSlice
}

func concatenateSlices(evenSlice []string, oddSlice []string) string{
    return strings.Join(evenSlice, "") + " " + strings.Join(oddSlice, "")
}

func isEven(x int) bool { 
        return x%2 == 0
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

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
    testCases, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    for i:=0; int64(i)<testCases; i++{
        S := strings.Split(strings.Fields(readLine(reader))[0], "")
        N := len(S)
        evenSlice, oddSlice := SplitByIndex(S, N)
        final_string := concatenateSlices(evenSlice, oddSlice)
        fmt.Println(final_string)
    }
}
