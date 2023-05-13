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

	N    int
	arr  [MAX + 1]int // arr[i] = 원래 수열에서 i보다 앞에 있는 i보다 큰 수의 개수
	tree [MAX*4 + 1]int
)

const MAX = 100000

// 난이도: Platinum 5
// 메모리: 5936KB
// 시간: 72ms
// 분류: 세그먼트 트리, 이분 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 1; i <= N; i++ {
		arr[i] = scanInt()
		update(i, 1, 1, N, 1) // 1~i까지 구간에 포함된 빈 칸의 개수를 세그먼트 트리에 저장
	}
}

func Solve() {
	ans := make([]int, N+1)

	for i := 1; i <= N; i++ {
		pos := query(arr[i]+1, 1, N, 1) // i가 들어갈 위치는 전체 구간의 빈 칸 중 arr[i]+1번째 빈 칸
		update(pos, -1, 1, N, 1) // pos번째 빈 칸을 차지했으므로 빈 칸의 개수를 1 줄임
		ans[pos] = i // pos번째 빈 칸에 i를 넣음
	}

	for i := 1; i <= N; i++ {
		fmt.Fprintln(writer, ans[i])
	}
}

func query(target, left, right, node int) int {
	if left == right {
		return left
	}

	mid := (left + right) / 2

	if tree[node*2] >= target {
		return query(target, left, mid, node*2)
	}
	return query(target-tree[node*2], mid+1, right, node*2+1)
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

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
