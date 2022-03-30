package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	s1, s2 := scanString(), scanString()

	solve(s1, s2)
	fmt.Fprintln(writer, 0)
}

// s2의 마지막이 A인 경우와 시작이 B인 경우를 동시에 고려해야 한다
func solve(s1, s2 string) {
	if len(s1) == len(s2) {
		if s1 == s2 {
			fmt.Fprintln(writer, 1)
			writer.Flush()
			os.Exit(0)
		}
	} else {
		if s2[len(s2)-1] == 'A' {
			solve(s1, s2[:len(s2)-1])
		}
		if s2[0] == 'B' {
			tmp := []rune(s2)
			n := len(tmp)
			for i := 0; i < n/2; i++ {
				tmp[i], tmp[n-1-i] = tmp[n-1-i], tmp[i]
			}
			solve(s1, string(tmp)[:len(tmp)-1])
		}
	}
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}
