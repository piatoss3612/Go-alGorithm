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
	N       int
	x, y    float64
	result  [6]map[float64]int // 4분면 + x가 0인경우 + y가 0인 경우
	INF     float64            = 987654321
)

// 메모리: 12552KB
// 시간: 76ms
// 해시를 사용한 집합과 맵
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanLines)
	N = scanInt()

	for i := 0; i < 6; i++ {
		result[i] = make(map[float64]int)
	}

	for i := 1; i <= N; i++ {
		x, y = scanLine()

		// 0. 문제 조건상 x가 0이고 y가 0인 경우는 불가능

		// 1. x가 0인 경우
		if x == 0 {
			// x가 0인 경우 화살을 쏘는 방향을 고려해야 한다
			if y > 0 {
				result[0][1]++
			} else {
				result[0][0]++
			}
			continue
		}

		// 2. y가 0인 경우
		if y == 0 {
			// y가 0인 경우에도 화살을 쏘는 방향을 고려해야 한다
			if x > 0 {
				result[1][1]++
			} else {
				result[1][0]++
			}
			continue
		}

		// 3. 제1사분면 방향으로 화살을 쏘는 경우
		if x > 0 && y > 0 {
			// 화살을 쏘는 위치는 항상 (0,0)이다
			// 따라서 그래프의 기울기는 y/x가 되며
			// 기울기는 곧 화살을 쏘는 방향이라고 할 수 있다
			// 따라서 그래프의 기울기 y/x 값에 매핑된 정수값을 증가시킨다
			result[2][y/x]++
			continue
		}

		// 3. 제2사분면 방향으로 화살을 쏘는 경우
		if x < 0 && y > 0 {
			result[3][y/x]++
			continue
		}

		// 3. 제3사분면 방향으로 화살을 쏘는 경우
		if x < 0 && y < 0 {
			result[4][y/x]++
			continue
		}

		// 3. 제4사분면 방향으로 화살을 쏘는 경우
		if x > 0 && y < 0 {
			result[5][y/x]++
		}
	}

	max := 0

	// 화살 한 개를 쏘아 얻을 수 있는 가장 높은 점수 갱신
	for i := 0; i < 6; i++ {
		for _, v := range result[i] {
			if v > max {
				max = v
			}
		}
	}

	fmt.Fprintln(writer, max)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanLine() (float64, float64) {
	scanner.Scan()
	fields := strings.Fields(scanner.Text())

	x, _ := strconv.ParseFloat(fields[0], 64)
	y, _ := strconv.ParseFloat(fields[1], 64)
	return x, y
}
