package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)

	N int

	arr  [][2]int
	tree [MAX * 4]int
)

const MAX = 500000

// 난이도: Platinum 5
// 메모리: 27976KB
// 시간: 448ms
// 분류: 자료 구조, 세그먼트 트리
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
}

func Setup() {
	for {
		N = scanInt()
		if N == 0 {
			return
		}

		arr = make([][2]int, N)
    	tree = [MAX*4]int{}

		for i := 0; i < N; i++ {
			arr[i][0] = scanInt()
			arr[i][1] = i + 1
		}

		sort.Slice(arr, func(i, j int) bool {
			return arr[i][0] < arr[j][0]
		})

		Solve()
	}
}

func Solve() {
	swap := 0

	for i := 0; i < N; i++ {
		pos := arr[i][1]
		swap += query(pos+1, N, 1, N, 1)
		update(pos, 1, 1, N, 1)
	}

	fmt.Fprintln(writer, swap)
}

func query(left, right, nodeLeft, nodeRight, node int) int {
	if right < nodeLeft || nodeRight < left {
		return 0
	}

	if left <= nodeLeft && nodeRight <= right {
		return tree[node]
	}

	mid := (nodeLeft + nodeRight) / 2
	return query(left, right, nodeLeft, mid, node*2) + query(left, right, mid+1, nodeRight, node*2+1)
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
