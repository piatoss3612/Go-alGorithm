package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	N, M       int
	founder    string             // 유토피아를 세운 사람
	rels       map[string]*Parent // 자식-부모 관계
	pureBlood  map[string]float64 // 혈통값
	candidates []string           // 왕위 계승 후보들
)

type Parent struct {
	p1, p2 string
}

// 메모리: 932KB
// 시간: 4ms
// 해시 맵
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()

	//
	rels = make(map[string]*Parent)
	pureBlood = make(map[string]float64)
	candidates = make([]string, M)

	// 왕국 설립자의 혈통값은 1
	founder = scanName()
	pureBlood[founder] = 1

	for i := 0; i < N; i++ {
		child, p1, p2 := scanName(), scanName(), scanName()
		rels[child] = &Parent{p1, p2} // 자식-부모 관계 저장
	}

	var nextKing string
	var max float64 = -987654321

	for i := 0; i < M; i++ {
		candidates[i] = scanName()
		pb := find(candidates[i]) // i번째 후보의 혈통값 구하기
		// i번째 후보의 혈통값이 더 큰 경우
		if pb > max {
			nextKing = candidates[i] // 다음 왕은 i번째 후보가 유력
			max = pb
		}
	}

	fmt.Fprintln(writer, nextKing)
}

func find(name string) float64 {
	// 이미 혈통값이 계산된 경우
	if b, ok := pureBlood[name]; ok {
		return b
	}

	// 부모가 없는 경우
	if _, ok := rels[name]; !ok {
		return 0
	}

	parents := rels[name]
	// 현재 후보의 혈통값은 부모의 혈통값을 더한 뒤 절반으로 나눈 값
	pureBlood[name] = (find(parents.p1) + find(parents.p2)) / 2
	return pureBlood[name]
}

func scanName() string {
	scanner.Scan()
	return scanner.Text()
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
