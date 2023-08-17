package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)

	S     string
	sec   int
	check [3601][2]int // [초][전자레인지 상태] = 버튼 누른 횟수
)

// 난이도: Silver 1
// 메모리: 1080KB
// 시간: 4ms
// 분류: 그래프 이론, 그래프 탐색, 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	S = scanString()
	fields := strings.Split(S, ":")
	m := mustParseInt(fields[0])
	s := mustParseInt(fields[1])
	sec = m*60 + s
}

func Solve() {
	q := [][2]int{{0, 0}}

	for len(q) > 0 {
		front := q[0]
		q = q[1:]

		curr, state, cnt := front[0], front[1], check[front[0]][front[1]] // 현재 시간, 전자레인지 상태, 버튼 누른 횟수

		// 조리시간을 sec로 맞추고 조리중인 상태라면 버튼 누른 횟수 출력하고 종료
		if curr == sec && state == 1 {
			fmt.Fprintln(writer, cnt)
			return
		}

		// 10초, 1분, 10분 버튼 누르기
		if curr+10 <= 3600 && check[curr+10][state] == 0 {
			check[curr+10][state] = cnt + 1
			q = append(q, [2]int{curr + 10, state})
		}

		if curr+60 <= 3600 && check[curr+60][state] == 0 {
			check[curr+60][state] = cnt + 1
			q = append(q, [2]int{curr + 60, state})
		}

		if curr+600 <= 3600 && check[curr+600][state] == 0 {
			check[curr+600][state] = cnt + 1
			q = append(q, [2]int{curr + 600, state})
		}

		// 조리중이 아니라면 시작 버튼 누르기
		if state == 0 {
			if check[curr][1] == 0 {
				// 조리중이 아니고 조리시간이 0초라면 30초가 추가된다.
				if curr == 0 {
					if curr+30 <= 3600 && check[curr+30][1] == 0 {
						check[curr+30][1] = cnt + 1
						q = append(q, [2]int{curr + 30, 1})
					}
				} else {
					// 조리중이 아니고 조리시간이 0초가 아니라면 그대로 조리가 시작된다.
					check[curr][1] = cnt + 1
					q = append(q, [2]int{curr, 1})
				}
			}
		} else {
			// 조리중이라면 30초가 추가된다.
			if curr+30 <= 3600 && check[curr+30][state] == 0 {
				check[curr+30][state] = cnt + 1
				q = append(q, [2]int{curr + 30, state})
			}
		}
	}
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}