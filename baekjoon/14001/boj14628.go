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
	N, M, K int
	X, Y    int
	dp      [101]int
	// i번째 스킬을 조합하는 과정에서 1~i-1번째 스킬을 조합한 결과가 필요,
	//따라서 슬라이딩 윈도우 비슷한 방식으로 단일 dp 라인만으로 결괏값을 구할 수 있다
)

// 메모리: 908KB
// 시간: 4ms
// 다이나믹 프로그래밍, 배낭 문제
// 정확히 적의 체력 M을 깎을 수 있는 스킬의 조합과 그 비용의 최소 필요 마나를 찾는 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M, K = scanInt(), scanInt(), scanInt()

	// 1. N개의 스킬 정보 입력
	for i := 1; i <= N; i++ {
		X, Y = scanInt(), scanInt() // 스킬에 필요한 마나, 체력을 깎는 수치

		// 2. i번째 스킬로 상대 체력 깎기
		// j가 M부터 0으로 줄어드는 순서로 진행하여 중복되는 연산을 배제
		for j := M; j >= 0; j-- {
			temp := X // 현재 스킬 사용에 필요한 마나, 사용할 때마다 필요 마나가 K만큼 늘어난다

			// 2-1. 적의 체력이 전혀 깎이지 않은 상태에서 시작하는 경우
			if j == 0 {
				for k := j; k <= M; {
					// 0에서부터 시작하므로 체력이 깎여 있는지 여부는 체크하지 않는다
					if k+Y <= M {
						// k+Y만큼 체력을 깎아본 적이 없는 경우
						if dp[k+Y] == 0 {
							dp[k+Y] = dp[k] + temp
						} else {
							// k+Y만큼 체력을 깎아본 적이 있는 경우
							dp[k+Y] = min(dp[k+Y], dp[k]+temp)
						}
						// temp와 k값 조정
						temp += K
						k += Y
					} else {
						break
					}
				}
			} else {
				// 2-2. 적의 체력이 j만큼 깎인 상태에서 시작하는 경우
				for k := j; k <= M; {
					// k만큼 적의 체력이 깎여있고 Y만큼 더 깎았을 때 M보다 작거나 같은 경우
					if dp[k] > 0 && k+Y <= M {
						// k+Y만큼 체력을 깎아본 적이 없는 경우
						if dp[k+Y] == 0 {
							dp[k+Y] = dp[k] + temp
						} else {
							// k+Y만큼 체력을 깎아본 적이 있는 경우
							dp[k+Y] = min(dp[k+Y], dp[k]+temp)
						}
						temp += K
						k += Y
					} else {
						break
					}
				}
			}
		}
	}

	fmt.Fprintln(writer, dp[M]) // 적의 체력을 M만큼 깎았을 때의 최소 필요 마나 출력
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
