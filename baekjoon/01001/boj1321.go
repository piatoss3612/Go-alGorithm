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

	N, M     int
	camps    [500001]int
	segments [1200000]int
)

// 난이도: Platinum 5
// 메모리: 14732KB
// 시간: 88ms
// 분류: 세그먼트 트리, 이분 탐색
// 풀이: 세그먼트 트리로 구간 합을 구하고 이분 탐색으로 1~m(m = (l+r)/2)번째 구간의 합이 x 이상이 되는 m의 최솟값을 구한다.
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 1; i <= N; i++ {
		camps[i] = scanInt()
	}

	NewSegmentTree(1, N, 1)

	M = scanInt()
}

func Solve() {
	for i := 1; i <= M; i++ {
		q := scanInt()

		switch q {
		case 1:
			x, y := scanInt(), scanInt()
			Update(x, y, 1, N, 1)
			camps[x] += y
		case 2:
			x := scanInt()
			l, r := 1, N
			for l <= r {
				m := (l + r) / 2
				y := Query(1, m, 1, 1, N)

				if y >= x {
					r = m - 1
				} else {
					l = m + 1
				}
			}

			fmt.Fprintln(writer, l)
		}
	}
}

func NewSegmentTree(left, right, node int) int {
	if left == right {
		segments[node] = camps[left]
		return segments[node]
	}

	mid := (left + right) / 2
	segments[node] = NewSegmentTree(left, mid, node*2) + NewSegmentTree(mid+1, right, node*2+1)

	return segments[node]
}

func Update(target, amount, left, right, node int) {
	if target < left || right < target {
		return
	}

	if left == right {
		segments[node] += amount
		return
	}

	segments[node] += amount

	mid := (left + right) / 2
	Update(target, amount, left, mid, node*2)
	Update(target, amount, mid+1, right, node*2+1)
}

func Query(left, right, node, nodeLeft, nodeRight int) int {
	if right < nodeLeft || nodeRight < left {
		return 0
	}

	if left <= nodeLeft && nodeRight <= right {
		return segments[node]
	}

	mid := (nodeLeft + nodeRight) / 2
	return Query(left, right, node*2, nodeLeft, mid) + Query(left, right, node*2+1, mid+1, nodeRight)
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
