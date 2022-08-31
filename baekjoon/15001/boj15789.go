package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	N, M     int
	alliance []int // 동맹 관계
	power    []int // 국력
	X, Y     int
	C, H, K  int
)

// 메모리: 5480KB
// 시간: 84ms
// 분리 집합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	// 1. 데이터 초기화
	N, M = scanInt(), scanInt()
	alliance = make([]int, N+1)
	power = make([]int, N+1)
	for i := 1; i <= N; i++ {
		alliance[i] = i // 동맹은 자기자신이 맹주인 상태에서 시작
		power[i] = 1    // 동맹이 없는 상태의 왕국의 힘은 1
	}

	// 2. 동맹 관계 설정
	for i := 1; i <= M; i++ {
		X, Y = scanInt(), scanInt()
		union(X, Y) // X와 Y의 동맹
	}

	// 3. 동맹 후보 찾기
	C, H, K = scanInt(), scanInt(), scanInt()

	CTPBelongsTo := find(C)    // CTP 왕국이 속해있는 동맹
	HansolBelongsTo := find(H) // 한솔 왕국이 속해있는 동맹

	candidates := []int{} // 동맹 후보들의 국력

	for i := 1; i <= N; i++ {
		tmp := find(i) // i번째 왕국이 속한 동맹의 맹주

		// i가 속한 동맹의 맹주가 자기자신이고
		// 해당 동맹에 CTP 왕국이 포함되어 있지 않으면서
		// 한솔 왕국도 포함되어 있지 않다면
		// 동맹 후보로 고려할 수 있다
		if tmp == i && tmp != CTPBelongsTo && tmp != HansolBelongsTo {
			candidates = append(candidates, power[tmp])
		}
	}

	// 동맹 후보들의 국력을 가장 큰 것부터 내림차순으로 정렬
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] > candidates[j]
	})

	// 4. CTP 왕국의 힘의 최댓값 구하기
	ans := power[CTPBelongsTo]

	// K만큼의 동맹 기회와 남은 동맹 후보들을 고려하여
	// 가장 국력이 강한 동맹 후보부터 선택
	for K > 0 && len(candidates) > 0 {
		ans += candidates[0]
		candidates = candidates[1:]
		K--
	}

	fmt.Fprintln(writer, ans)
}

func find(x int) int {
	if alliance[x] == x {
		return x
	}

	alliance[x] = find(alliance[x])
	return alliance[x]
}

func union(x, y int) {
	x, y = find(x), find(y)
	if x != y {
		alliance[y] = x
		power[x] += power[y]
		power[y] = power[x]
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
