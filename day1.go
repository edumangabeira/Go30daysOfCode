package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	var int_var uint64
	var double_var float64
	var string_var string
	scanner.Scan()
	int_var, _ = strconv.ParseUint(scanner.Text(), 10, 64)
	scanner.Scan()
	double_var, _ = strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	string_var = scanner.Text()
	fmt.Printf("%d\n", i+int_var)
	fmt.Printf("%.1f\n", d+double_var)
	fmt.Println(s + string_var)

}
