package main
import ("fmt"; "bufio"; "os")


func main() {
    reader := bufio.NewReader(os.Stdin)
    inputString, _ := reader.ReadString('\n')
    fmt.Printf("Hello, World.\n%s", inputString)
}
