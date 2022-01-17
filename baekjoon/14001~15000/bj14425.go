package main

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
	n, m := scanInt(), scanInt()
	s := make(map[string]bool)
	for i := 0; i < n; i++ {
		str := scanStr()
		s[str] = true
	}
	cnt := 0
	for i := 0; i < m; i++ {
		str := scanStr()
		_, ok := s[str]
		if ok {
			cnt += 1
		}
	}
	fmt.Fprintln(writer, cnt)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanStr() string {
	scanner.Scan()
	return scanner.Text()
}
