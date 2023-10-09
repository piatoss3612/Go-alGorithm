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
	input := make([]int, n+1)
	for i := 1; i <= n; i++ {
		input[i] = scanInt()
	}

	inc := make([]int, n+1)
	dec := make([]int, n+1)

	/*
		가장 긴 바이토닉 부분수열

		예제 입력:
		10
		1 5 2 1 4 3 4 5 2 1

		증가하는 dp:
		[0 1 2 2 1 3 3 4 5 2 1]

		감소하는 dp:
		[0 1 5 2 1 4 3 3 3 2 1]

		예제 출력:
		7

		(슬라이스의 1에서부터 i까지 증가하는 가장 긴 부분수열의 길이) +
		(n에서부터 i까지 역순으로 감소하는 부분수열의 길이) - 1이 최댓값이 되는 i를 찾는다
	*/

	for i := n; i >= 1; i-- {
		dec[i] = 1
		for j := i + 1; j <= n; j++ {
			if input[j] < input[i] {
				dec[i] = max(dec[i], dec[j]+1)
			}
		}
	}

	ans := 0

	for i := 1; i <= n; i++ {
		inc[i] = 1
		for j := i - 1; j >= 1; j-- {
			if input[j] < input[i] {
				inc[i] = max(inc[i], inc[j]+1)
			}
		}
		ans = max(ans, inc[i]+dec[i]-1)
	}
	fmt.Fprintln(writer, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
