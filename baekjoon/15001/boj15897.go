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
	N       int
)

// 메모리: 920KB
// 시간: 8ms
// 잘못된 에라토스테네스의 체의 연산횟수가 같은 값이 같은 구간을 구함으로써
// 연산 속도를 O(N)에서 O(N^1/2)로 줄일 수 있다
// 참고: https://ahgus89.github.io/algorithm/Harmonic-Lemma/
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()

	/*
		// 시간복잡도 O(N)으로 시간 초과 발생

		sum := N

		for i := 1; i <= N; i++ {
			sum += (N - 1) / i
		}
	*/

	// 잘못 구현한 에라토스테네스의 체에서
	// j가 1인 경우의 연산 횟수: N
	// i=1인 경우의 연산 횟수: N-1
	// 이것들을 미리 더해놓고 남은 연산횟수를 구한다
	sum := 2*N - 1
	k := 1

	for i := 2; i < N; i = k + 1 {
		k = (N - 1) / ((N - 1) / i) // k는 (N-1)/i만큼의 연산횟수가 필요한 최댓값
		cnt := (N - 1) / i          // 필요한 연산횟수
		sum += (k - i + 1) * cnt    // i~k구간의 연산횟수
	}

	fmt.Fprintln(writer, sum)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
