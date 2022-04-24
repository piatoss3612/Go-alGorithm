package bj2908

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Scan()
	nums := strings.Split(scanner.Text(), " ")
	a := strings.Split(nums[0], "")
	b := strings.Split(nums[1], "")
	for i := 2; i >= 0; i-- {
		cmp := strings.Compare(a[i], b[i])
		if cmp == 1 {
			printReverse(a)
			break
		} else if cmp == -1 {
			printReverse(b)
			break
		} else {
			continue
		}
	}
}

func printReverse(s []string) {
	result := ""
	for _, v := range s {
		result = v + result
	}
	fmt.Fprintln(writer, result)
}
