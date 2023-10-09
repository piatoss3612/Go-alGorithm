package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner         = bufio.NewScanner(os.Stdin)
	writer          = bufio.NewWriter(os.Stdout)
	board           [101][101]int
	directionChange map[int]rune
	N, K, L         int
)

type Coord struct {
	r, c int
}

// 메모리: 1312KB
// 시간: 4ms
// 덱을 사용한 게임 구현
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, K = scanInt(), scanInt()
	for i := 1; i <= K; i++ {
		board[scanInt()][scanInt()] = 1 // 사과가 있는 위치는 1
	}

	directionChange = make(map[int]rune) // 방향 변경 정보를 저장하는 맵
	L = scanInt()
	for i := 1; i <= L; i++ {
		directionChange[scanInt()] = scanRune()
	}

	board[1][1] = 2 // 뱀의 몸이 있는 위치는 2
	direction := 1  // 1: R, 2: D, 3: L, 4: U
	time := 0
	coords := []Coord{{1, 1}} // 뱀의 몸이 있는 좌표 정보를 저장하는 덱

	for {
		// 0. 시간 추가
		time++

		// 1. 이동할 좌표 구하기
		next := coords[len(coords)-1]

		switch direction {
		case 1:
			next.c++
		case 2:
			next.r++
		case 3:
			next.c--
		case 4:
			next.r--
		}

		// 2. 이동할 좌표에 벽이 있거나 자기자신의 몸이 있는 경우
		if !valid(next) || board[next.r][next.c] == 2 {
			fmt.Fprintln(writer, time)
			return
		}

		// 3. 이동할 좌표에 사과가 있는지
		if board[next.r][next.c] != 1 {
			board[coords[0].r][coords[0].c] = 0
			coords = coords[1:]
		}

		board[next.r][next.c] = 2
		coords = append(coords, next)

		// 4. 방향 변경 여부 확인
		change, ok := directionChange[time]

		if ok {
			switch change {
			case 'L':
				if direction == 1 {
					direction = 4
				} else {
					direction--
				}
			case 'D':
				if direction == 4 {
					direction = 1
				} else {
					direction++
				}
			}
		}
	}
}

func valid(coord Coord) bool {
	if coord.r >= 1 && coord.r <= N && coord.c >= 1 && coord.c <= N {
		return true
	}
	return false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanRune() rune {
	scanner.Scan()
	b := scanner.Bytes()
	return rune(b[0])
}
