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
	memory  = (1 << 26) - 1 // 모든 알파벳을 외우고 있는 상태
	words   []int
	N, M    int
)

// 메모리: 1372KB
// 시간: 800ms
// 비트마스크, 브루트포스 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()

	words = make([]int, N+1)
	for i := 1; i <= N; i++ {
		words[i] = scanWordToBitMask() // 각 단어를 비트마스크로 표현
	}

	for i := 1; i <= M; i++ {
		o, x := scanInt(), scanChar()

		switch o {
		// 알파벳 x를 잊는 경우
		case 1:
			memory &^= x // and not 연산으로  x 제거
		// 알파벳 x를 기억하는 경우
		case 2:
			memory |= x // or 연산으로 x 기억
		}

		query() // 알고 있는 단어의 개수 출력
	}
}

// 브루트 포스
func query() {
	cnt := 0
	for i := 1; i <= N; i++ {
		// 기억과 i번째 단어를 and 연산한 결과가 i번째 단어인 경우
		if memory&words[i] == words[i] {
			cnt++ // i번째 단어는 완전히 기억하는 단어
		}
	}
	fmt.Fprintln(writer, cnt)
}

func scanWordToBitMask() int {
	scanner.Scan()
	word := scanner.Bytes()
	bm := 0
	for _, c := range word {
		bm |= 1 << (c - 'a')
	}
	return bm
}

func scanChar() int {
	scanner.Scan()
	return 1 << int(scanner.Bytes()[0]-'a')
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
