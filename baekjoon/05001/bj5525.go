package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
이 문제도 결국 bufio.Scanner의 버펴 용량 문제...
*/

// const MaxBuf int = 1000000 // 문자열 크기가 최대 1백만인데 50점 부분점수
// const MaxBuf int = 1000001 // 100점; 메모리: 1904KB, 시간: 20ms
const MaxBuf int = 2000000 // 100점; 메모리: 1932KB, 시간: 12ms

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Buffer(make([]byte, 0, MaxBuf), MaxBuf)
	scanner.Split(bufio.ScanWords)
	n, m := scanInt(), scanInt()
	s := scanBytes()

	ans := 0
	cnt := 0

	for i := 1; i < m-1; i++ {
		if s[i-1] == 73 && s[i] == 79 && s[i+1] == 73 {
			cnt += 1
			if cnt == n {
				cnt -= 1
				ans += 1
			}
			i += 1
		} else {
			cnt = 0
		}
	}
	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
