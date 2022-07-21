package main

import (
	"bufio"
	"fmt"
	_ "math"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	sentence = (1 << 26) - 1 // 테스트 문장과 비교할 모든 알파벳이 포함된 문장의 비트마스크
	words    [26]int         // 각 단어를 비트마스크로 변환하여 저장
	N        int
	ans      = 0
)

// 메모리: 900KB
// 시간: 196ms
// 비트마스크, 브루트포스
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()

	for i := 1; i <= N; i++ {
		words[i] = scanWordToBitmask()
	}

	// 브루트포스
	for i := 1; i <= N; i++ {
		rec(i, words[i])
	}

	fmt.Fprintln(writer, ans)
}

func rec(idx, test int) {
	// 모든 알파벳이 포함된 sentence와 test 문장을 and 연산한 결과가
	// sentence인 경우는 test 문장에 모든 알파벳이 포함되어 있다는 의미
	if (sentence & test) == sentence {
		ans++
	}

	for i := idx + 1; i <= N; i++ {
		rec(i, test|words[i]) // or 연산으로 test 문장에 i번째 단어에 포함된 알파벳 추가
	}
}

// 단어를 비트마스크로 변환
func scanWordToBitmask() int {
	scanner.Scan()
	word := scanner.Bytes()
	bitmask := 0
	// 바이트값으로 읽어온 알파벳 소문자에 -'a'를 할 경우
	// a~z는 0~25 정숫값이 된다
	// 이 정숫값만큼 1을 왼쪽으로 shift 연산하면
	// 알파벳을 비트마스크로 표현한 값이 된다
	// 이렇게 구한 값을 누적해서 or 연산 해줌으로써
	// 하나의 단어를 비트마스크로 표현한다
	for _, c := range word {
		bitmask |= 1 << (c - 'a')
	}
	return bitmask
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
