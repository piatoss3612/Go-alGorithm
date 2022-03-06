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
	n := scanInt()

	/*
		n이 조 단위라 슬라이스에 할당할 수 없기 때문에
		다이나믹 프로그래밍으로 풀기 어려운 문제
		따라서 반복되는 규칙을 찾아서 문제를 풀어야 한다

		1개: SK
		2개: CY
		3개: SK
		4개: SK
		5개: SK
		6개: SK
		7개: CY
		...

		이후로 7개씩 반복
	*/

	if n%7 == 0 || n%7 == 2 {
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
