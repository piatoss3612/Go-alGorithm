package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	inDegree []int    // 진입 차수
	items    [][]item // 각 부품을 만들기 위해 필요한 부품들에 대한 정보
	n, m     int
)

// 부품 정보
type item struct {
	part int
	need int
}

// 메모리: 924KB
// 시간: 4ms
// 위상 정렬과 다이나믹 프로그래밍을 이용한 문제 풀이
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()

	inDegree = make([]int, n+1)
	items = make([][]item, n+1)

	// x를 만들기 위해 y가 k개 필요
	var x, y, k int
	for i := 1; i <= m; i++ {
		x, y, k = scanInt(), scanInt(), scanInt()
		inDegree[y] += 1                        // 하위 부품 y의 진입 차수 증가
		items[x] = append(items[x], item{y, k}) // 상위 부품 x를 만들기 위해 필요한 부품 정보 추가
	}

	TopologicalSort() // 위상 정렬 실행
}

// 위상 정렬
func TopologicalSort() {
	dp := make([]int, n+1) // 필요한 부품의 수를 메모이제이션할 슬라이스

	queue := make([]int, 0, n+1)

	// 진입 차수가 0인 부품, 즉 완성품 또는 상위 부품부터 탐색
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		var temp item
		for i := 0; i < len(items[front]); i++ {
			temp = items[front][i]

			// 다른 상위 부품을 구성하는데 front가 포함되지 않는 경우
			if dp[front] == 0 {
				dp[temp.part] += temp.need
			} else { // front가 dp[front]개 만큼 필요하고 front를 만드는데 temp.part가 temp.need개 필요한 경우
				dp[temp.part] += dp[front] * temp.need
			}

			inDegree[temp.part] -= 1      // 하위 부품의 진입 차수 감소
			if inDegree[temp.part] == 0 { // 하위 부품의 진입 차수가 0이 됬을 경우
				queue = append(queue, temp.part)
			}
		}
	}

	// 기본 부품인지 판별하는 방법
	// 기본 부품을 구성하기 위한 하위 부품이 없는 경우
	// 즉, len(items[i])가 0이라면 기본 부품이라고 할 수 있다
	for i := 1; i <= n; i++ {
		if len(items[i]) == 0 {
			fmt.Fprintln(writer, i, dp[i]) // 완성품을 만들기 위해 필요한 부품 i와 개수를 출력
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
