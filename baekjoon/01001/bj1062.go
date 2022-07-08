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
	N, K     int
	words    []int // 각 단어
	alphabet int   // 배운 알파벳들을 비트마스크로 표현
	ans      = 0
)

// 메모리: 912KB
// 시간: 44ms
// 비트마스크, 브루트포스 문제: 비트마스크로 풀려니까 더 어려운 것 같다...
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, K = scanInt(), scanInt()
	words = make([]int, N+1)
	for i := 1; i <= N; i++ {
		words[i] = scanWordToBitMask() // 각 단어를 비트마스크로 표현하여 저장
	}

	// 모든 단어는 "anta", "tica"를 포함하므로 못해도 5개의 알파벳은 배워야 한다
	// 그렇지 않은 경우는 어떠한 단어도 배울 수 없다
	if K < 5 {
		fmt.Fprintln(writer, 0)
		return
	}

	// 모든 알파벳을 배운 경우는 모든 단어를 배울 수 있다
	if K == 26 {
		fmt.Fprintln(writer, N)
		return
	}

	// "anta", "tica"에 포함된 알파벳 전처리
	alphabet |= 1 << ('a' - 'a')
	alphabet |= 1 << ('n' - 'a')
	alphabet |= 1 << ('t' - 'a')
	alphabet |= 1 << ('i' - 'a')
	alphabet |= 1 << ('c' - 'a')

	// 브루트포스: 전수 조사 시작
	rec(0, 0)

	fmt.Fprintln(writer, ans)
}

func rec(index, count int) {
	// 배운 알파벳의 수가 총 K-5개(a, c, i, n, t를 제외)일 때
	// 배울 수 있는 단어들을 확인하고 결과의 최댓값을 갱신
	if count == K-5 {
		temp := 0
		for i := 1; i <= N; i++ {
			// i번째 단어와 지금까지 배운 알파벳들을 and 연산을 했을 때
			// i번째 단어 자기 자신이 나온다면 배울 수 있는 단어
			if words[i]&alphabet == words[i] {
				temp++
			}
		}
		ans = max(ans, temp)
		return
	}

	// 백트래킹
	for i := index; i < 26; i++ {
		// i번째 알파벳을 아직 배우지 않았다면
		if alphabet&(1<<i) == 0 {
			alphabet |= 1 << i // or 연산으로 i번째 알파벳을 배운다
			rec(i+1, count+1)
			alphabet &= ^(1 << i) // and not 연산으로 i번째 알파벳을 제거한다
		}
	}
}

// 시간 초과 원인: 비트마스크 초짜라 이런 단순한 생각을 먼저 해버렸다
func letterCount(word int) int {
	cnt := 0
	for word > 0 {
		word &= word - 1
		cnt++
	}
	return cnt
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanWordToBitMask() int {
	scanner.Scan()
	word := scanner.Bytes()
	word = word[4:]           // "anta" 제거
	word = word[:len(word)-4] // "tica" 제거

	// 비트마스크로 변환하여 반환
	bitMask := 0
	for _, w := range word {
		bitMask |= 1 << (w - 'a')
	}
	return bitMask
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
