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
	pos  [MAX + 1]int // A열 기계의 위치
	tree [MAX * 4]int // 세그먼트 트리
)

const MAX = 1000000

// 난이도: Platinum 5
// 메모리: 26156KB
// 시간: 468ms
// 분류: 자료 구조, 세그먼트 트리
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	// A열 기계의 위치를 저장
	for i := 1; i <= N; i++ {
		x := scanInt()
		pos[x] = i
	}
}

func Solve() {
	total := 0

	for i := 1; i <= N; i++ {
		// B열 기계의 식별번호
		x := scanInt()
		// A열에서 식별번호가 x인 기계의 위치를 찾아서 세그먼트 트리에서 그 위치보다 오른쪽에 있는 기계의 수를 구한다
		// B열의 i번째 기계와 연결된 A열의 기계보다 오른쪽에 있다는 것은
		// 1~i-1번째 기계들 중에서 i번째 기계의 연결선과 교차하는 선이 존재함을 의미한다
		total += query(pos[x], MAX, 0, MAX, 1)
		// 식별번호가 x인 기계들을 연결한다 (세그먼트 트리에서 그 위치에 1을 더한다)
		update(pos[x], 0, MAX, 1)
	}

	fmt.Fprintln(writer, total)
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

func update(target, left, right, node int) {
	if target < left || right < target {
		return
	}

	tree[node] += 1

	if left == right {
		return
	}

	mid := (left + right) / 2
	update(target, left, mid, node*2)
	update(target, mid+1, right, node*2+1)
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
