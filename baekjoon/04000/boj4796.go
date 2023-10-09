package bj4796

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
	cnt := 1
	for scanner.Scan() {
		if scanner.Text() == "0 0 0" {
			break
		}
		s := strings.Split(scanner.Text(), " ")
		l, _ := strconv.Atoi(s[0])
		p, _ := strconv.Atoi(s[1])
		v, _ := strconv.Atoi(s[2])
		result := maxDay(l, p, v)
		fmt.Fprintf(writer, "Case %d: %d\n", cnt, result)
		cnt++
	}
}

func maxDay(l, p, v int) int {
	r := v % p
	d := v / p
	if r > l {
		r = l
	}
	return (l * d) + r
}
