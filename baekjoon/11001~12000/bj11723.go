package bj11723

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	m := scanInt()
	s := make([]bool, 21)
	for i := 0; i < m; i++ {
		op := scanString()
		switch op {
		case "add":
			{
				x := scanInt()
				if s[x] == false {
					s[x] = true
				}
			}
		case "remove":
			{
				x := scanInt()
				s[x] = false
			}
		case "check":
			{
				x := scanInt()
				if s[x] {
					fmt.Fprintln(writer, 1)
				} else {
					fmt.Fprintln(writer, 0)
				}
			}
		case "toggle":
			{
				x := scanInt()
				if s[x] {
					s[x] = false
				} else {
					s[x] = true
				}
			}
		case "all":
			{
				for i := 1; i <= 20; i++ {
					s[i] = true
				}
			}
		case "empty":
			s = make([]bool, 21)
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}
