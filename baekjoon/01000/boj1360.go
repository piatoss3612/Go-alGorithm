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
	N       int
)

// 1360번: 되돌리기
// hhttps://www.acmicpc.net/problem/1360
// 난이도: 골드 5
// 메모리: 876 KB
// 시간: 4 ms
// 분류: 구현
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
}

func Solve() {
	type Element struct {
		time int
		text string
	}

	arr := []Element{}

	arr = append(arr, Element{time: 0, text: ""})

	for i := 1; i <= N; i++ {
		cmd := scanString()
		text := scanString()
		time := scanInt()

		switch cmd {
		case "type":
			newText := arr[i-1].text + text                       // 이전 텍스트에 새로운 텍스트를 추가
			arr = append(arr, Element{time: time, text: newText}) // 새로운 텍스트를 time초에 생성
		case "undo":
			undoTime := mustParseInt(text) // time-undoTime초부터 time초까지의 텍스트를 제거해야 함
			flag := false
			for j := i - 1; j >= 0; j-- {
				// time-undoTime초 이전의 텍스트를 찾은 경우
				if arr[j].time < time-undoTime {
					arr = append(arr, Element{time: time, text: arr[j].text})
					flag = true
					break
				}
			}
			// 모든 텍스트를 제거해야 하는 경우
			if !flag {
				arr = append(arr, Element{time: time, text: ""})
			}
		}
	}

	fmt.Fprintln(writer, arr[N].text)
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
