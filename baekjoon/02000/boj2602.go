package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	scanner      = bufio.NewScanner(os.Stdin)
	writer       = bufio.NewWriter(os.Stdout)
	scroll       []byte          // 마법의 두루마리
	devil, angel []byte          // 악마 다리, 천사 다리
	dp           [2][20][100]int // dp[i][j][k]: i는 다리의 종류, j는 마법 두루마리에 적힌 문자의 위치, k는 다리에서 현재 위치
)

// 난이도: Gold 4
// 메모리: 960KB
// 시간: 112ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	Input()
	Solve()
}

func Input() {
	scanner.Split(bufio.ScanWords)
	scroll = scanBytes()
	devil = scanBytes()
	angel = scanBytes()
}

func Solve() {
	ans := 0

	// 악마 다리에서 시작
	for i, d := range devil {
		if d == scroll[0] {
			ans += rec(0, 0, i)
		}
	}
	// 천사 다리에서 시작
	for i, a := range angel {
		if a == scroll[0] {
			ans += rec(1, 0, i)
		}
	}

	fmt.Fprintln(writer, ans)
}

func rec(bridge, scrollPos, bridgePos int) int {
	// 기저 사례: 마법 두루마리에 적힌 문자열을 생성할 수 있는 경우
	if scrollPos == len(scroll)-1 {
		return 1
	}

	ret := &dp[bridge][scrollPos][bridgePos]
	if *ret != 0 {
		return *ret
	}

	next := scroll[scrollPos+1] // 다음으로 찾을 문자

	// 악마-천사 다리를 번갈아가면서 이동
	switch bridge {
	case 0:
		for i := bridgePos + 1; i < len(angel); i++ {
			if angel[i] == next {
				*ret += rec(1, scrollPos+1, i)
			}
		}
	case 1:
		for i := bridgePos + 1; i < len(devil); i++ {
			if devil[i] == next {
				*ret += rec(0, scrollPos+1, i)
			}
		}
	}

	return *ret
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
