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
	l, c    int
	apb     []string
	visited []bool
	ans     []string
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	l, c = scanInt(), scanInt()
	apb = make([]string, c+1)
	for i := 1; i <= c; i++ {
		apb[i] = scanString()
	}

	// 알파벳 순으로 정렬
	sort.Strings(apb)

	visited = make([]bool, c+1)

	getPwd("", 1, 0, 0, 0)

	for _, v := range ans {
		fmt.Fprintln(writer, v)
	}
}

// 백트래킹으로 풀은 문제
func getPwd(pwd string, current, vowel, consonant, length int) {
	if length == l {
		// 모음이 1개 이상, 자음이 2개 이상인 경우만 ans 슬라이스에 추가
		if vowel >= 1 && consonant >= 2 {
			ans = append(ans, pwd)
		}
		return
	} else if length < l {
		for i := current; i <= c; i++ {
			if visited[i] {
				continue
			}

			visited[i] = true

			if isVowel(apb[i]) {
				getPwd(pwd+apb[i], i, vowel+1, consonant, length+1)
			} else {
				getPwd(pwd+apb[i], i, vowel, consonant+1, length+1)
			}

			visited[i] = false
		}
	}
}

// 알파벳의 모음이 5개이므로 모음만 판별해주면 간단해진다
func isVowel(s string) bool {
	if s == "a" || s == "e" || s == "i" || s == "o" || s == "u" {
		return true
	}
	return false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}
