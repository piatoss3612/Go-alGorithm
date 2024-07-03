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
	S       string
	n       int
	mapping [130]byte
)

// 27396번:
// hhttps://www.acmicpc.net/problem/27396
// 난이도: 골드 5
// 메모리: 13868 KB
// 시간: 1208 ms
// 분류: 문자열, 해시를 사용한 집합과 맵
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 0, 300000), 300000) // 버퍼 크기에 따라 부분 점수가 다르게 매겨짐?

	Setup()
	Solve()
}

func Setup() {
	S = scanString()
	n = scanInt()

	// 'A'~'z'에 자기자신에 해당하는 바이트값 할당
	for i := 65; i <= 122; i++ {
		mapping[i] = byte(i)
	}
}

func Solve() {
	for ; n > 0; n-- {
		query := scanInt()

		// 쿼리문 실행
		switch query {
		case 1:
			p, q := scanBytes()[0], scanBytes()[0] // p -> q로 변경해야 함

			// 이 때 문자열을 처음부터 끝까지 돌면서 변경하는 것은 최대 n * len(S), 100,000 * 300,000이 걸리므로 타임 아웃

			// 각 문자가 어떤 문자로 변경되었는지 메모해 놓으면 문자열을 출력할 때 이를 참고하기만 하면 된다
			for i := 65; i <= 122; i++ {
				if mapping[i] == p {
					mapping[i] = q
				}
			}
		case 2:
			// 변경 사항을 확인하여 문자열 재구성 및 출력
			for _, r := range S {
				fmt.Fprint(writer, string(mapping[r]))
			}
			fmt.Fprintln(writer)
		}
	}
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
