package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	N, M, Q  int
	tank     []status // 물탱크
	pipeline []int    // 물탱크 연결 정보
)

type status struct {
	clean, stagnant int // 특정 물탱크와 연결된 청정수 탱크 또는 고인물 탱크의 수
}

// 메모리: 5596KB
// 시간: 112ms
// 분리 집합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M, Q = scanInt(), scanInt(), scanInt()
	tank = make([]status, N+1)
	pipeline = make([]int, N+1)
	for i := 1; i <= N; i++ {
		pipeline[i] = i
		s := scanInt()
		if s == 1 {
			tank[i].clean++
		} else {
			tank[i].stagnant++
		}
	}

	for i := 1; i <= M; i++ {
		u, v := scanInt(), scanInt()
		union(u, v) // u, v 물탱크 연결
	}

	for i := 1; i <= Q; i++ {
		K := scanInt()
		K = find(K)

		// K번 탱크에 연결된 모든 청정수와 고인물 물탱크의 수를 비교
		if tank[K].clean > tank[K].stagnant {
			fmt.Fprintln(writer, 1)
		} else {
			fmt.Fprintln(writer, 0)
		}
	}
}

func find(x int) int {
	if pipeline[x] == x {
		return x
	}
	pipeline[x] = find(pipeline[x])
	return pipeline[x]
}

func union(x, y int) {
	x, y = find(x), find(y)

	if x != y {
		pipeline[y] = x
		tank[x].clean += tank[y].clean
		tank[y].clean = tank[x].clean
		tank[x].stagnant += tank[y].stagnant
		tank[y].stagnant = tank[x].stagnant
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
