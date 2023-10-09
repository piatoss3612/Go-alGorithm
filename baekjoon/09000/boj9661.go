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

// 메모리: 916KB
// 시간: 4ms
// 돌겜 이론
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()

	/*
		돌이 1개 ~15개일 때 게임의 결과: 돌을 1개 또는 4개 가져갈 수 있다

		1 2 3 4 5 6 7 8 9 10 11 12 13 14 15
		S C S S C S C S S C  S  C  S  S  C

		이처럼 5를 단위로 S C S S C가 반복된다

		돌을 16개를 가져갈 수 있는 경우는?

		상근이는 돌을 1개 | 4개 | 16개를 가져갈 수 있는데
		돌 16: 어떻게 가져가더라도 상근이의 필승
		돌 17: 어떻게 가져가더라도 상근이의 필패
		돌 18: 돌을 1개 가져가면 상근이의 필승
		돌 19: 돌을 4개 가져가면 상근이의 필승
		돌 20: 어떻게 가져가더라도 상근이의 필패

		...

		상근이는 돌을 1개 | 4개 | 16개 | 64개를 가져갈 수 있는데
		돌 64: 64개 또는 4개를 가져가면 상근이의 필승
		돌 65: 어떻게 가져가더라도 상근이의 필패
		돌 66: 어떻게 가져가더라도 상근이의 필승
		돌 67: 어떻게 가져가더라도 상근이의 필패
		돌 68: 16개 또는 1개를 가져가면 상근이의 승리
		돌 69: 64개 또는 4개를 가져가면 상근이의 승리
		돌 70: 어떻게 가져가더라도 상근이의 필패

		가져가는 돌의 수 4^x가 쓸데없이 커질 필요도 없이
		1개 또는 4개를 가져가는 선에서 게임의 결과는 확정되므로
		5개를 단위로 SCSSC가 반복된다는 것을 확인할 수 있다
	*/

	if n%5 == 2 || n%5 == 0 {
		fmt.Fprintln(writer, "CY")
	} else {
		fmt.Fprintln(writer, "SK")
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
