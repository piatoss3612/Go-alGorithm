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
	n, w, L int
	truck   []int
)

// 난이도: Silver 1
// 메모리: 1556KB
// 시간: 8ms
// 분류: 구현, 자료 구조, 큐
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	n, w, L = scanInt(), scanInt(), scanInt()
	truck = make([]int, n)
	for i := 0; i < n; i++ {
		truck[i] = scanInt()
	}
}

func Solve() {
	queue := make([]int, w) // 다리의 상태
	onBridge := 0           // 다리 위의 트럭의 무게
	time := 0               // 모든 트럭이 다리를 지나는데 걸린 시간

	for len(truck) > 0 {
		curr := truck[0] // 다리를 건널 차례가 온 트럭

		// 다리 위의 트럭들이 한 칸씩 이동했다고 가정
		// 다리의 끝에 트럭이 있었다면 다리를 빠져나갔을 것
		onBridge -= queue[0]
		queue = queue[1:]

		if onBridge+curr > L {
			// curr이 아직 다리를 건널 수 없는 경우
			queue = append(queue, 0)
		} else {
			// curr이 다리를 건너기 시작한 경우
			queue = append(queue, curr)
			onBridge += curr
			truck = truck[1:]
		}

		time++
	}

	time += len(queue) // 마지막으로 다리를 건너기 시작한 트럭이 다리를 건너는데 필요한 시간 추가

	fmt.Fprintln(writer, time)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
