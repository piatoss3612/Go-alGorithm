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
	N, B, W int
	pebbles []byte
)

const LIMIT = 300000

// 난이도: Gold 4
// 메모리: 1224KB
// 시간: 12ms
// 분류: 두 포인터
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 0, LIMIT), LIMIT) // 250점 만점을 받으려면 입력을 받을 버퍼 사이즈를 충분히 늘려야 한다
	Input()
	Solve()
}

func Input() {
	N, B, W = scanInt(), scanInt(), scanInt()
	pebbles = scanBytes()
}

func Solve() {
	l, r := 0, 0
	bCount, wCount := 0, 0
	ans := 0

	for r < N {
		// 반복: 검은 조약돌의 개수가 B 이하이면서 r이 N보다 작은 경우
		for bCount <= B && r < N {
			if pebbles[r] == 'W' {
				wCount++
			} else {
				bCount++
			}

			// 하얀 조약돌의 개수가 W 이상이고 검은 조약돌의 개수가 B 이하인 경우
			if wCount >= W && bCount <= B {
				ans = max(ans, r-l+1) // 산책 구간의 최대 길이 갱신
			}
			r++ // 오른쪽 경곗값 이동
		}

		// 반복: 검은 조약돌의 개수가 B보다 큰 경우
		for bCount > B {
			if pebbles[l] == 'W' {
				wCount--
			} else {
				bCount--
			}
			l++ // 왼쪽 경곗값 이동
		}
	}

	fmt.Fprintln(writer, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
