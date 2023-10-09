package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner      = bufio.NewScanner(os.Stdin)
	writer       = bufio.NewWriter(os.Stdout)
	N, M         int
	drain        []int  // i번 도시의 배수로의 용량
	rainfall     []int  // i번 도시의 강수량
	parent       []int  // 각 도시와 연결된 루트 도시
	flooded      []bool // i번 도시가 홍수가 나는지 여부
	connected    []int  // 자신을 포함하여 연결된 도시의 개수
	totalFlooded int    // 홍수가 나는 모든 도시의 수
)

// 난이도: Gold 3
// 메모리: 12096KB
// 시간: 72ms
// 분류: 분리 집합

// 2번 쿼리를 실행할 때마다 N개의 도시를 반복문으로 돌아보며 홍수가 나는 도시의 개수를 구하면
// 최악의 경우 100000 * 100000 (3 <= N <= 100000, 1 <= M <= 100000) 이상의 시간이 걸려 시간 초과가 발생한다
// 따라서 2번 쿼리가 입력되면 바로 홍수가 나는 도시의 개수를 출력할 수 있도록
// 1번 쿼리인 유니온 연산을 실행할 때 미리 값을 계산해 놓는 방법을 고민했다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Query()
}

func Input() {
	N, M = scanInt(), scanInt()
	drain = make([]int, N+1)
	rainfall = make([]int, N+1)
	parent = make([]int, N+1)
	flooded = make([]bool, N+1)
	connected = make([]int, N+1)

	for i := 1; i <= N; i++ {
		drain[i] = scanInt() // i번 도시의 배수로의 용량 입력
		parent[i] = i        // i번 도시의 초기 루트 도시는 i
		connected[i] = 1     // i번 도시와 초기에 연결된 도시는 자기자신뿐
	}

	for i := 1; i <= N; i++ {
		rainfall[i] = scanInt() // i번 도시의 강수량 입력
		// i번 도시의 배수로의 용량보다 강수량이 큰 경우
		if rainfall[i] > drain[i] {
			flooded[i] = true // i번 도시에 홍수 발생 예상
			totalFlooded++    // 홍수가 발생할 도시의 개수 증가
		}
	}
}

func Query() {
	for i := 1; i <= M; i++ {
		q := scanInt()
		switch q {
		case 1:
			union(scanInt(), scanInt()) // 1번 쿼리: 유니온 연산 실행
		case 2:
			fmt.Fprintln(writer, totalFlooded) // 2번 쿼리: 홍수가 날 도시의 개수 출력
		}
	}
}

func find(x int) int {
	if x == parent[x] {
		return x
	}
	parent[x] = find(parent[x])
	return parent[x]
}

func union(x, y int) {
	x, y = find(x), find(y)
	if x != y {
		// x번 도시를 기준으로 값갱신
		parent[y] = x

		// 배수로 용량 갱신
		drain[y] += drain[x]
		drain[x] = drain[y]

		// 강수량 갱신
		rainfall[y] += rainfall[x]
		rainfall[x] = rainfall[y]

		if drain[x] >= rainfall[x] {
			// 배수로의 용량이 강수량보다 크거나 같아서 홍수를 예방할 수 있는 경우
			switch {
			// 1. x번 도시와 y번 도시가 모두 홍수가 날 것이라 예상된 경우
			case flooded[x] && flooded[y]:
				// 홍수는 예방할 수 있으므로 관련된 값 갱신
				// x번 도시가 홍수가 나지 않는다면 x번 도시와 연결된 다른 도시들도 홍수가 나지 않을 것이다
				// y번 도시가 홍수가 나지 않는다면 y번 도시와 연결된 다른 도시들도 홍수가 나지 않을 것이다
				totalFlooded -= (connected[x] + connected[y])
				flooded[x] = false
				flooded[y] = false
			// 2. x번 도시만 홍수가 날 것이라 예상된 경우
			case flooded[x]:
				totalFlooded -= connected[x]
				flooded[x] = false
			// 3. y번 도시만 홍수가 날 것이라 예상된 경우
			case flooded[y]:
				totalFlooded -= connected[y]
				flooded[y] = false
			}
		} else if drain[x] < rainfall[x] {
			// 배수로의 용량이 강수량보다 작아서 홍수가 날 것이라 예상되는 경우
			switch {
			// 1. x번 도시만 홍수로부터 안전했던 경우
			case !flooded[x]:
				// 홍수가 날 것이라 예상되었으므로 관련된 값 갱신
				totalFlooded += connected[x]
				flooded[x] = true
			// 2. y번 도시만 홍수로부터 안전했던 경우
			case !flooded[y]:
				totalFlooded += connected[y]
				flooded[y] = true
			}
		}

		// 연결된 도시의 개수 갱신
		connected[y] += connected[x]
		connected[x] = connected[y]
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
