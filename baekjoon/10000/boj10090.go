package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// 난이도: Platinum 5
// 메모리: 50356KB
// 시간: 732ms
// 분류: 자료 구조, 세그먼트 트리

/*
var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)

	N int

	arr  [][2]int
	tree [MAX * 4]int
)

const MAX = 1000000

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()

	arr = make([][2]int, N)

	for i := 0; i < N; i++ {
		arr[i][0] = scanInt()
		arr[i][1] = i + 1
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

}

func Solve() {
	inversions := 0

	for i := 0; i < N; i++ {
		pos := arr[i][1]
		inversions += query(pos+1, N, 1, N, 1)
		update(pos, 1, 1, N, 1)
	}

	fmt.Fprintln(writer, inversions)
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

func update(pos, val, left, right, node int) {
	if pos < left || right < pos {
		return
	}

	tree[node] += val

	if left != right {
		mid := (left + right) / 2
		update(pos, val, left, mid, node*2)
		update(pos, val, mid+1, right, node*2+1)
	}
}
*/

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)

	N int

	arr  [][2]int
	tree [MAX + 1]int
)

const MAX = 1000000

// 난이도: Platinum 5
// 메모리: 25904KB
// 시간: 192ms
// 분류: 자료 구조, 세그먼트 트리, 펜윅 트리
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

	N int

	arr  [MAX + 1]int
	tree [MAX + 1]int
)

const MAX = 1000000

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
	}
}

func Solve() {
	inversions := 0

	for i := 1; i <= N; i++ {
		inversions += sum(N) - sum(arr[i])
		update(arr[i], 1)
	}

	fmt.Fprintln(writer, inversions)
}

func sum(pos int) (ret int) {
	for pos > 0 {
		ret += tree[pos]
		pos &= pos - 1
	}
	return
}

func update(pos, val int) {
	for pos <= N {
		tree[pos] += val
		pos += pos & -pos
	}
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}


func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
