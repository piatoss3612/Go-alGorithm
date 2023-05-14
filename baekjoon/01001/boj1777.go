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
	arr  [MAX + 1]int // arr[i] = 원래 수열에서 i보다 뒤에 나오면서 i보다 작은 수의 개수
	tree [MAX*4 + 1]int
)

const MAX = 100000

// 난이도: Platinum 5
// 메모리: 5928KB
// 시간: 80ms
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
		update(i, 1, 1, N, 1)
	}
}

func Solve() {
	ans := make([]int, N+1)

	// 뒤에서부터 수열을 채워나간다
	for i := N; i >= 1; i-- {
		pos := query(i-arr[i], 1, N, 1) // i보다 뒤에 나오면서 i보다 작은 수의 개수가 arr[i]개일 때 i의 위치
		update(pos, -1, 1, N, 1) // pos번째 자리를 제외하고 남은 빈 자리의 개수 갱신
		ans[pos] = i // pos번째 자리에 i를 넣는다
	}

	for i := 1; i <= N; i++ {
		fmt.Fprintf(writer, "%d ", ans[i])
	}
	fmt.Fprintln(writer)
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
