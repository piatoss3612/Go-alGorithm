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
	robot   [1000001]int // 로봇(부품) 번호
	part    [1000001]int // 부품의 개수
	N       int
)

// 주의! 부품 번호는 1부터 N까지가 아니라 10^6까지다
func init() {
	for i := 1; i <= 1000000; i++ {
		robot[i] = i
		part[i] = 1
	}
}

// 메모리: 22236KB
// 시간: 436ms
// 분리 집합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	N = scanInt()

	for i := 1; i <= N; i++ {
		ops := scanRune()

		switch ops {
		case 'I':
			a, b := scanInt(), scanInt()
			union(a, b) // a와 b를 합친다
		case 'Q':
			c := scanInt()
			fmt.Fprintln(writer, part[find(c)]) // c가 부품으로 사용되는 로봇의 부품 수를 출력
		}
	}
}

func find(x int) int {
	if robot[x] == x {
		return x
	}
	robot[x] = find(robot[x])
	return robot[x]
}

func union(x, y int) {
	x, y = find(x), find(y)
	if x != y {
		robot[y] = x
		part[x] += part[y]
		part[y] = part[x]
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanRune() rune {
	scanner.Scan()
	return rune(scanner.Bytes()[0])
}
