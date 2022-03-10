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

	// n개의 소시지와 m명의 평론가
	// 평론가 1명 당 n/m 개의 소시지를 받을 수 있다
	/*
		예제 입력: 3 4

		프로세스:
		소시지1 == | ======
		소시지2 == | ======
		소시지3 == | ======

		소시지를 1/4, 3/4로 나누어 총 3번의 칼질
		최대 m - 1번, 최소 m - gcd(n, m)의 칼질이 필요
	*/
	fmt.Fprintln(writer, m-gcd(n, m))
}

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
