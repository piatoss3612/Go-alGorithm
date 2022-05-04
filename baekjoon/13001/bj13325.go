package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	n       int
	tree    [21][1048577]int // 높이가 20일 때 리프들의 수: 2^20 = 1048576
	dp      [21][1048577]int
	sum     = 0 // 가중치의 총합
)

// 1차 풀이 제출
// 메모리: 38840KB
// 시간: 344ms
// func main() {
// 	defer writer.Flush()
// 	scanner.Split(bufio.ScanWords)
// 	n = scanInt()
// 	for i := 1; i <= n; i++ {
// 		for j := 1; j <= int(math.Pow(2, float64(i))); j++ {
// 			tree[i][j] = scanInt()
// 			sum += tree[i][j]
// 		}
// 	}

// 	dp[n] = tree[n]

// 	for i := n; i >= 2; i-- {
// 		for j := 2; j <= int(math.Pow(2, float64(i))); j += 2 {
// 			dp[i-1][j/2] = tree[i-1][j/2] + max(dp[i][j-1], dp[i][j])
// 		}
// 	}

// 	for i := 1; i <= n; i++ {
// 		for j := 2; j <= int(math.Pow(2, float64(i))); j += 2 {
// 			if dp[i][j-1] > dp[i][j] {
// 				sum += dp[i][j-1] - dp[i][j]
// 			} else {
// 				sum += dp[i][j] - dp[i][j-1]
// 			}
// 		}
// 	}
// 	fmt.Fprintln(writer, sum)
// }

// 2차 풀이: 재귀 함수로 변형
// 메모리: 30572KB
// 시간: 284ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n = scanInt()
	for i := 1; i <= n; i++ {
		// i번 높이의 노드들의 개수: int(math.Pow(2, float64(i)))
		for j := 1; j <= int(math.Pow(2, float64(i))); j++ {
			tree[i][j] = scanInt()
			sum += tree[i][j]
		}
	}
	_ = rec(1, 1, 2) // (1,1), (1,2)에서 시작하는 재귀 호출

	fmt.Fprintln(writer, sum)
}

/*
예제 입력:
2
2 2 2 1 1 3

tree:
2 2
2 1 1 3

dp:
4 5
0 0 0 0


예제 출력:
15
*/

/*
최적화, 최솟값을 구하는 방법:
하나의 부모 노드에서 갈라져 나온 두 개의 자식 노드가 각각 가지는 가중치의 합 a, b를 비교하여
a가 b보다 큰 경우, b를 a로
b가 a보다 큰 경우, a를 b로 변경해 준다고 가정하고
어느 한 쪽을 변경시키기 위해 증가시킨 가중치(a-b의 절댓값)를 sum에 누적해서 더해 준다
*/
func rec(x, y1, y2 int) int {
	// 기저 사례: 리프 노드 -> 부모 노드
	if x == n {
		sum += absSub(tree[x][y1], tree[x][y2])
		return max(tree[x][y1], tree[x][y2])
	}

	a := &dp[x][y1]
	b := &dp[x][y2]

	*a = rec(x+1, y1*2-1, y1*2) + tree[x][y1] // 부모 노드 (x, y1)에서 시작하는 가중치 합의 최댓값
	*b = rec(x+1, y2*2-1, y2*2) + tree[x][y2] // 부모 노드 (x, y2)에서 시작하는 가중치 합의 최댓값

	sum += absSub(*a, *b) // a-b의 절댓값만큼 가중치를 추가하고 sum에 누적
	return max(*a, *b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func absSub(a, b int) int {
	temp := a - b
	if temp < 0 {
		return -temp
	}
	return temp
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
