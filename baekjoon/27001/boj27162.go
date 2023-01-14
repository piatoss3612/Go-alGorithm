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
	options string
	a, b, c int
)

// 난이도: Silver 4
// 메모리: 912KB
// 시간: 4ms
// 분류: 구현, 브루트포스 알고리즘, 많은 조건 분기
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	options = scanString()
	a, b, c = scanInt(), scanInt(), scanInt()
}

func Solve() {
	ans := 0

	for i, v := range options {
		if v == 'Y' {
			ans = max(ans, calculate(i+1))
		}
	}

	fmt.Fprintln(writer, ans)
}

func calculate(option int) int {
	switch option {
	case 1:
		return diceNs(1)
	case 2:
		return diceNs(2)
	case 3:
		return diceNs(3)
	case 4:
		return diceNs(4)
	case 5:
		return diceNs(5)
	case 6:
		return diceNs(6)
	case 7:
		return fourOfAKind()
	case 8:
		return fullHouse()
	case 9:
		return littleStraight()
	case 10:
		return bigStraight()
	case 11:
		return yacht()
	case 12:
		return a + b + c + 12
	}
	return 0
}

func diceNs(n int) int {
	var sum int
	if a == n {
		sum += n
	}
	if b == n {
		sum += n
	}
	if c == n {
		sum += n
	}
	return sum + n*2
}

func fourOfAKind() int {
	if a == b && b == c {
		return a * 4
	}

	if a == b {
		return a * 4
	} else if b == c {
		return b * 4
	} else if c == a {
		return c * 4
	}
	return 0
}

func fullHouse() int {
	if a == b && b == c {
		m := 1
		for i := 2; i <= 6; i++ {
			if i != a {
				m = i
			}
		}
		return a*3 + m*2
	}
	if a == b || b == c {
		return max(a, c)*3 + min(a, c)*2
	}

	if a == c {
		return max(b, c)*3 + min(b, c)*2
	}
	return 0
}

func littleStraight() int {
	if a <= 5 && b <= 5 && c <= 5 && a != b && b != c && c != a {
		return 30
	}
	return 0
}

func bigStraight() int {
	if a >= 2 && b >= 2 && c >= 2 && a != b && b != c && c != a {
		return 30
	}
	return 0
}

func yacht() int {
	if a == b && b == c && c == a {
		return 50
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}
