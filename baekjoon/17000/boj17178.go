package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N       int
	inp     []*ticket // 줄에서 기다리고 있는 사람들
	target  []*ticket // 주어진 티켓 번호에 따라 입장해야 하는 순서
)

type ticket struct {
	alpha  string
	number int
}

// 메모리: 996KB
// 시간: 4ms
// 스택
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanLines)
	N = scanInt()

	inp = make([]*ticket, 0, N*5)
	target = make([]*ticket, 0, N*5)
	for i := 0; i < N; i++ {
		line := scanLine()
		inp = append(inp, line...)
		target = append(target, line...)
	}

	// target 슬라이스를 알파벳, 숫자에 따라 오름차순으로 정렬
	sort.Slice(target, func(i, j int) bool {
		if target[i].alpha == target[j].alpha {
			return target[i].number < target[j].number
		}
		return target[i].alpha < target[j].alpha
	})

	stack := []*ticket{} // 스택, 콘서트 입장 대기열

	// 입장 순서 확인
	for i := 0; i < len(inp); i++ {
		// 스택의 가장 뒤에 있는 티켓이 입장해야 하는 순서의 가장 앞에 있다면
		// 스택에서 해당 티켓을 제거하고 콘서트장에 입장시킨다
		for len(stack) > 0 && stack[len(stack)-1] == target[0] {
			stack = stack[:len(stack)-1]
			target = target[1:]
		}

		// 그렇지 않은 경우 대기열로 들어간다
		stack = append(stack, inp[i])
	}

	// 대기열에 남은 사람들의 입장 순서 확읹
	for len(stack) > 0 {
		if stack[len(stack)-1] == target[0] {
			stack = stack[:len(stack)-1]
			target = target[1:]
		} else {
			// 입장 순서가 올바르지 않은 경우
			// "BAD"를 출력하고 프로그램 종료
			fmt.Fprintln(writer, "BAD")
			return
		}
	}

	// 모든 인원을 순서에 따라 콘서트장에 입장 시킴
	fmt.Fprintln(writer, "GOOD")
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanLine() []*ticket {
	scanner.Scan()
	fields := strings.Fields(scanner.Text())

	var t []*ticket

	for _, field := range fields {
		s := strings.Split(field, "-")
		n, _ := strconv.Atoi(s[1])
		t = append(t, &ticket{s[0], n})
	}

	return t
}
