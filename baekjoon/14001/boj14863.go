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
	N, K    int
	options []*Option        // option[i]: i-1 도시에서 i 도시로 이동수단의 정보
	dp      [101][100001]int // dp[i][j]: i 도시에서 시작하여 (K-j) 만큼의 시간이 남았을 때 모금할 수 있는 최대 모금액
)

// 이동수단 정보
type Option struct {
	// 도보
	WalkTime   int
	WalkIncome int
	// 자전거
	BikeTime   int
	BikeIncome int
}

const INF = -987654321 // 모금액의 최댓값 10^8보다 절댓값이 큰 음수

// 메모리: 61456KB
// 시간: 108ms
// 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, K = scanInt(), scanInt()
	options = make([]*Option, N+1)

	// 이동수단 정보 입력
	for i := 1; i <= N; i++ {
		options[i] = &Option{
			scanInt(), scanInt(), scanInt(), scanInt(),
		}
	}

	// 재귀 함수 호출: 서울에서 시작하여 K 만큼의 시간이 남아있을 때의 모금액의 최댓값 구하기
	ans := solve(0, 0)
	fmt.Fprintln(writer, ans)
}

func solve(pos, time int) int {
	// 기저 사례1: 사용할 수 있는 시간 이상이 소비된 경우
	if time > K {
		return INF
	}

	// 기저 사례2: 마지막 도시까지 성공적으로 도착한 경우
	if pos == N {
		return 0
	}

	ret := &dp[pos][time]
	if *ret != 0 {
		return *ret
	}

	// 중요!: N번째 도시까지 도달하는 경우에만 최댓값을 비교하도록
	// 걸어서 가는 경우와 자전거를 가는 경우 모두 불가능한 경우에는 음수값을 반환하도록 해야 한다
	*ret = INF

	// 현재 도시에서 다음 도시까지
	*ret = max(*ret, solve(pos+1, time+options[pos+1].WalkTime)+options[pos+1].WalkIncome) // 걸어서 가는 경우
	*ret = max(*ret, solve(pos+1, time+options[pos+1].BikeTime)+options[pos+1].BikeIncome) // 자전거를 타고 가는 경우

	return *ret
}

var (
	dp2 [100001]int
)

// 메모리: 2480KB
// 시간: 24ms
// 다이나믹 프로그래밍, 슬라이딩 윈도우
func main2() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, K = scanInt(), scanInt()

	dp2[0] = 1 // 첫번째 입력값에 대한 연산을 수행하기 위해 0번 인덱스 값을 1로 초기화

	for i := 1; i <= N; i++ {
		wt, wi, bt, bi := scanInt(), scanInt(), scanInt(), scanInt()
		// 중복 연산을 배제하기 위해 j를 K부터 0까지 역순으로 진행
		for j := K; j >= 0; j-- {
			if dp2[j] != 0 {
				// 도보로 이동하는 경우
				if j+wt <= K {
					dp2[j+wt] = max(dp2[j+wt], dp2[j]+wi)
				}
				// 자전거로 이동하는 경우
				if j+bt <= K {
					dp2[j+bt] = max(dp2[j+bt], dp2[j]+bi)
				}
				dp2[j] = 0 // N번째 도시까지 성공적으로 이동한 경우의 모금액만 비교해야 하므로 이동을 마친 기존의 값은 제거
			}
		}
	}

	ans := 0
	// 최댓값 비교
	for i := 0; i <= K; i++ {
		ans = max(ans, dp2[i])
	}
	fmt.Fprintln(writer, ans-1) // 최댓값에서 dp[0]에 임의로 부여한 1만큼을 빼준 값이 모금액의 최댓값
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
