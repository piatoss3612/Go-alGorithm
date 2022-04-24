package bj9012

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		checkVPS()
	}
}

func checkVPS() {
	scanner.Scan()
	slice := strings.Split(scanner.Text(), "")
	left := 0
	for _, v := range slice {
		if left == 0 && v == ")" {
			fmt.Fprintln(writer, "NO")
			return
		}
		if v == "(" {
			left += 1
		} else if v == ")" && left > 0 {
			left -= 1
		}
	}
	if left == 0 {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
	}
}
