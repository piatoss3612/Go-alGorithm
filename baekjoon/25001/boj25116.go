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
	stage   [100001]int
)

// 난이도: Gold 4
// 메모리: 3288KB
// 시간: 40ms
// 분류: 이분 탐색, 두 포인터, 매개 변수 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, M, K = scanInt(), scanInt(), scanInt()
	for i := 0; i < N; i++ {
		stage[i] = scanInt()
	}
}

func Solve() {
	// 연속되는 구간의 합이 M 이상이되는 구간(interesting section)의 개수가 K개 이상이 되도록
	// stage의 모든 원소에 동일한 값 X를 더해주어야 할 때
	// X의 최솟값을 구하는 문제

	// 이분 탐색을 사용하여 매개 변수 X를 구하는 과정에서
	// 구간의 개수 K는 두 포인터를 사용해 구했다

	l, r := 0, 1000000000000000000
	for l <= r {
		m := (l + r) / 2     // X
		sl, sr := 0, 0       // 두 포인터
		cnt := 0             // interesting section의 개수
		sum := stage[sr] + m // sl~sr 구간 합
		for sr < N {
			// 1) 구간 합이 M보다 작은 경우
			if sum < M {
				// 오른쪽 포인터 이동
				sr++
				sum += (stage[sr] + m)
				continue
			}

			// 2) 구간 합이 M보다 크거나 같은 경우
			// sl~sr 구간의 합이 M보다 크거나 같으므로 sr을 오른쪽으로 이동해도 구간 합은 항상 M보다 크거나 같다
			// 따라서 일일히 이동해 볼 필요없이 (N - sr)개의 연속되는 구간의 합이 M보다 크거나 같다는 것을 알 수 있다
			cnt += (N - sr)

			// 왼쪽 포인터 이동
			sum -= (stage[sl] + m)
			sl += 1

			// 구간의 길이가 1인 상태에서 왼쪽 포인터를 이동한 경우
			if sl > sr {
				// 오른쪽 포인터를 왼쪽 포인터와 동일한 지점으로 이동
				sr = sl
				sum += (stage[sr] + m)
			}

			// 구간의 합이 M보다 크거나 같은 연속된 구간의 개수가 K보다 크거나 같은 경우
			if cnt >= K {
				break
			}
		}

		// 구간의 합이 M보다 크거나 같은 연속된 구간의 개수가 K보다 크거나 같은 경우
		if cnt >= K {
			r = m - 1 // X의 크기를 줄인다
		} else {
			l = m + 1 // X의 크기를 늘린다
		}
	}

	fmt.Fprintln(writer, l) // 조건을 만족하는 최솟값을 구하는 것이므로 X의 lower bound인 l을 출력
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
