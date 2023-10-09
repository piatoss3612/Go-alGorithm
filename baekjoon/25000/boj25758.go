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

	N     int
	genes [100001]string
	first [26]int
)

// 난이도: Silver 1
// 메모리: 2732KB
// 시간: 80ms
// 분류: 브루트포스 알고리즘, 해시를 사용한 집합과 맵
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()

	for i := 1; i <= N; i++ {
		gene := scanString()
		genes[i] = gene
		first[gene[0]-'A']++ // 앞자리가 int(gene[0]-'A')인 유전자의 개수 증가
	}
}

func Solve() {
	set := make(map[int]bool) // 2세대 유전자의 표현형을 저장할 집합

	for i := 0; i < 26; i++ {
		switch {
		// i번째 알파벳에 해당하는 유전자의 앞자리가 2개 이상이면 서로 다른 i번째 알파벳을 가진 유전자를 조합할 수 있음
		case first[i] > 1:
			for j := 1; j <= N; j++ {
				set[max(i, int(genes[j][1]-'A'))] = true
			}
		// i에 해당하는 유전자의 앞자리가 1개뿐이면 i번째 알파벳을 i번째 알파벳을 가진 유전자와 조합할 수 없음
		case first[i] == 1:
			for j := 1; j <= N; j++ {
				if int(genes[j][0]-'A') == i {
					continue
				}
				set[max(i, int(genes[j][1]-'A'))] = true
			}
		default:
			continue
		}
	}

	ans := []byte{}

	for i := 0; i < 26; i++ {
		if set[i] {
			ans = append(ans, byte(i)+'A')
		}
	}

	// 정답을 사전순으로 출력하기 위해 정렬
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})

	fmt.Fprintln(writer, len(ans))
	for _, v := range ans {
		fmt.Fprintf(writer, "%c ", v)
	}
	fmt.Fprintln(writer)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
