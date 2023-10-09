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
	N       int
	cables  []Cable // 랜선 정보
	parent  []int   // 부모 노드
	total   int     // 랜선 길이 합
)

// 랜선 정보
type Cable struct {
	u, v, l int
}

// 난이도: Gold 3
// 메모리: 992KB
// 시간: 4ms
// 분류: 문자열, 최소 스패닝 트리
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	cables = make([]Cable, 0, N*N)
	parent = make([]int, N)

	// 랜선 정보 입력
	for i := 0; i < N; i++ {
		line := scanBytes()
		for j := 0; j < N; j++ {
			// 랜선이 연결되어 있는 경우, 랜선 길이를 계산하여 케이블 정보를 저장하고 랜선 길이의 누적 합을 구한다
			switch {
			case line[j] >= 'a': // a~z
				l := int(line[j] - 'a' + 1)
				total += l
				cables = append(cables, Cable{i, j, l})
			case line[j] >= 'A': // A~Z
				l := int(line[j] - 'A' + 27)
				total += l
				cables = append(cables, Cable{i, j, l})
			}
		}
		parent[i] = i
	}

	// 랜선 길이를 기준으로 오름차순 정렬
	sort.Slice(cables, func(i, j int) bool {
		return cables[i].l < cables[j].l
	})
}

func Solve() {
	cnt := 0 // 최소 스패닝 트리에 포함된 간선의 수

	// 랜선 길이가 짧은 것부터 탐색하여 두 컴퓨터가 연결되어 있지 않은 경우, 두 컴퓨터를 연결하고 랜선 길이를 누적 합에서 뺀다
	for len(cables) > 0 && cnt < N-1 {
		c := cables[0]
		cables = cables[1:]

		pu, pv := Find(c.u), Find(c.v)
		if pu != pv {
			parent[pv] = pu
			total -= c.l
			cnt += 1
		}
	}

	// 모든 컴퓨터가 연결되어 있지 않은 경우(최소 스패닝 트리를 구성하지 못한 경우), -1을 출력한다
	if cnt != N-1 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, total)
	}
}

func Find(x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = Find(parent[x])
	return parent[x]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
