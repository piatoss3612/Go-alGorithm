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

	N       int
	squares []Square
)

type Square struct {
	h, w int
}

// 난이도: Gold 4
// 메모리: 195912KB
// 시간: 856ms
// 분류: 정렬, 스택
// 풀이: 1사분면에 채색된 사각형의 넓이를 구하여 4를 곱한 값을 결과로 출력한다.
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Input()
	Solve()
}

func Input() {
	N = scanInt()
	squares = make([]Square, N)

	for i := 0; i < N; i++ {
		h, w := scanInt(), scanInt()
		squares = append(squares, Square{h: h / 2, w: w / 2}) // 1사분면에 채색된 사각형의 높이와 너비를 저장
	}

	// x좌표가 작은 순(x좌표가 같다면 y좌표가 작은 순)으로 정렬
	sort.Slice(squares, func(i, j int) bool {
		if squares[i].w == squares[j].w {
			return squares[i].h < squares[j].h
		}
		return squares[i].w < squares[j].w
	})
}

func Solve() {
	stack := []Square{}

	// 동일한 x좌표를 가진 사각형들 중 가장 큰 높이를 가진 사각형만을 스택에 저장
	for len(squares) > 0 {
		square := squares[0]
		squares = squares[1:]

		if len(stack) == 0 {
			stack = append(stack, square)
			continue
		}

		n := len(stack)
		top := stack[n-1]

		if top.w < square.w {
			stack = append(stack, square)
			continue
		}

		if top.h < square.h {
			stack[n-1] = square
		}
	}

	ans := 0
	ans += stack[0].h * stack[0].w

	// 스택에 저장된 사각형들의 넓이를 구하여 더한다.
	for i := 1; i < len(stack); i++ {
		ans += stack[i].h * stack[i].w                      // i번째 사각형의 넓이
		ans -= stack[i-1].w * min(stack[i].h, stack[i-1].h) // i-1번째 사각형과 i번째 사각형의 높이 중 작은 값 * i번째 사각형의 너비를 빼준다.
	}

	fmt.Fprintln(writer, ans*4) // 1사분면에 채색된 사각형의 넓이를 구했으므로 4를 곱해준다.
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
