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
	root    *Element
)

type Element struct {
	Value int
	Left  *Element
	Right *Element
}

func NewElement(value int) *Element {
	return &Element{Value: value}
}

func (e *Element) Insert(value int) {
	if value < e.Value {
		if e.Left == nil {
			e.Left = NewElement(value)
		} else {
			e.Left.Insert(value)
		}
	} else {
		if e.Right == nil {
			e.Right = NewElement(value)
		} else {
			e.Right.Insert(value)
		}
	}
}

func (e *Element) PostOrder() {
	if e.Left != nil {
		e.Left.PostOrder()
	}
	if e.Right != nil {
		e.Right.PostOrder()
	}

	fmt.Fprintln(writer, e.Value)
}

// 5639번: 이진 검색 트리
// hhttps://www.acmicpc.net/problem/5639
// 난이도: 골드 5
// 메모리: 2852 KB
// 시간: 372 ms
// 분류: 트리, 재귀
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	initialized := false

	for scanner.Scan() {
		n := mustParseInt(scanner.Text())

		if !initialized {
			root = NewElement(n)
			initialized = true
		} else {
			root.Insert(n)
		}
	}
}

func Solve() {
	root.PostOrder()
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
