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

	N, K int
	arr  []int
	tree []int
)

const MAX = 65537

// 난이도: Platinum 5
// 메모리: 5552KB
// 시간: 148ms
// 분류: 세그먼트 트리, 이분 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, K = scanInt(), scanInt()
	arr = make([]int, N+1)
	tree = make([]int, MAX*3)
	for i := 1; i <= N; i++ {
		arr[i] = scanInt()
	}
}

func Solve() {
	sum := 0
	cnt := (K + 1) / 2
	for i := 1; i <= N; i++ {
		update(arr[i], 1, 0, MAX, 1) // 세그먼트 트리에 i번째 수를 추가
		// i가 K보다 크거나 같은 경우
		if i >= K {
			// i-K+1부터 i까지의 수 중 K/2번째 수를 찾는다
			v := query(cnt, 0, MAX, 1)
			// i-K+1번째 수를 세그먼트 트리에서 제거
			update(arr[i-K+1], -1, 0, MAX, 1)
			sum += v // 구한 수를 더한다
		}
	}
	fmt.Fprintln(writer, sum)
}

func update(target, diff, left, right, node int) {
	if target < left || right < target {
		return
	}

	tree[node] += diff

	if left == right {
		return
	}

	mid := (left + right) / 2
	update(target, diff, left, mid, node*2)
	update(target, diff, mid+1, right, node*2+1)
}

func query(cnt, left, right, node int) int {
	if left == right {
		return left
	}

	mid := (left + right) / 2
	if tree[node*2] >= cnt {
		return query(cnt, left, mid, node*2)
	}
	return query(cnt-tree[node*2], mid+1, right, node*2+1)
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
