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
	N, T, W, M   int
	customers    []*Customer       // 고객 정보를 저장한 큐
	newCustomers map[int]*Customer // key: 영업이 시작되고 지난 x초, value: 새로 입장하는 고객 정보
)

// 고객 정보
type Customer struct {
	id, time int
}

// 메모리: 37916KB
// 시간: 268ms
// 큐, 시뮬레이션
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, T, W = scanInt(), scanInt(), scanInt()

	// 1. 줄에 서있는 고객 정보
	customers = make([]*Customer, N)
	for i := 0; i < N; i++ {
		customers[i] = &Customer{scanInt(), scanInt()}
	}

	M = scanInt()

	// 2. 영업 중에 새로 입장하는 고객 정보
	newCustomers = make(map[int]*Customer)
	for i := 0; i < M; i++ {
		p, t, c := scanInt(), scanInt(), scanInt()
		newCustomers[c] = &Customer{p, t}
	}

	seconds := 0 // 영업을 시작하고 흐른 시간 (초 단위)

	// 3. 영업 시작
	for seconds < W && len(customers) > 0 {
		// 3-1. 줄의 가장 앞에 있는 고객의 업무부터 처리
		front := customers[0]
		customers = customers[1:]

		temp := T // 고객을 응대하는 최대 시간

		// 3-2. 고객 응대
		for temp > 0 && seconds < W {
			fmt.Fprintln(writer, front.id) // 응대하는 고객의 id 출력
			temp--                         // 고객 응대 시간 1초 감소
			front.time--                   // 고객의 업무를 처리하는데 필요한 시간 1초 감소
			seconds++                      // 영업을 시작하고 흐른 시간 1초 증가

			// 3-2-1. 영업이 시작하고 seconds만큼 시간이 흘렀을 때 새롭게 입장하는 고객이 있는지 확인
			_, ok := newCustomers[seconds]
			if ok {
				customers = append(customers, newCustomers[seconds])
				delete(newCustomers, seconds)
			}

			// 3-2-2. 고객의 업무 처리를 완료한 경우
			if front.time == 0 {
				break
			}
		}

		// 3-3. 고객을 응대하는 시간을 모두 사용하고도 고객의 업무 처리를 완료하지 못한 경우
		if front.time > 0 {
			customers = append(customers, front)
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
