package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N, M, H int
	blocks  [51][]int // 각 학생이 가지고 있는 블록의 높이
	dp      [1001]int // dp[i]: 높이가 i인 탑을 만드는 경우의 수
)

const MOD = 10007

// 난이도: Gold 4
// 메모리: 936KB
// 시간: 4ms
// 다이나믹 프로그래밍
// 재귀 함수로 풀려했으나 시간 초과 발생, 3중 반복문으로 푸는 과정에서 모듈로 연산을 빠트려서 틀림
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanLines)
	line := scanLine()
	N, M, H = line[0], line[1], line[2]
	for i := 1; i <= N; i++ {
		blocks[i] = scanLine()
	}

	dp[0] = 1 // 단일 블록으로 탑을 만드는 경우를 세기 위해 dp[0]를 1로 초기화

	for i := 1; i <= N; i++ {
		// 누적 합을 제거하기 위해 역순으로 진행
		for j := H; j >= 0; j-- {
			// 높이가 j인 탑을 만들 수 없는 경우
			if dp[j] == 0 {
				continue
			}

			// 높이가 j인 탑을 만들 수 있는 경우
			for _, k := range blocks[i] {
				// 높이가 k인 블록을 더했을 때 탑의 높이가 H보다 작거나 같다면
				if j+k <= H {
					dp[j+k] += dp[j] // 높이가 j+k인 탑을 만드는 경우의 수 갱신
					dp[j+k] %= MOD   // 모듈로 연산
				}
			}
		}
	}

	fmt.Fprintln(writer, dp[H]%MOD)
}

func scanLine() []int {
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	nums := make([]int, 0, len(fields))
	for _, f := range fields {
		n, _ := strconv.Atoi(f)
		nums = append(nums, n)
	}
	return nums
}
