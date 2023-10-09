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
	N, M    int
	sum     int   // 누적 합
	pmod    []int // pmod[i]: 구간 1~j까지의 합을 M으로 나눈 나머지가 i인 경우의 수
)

// 난이도: Gold 3
// 메모리: 6592KB
// 시간: 168ms
// 분류: 누적 합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	pmod = make([]int, M)

	for i := 1; i <= N; i++ {
		sum += scanInt() // 누적 합 갱신
		pmod[sum%M] += 1 // 나머지의 경우의 수 갱신
	}
}

func Solve() {
	/*
		1. 1에서 시작하여 크기가 1씩 증가하는 구간의 크기가 N(1~N)이 될 때까지의 각 구간합을 M으로 나눈 나머지가 i(0<=i<M)인 경우의 수를 pmod[i]라고 하자.
		2. pmod[i]는 나머지가 동일한 구간의 개수를 의미하며, 이 구간들은 서로 겹치지 않는다.
		3. 구간이 서로 겹치지 않으므로 나머지가 동일한 구간의 개수(pmod[i])가 2개 이상인 경우, 그 중에 겹치지 않게 2개를 고를 수 있다.
			3-1. 방금 고른 2개의 구간 중 구간합이 큰 쪽을 A, 작은 쪽을 B라고 하자.
			3-2. 구간합이 크다는 것은 A의 구간이 끝나는 위치(a)가 B의 구간이 끝나는 위치(b)보다 항상 큼을 의미한다. (구간이 겹치지 않으므로)
			3-3. 따라서 A의 구간(1~a)에서 B의 구간(1~b)을 뺀 새로운 구간(b+1~a)의 합은 A의 구간의 합에서 B의 구간의 합을 뺀 것과 같으며 M으로 나눈 나머지는 0이 된다.
		4. pmod[i]개 중에 2개를 고르는 조합의 수는 (pmod[i] * (pmod[i] - 1)) / 2로 계산할 수 있다.
		5. 모든 i에 대해 3~4번을 반복한다.
		6. pmod[0](구간 1~j의 합을 M으로 나눈 나머지가 0인 경우)는 이미 완성된 구간이므로 결과에 누적해서 더해준다.
	*/

	ans := pmod[0]
	for i := 0; i < M; i++ {
		if pmod[i] >= 2 {
			ans += (pmod[i] * (pmod[i] - 1)) / 2
		}
	}

	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
