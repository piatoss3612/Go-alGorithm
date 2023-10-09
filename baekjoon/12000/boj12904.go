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
	s1 := scanString()
	s2 := scanString()

	for len(s2) > len(s1) {
		if s2[len(s2)-1] == 'A' {
			s2 = s2[:len(s2)-1]
		} else {
			s2 = s2[:len(s2)-1]
			s2 = reverse(s2)
		}
	}

	if s1 == s2 {
		fmt.Fprintln(writer, 1)
	} else {
		fmt.Fprintln(writer, 0)
	}
}

// string 타입으로 연산하는 것보다 훨씬 빠르게 연산 결과를 얻을 수 있다
func reverse(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}
