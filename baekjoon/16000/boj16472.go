package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	N        int
	str      []byte
	alphabet [26]int
)

// 난이도: Gold 4
// 메모리: 964KB
// 시간: 4ms
// 분류: 두 포인터
func main() {
	defer writer.Flush()
	scanner.Buffer(make([]byte, 0, 200000), 200000) // 입력받는 문자열의 길이가 최대 100000이므로 충분한 크기 버퍼 할당 필요
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	str = scanBytes()
}

func Solve() {
	l := 0   // 두 포인터: 왼쪽
	cnt := 0 // 부분 문자열에 포함된 알파벳의 종류
	ans := 0 // 최대 N개의 알파벳을 포함하는 부분 문자열의 최대 길이

	for r, v := range str {
		//
		// 오른쪽 포인터(r)의 위치에 있는 문자를 부분 문자열에 포함
		c := v - 'a'
		if alphabet[c] == 0 {
			cnt++
		}
		alphabet[c]++

		// 부분 문자열에 포함된 알파벳의 종류가 N개보다 많은 경우
		if cnt > N {
			// 알파벳의 종류가 N개 이하가 되도록 왼쪽 포인터를 오른쪽으로 이동
			for cnt > N {
				c = str[l] - 'a'
				alphabet[c]--
				if alphabet[c] == 0 {
					cnt--
				}
				l++
			}
		}
		ans = max(ans, r-l+1) // 최댓값 갱신
	}

	fmt.Fprintln(writer, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
