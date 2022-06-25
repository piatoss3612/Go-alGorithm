package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	army     []int // 병력
	allience []int // 동맹 관계
	n, m     int
)

// 메모리: 7624KB
// 시간: 60ms
// 전국시대에 전국칠웅이 있었고... 만화 킹덤을 보시면...
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	n, m = scanInt(), scanInt()
	army = make([]int, n+1)
	allience = make([]int, n+1)

	for i := 1; i <= n; i++ {
		allience[i] = i
		army[i] = scanInt()
	}

	var action, p, q int
	for i := 1; i <= m; i++ {
		action, p, q = scanInt(), scanInt(), scanInt()
		switch action {
		case 1:
			union(p, q) // p와 q가 동맹을 맺는다
		case 2:
			war(p, q) // p와 q가 전쟁을 한다
		}
	}

	res := []int{}
	for i := 1; i <= n; i++ {
		// i번째 국가가 속한 동맹의 수장이 자기 자인인 경우에만 남은 병력 수를 센다
		if allience[i] == i {
			res = append(res, army[i])
		}
	}

	sort.Ints(res) // 남은 병력 수 오름차 순으로 정렬
	fmt.Fprintln(writer, len(res))
	for _, v := range res {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer)
}

// x가 속한 동맹의 수장을 찾는다
func find(x int) int {
	if allience[x] == x {
		return x
	}
	allience[x] = find(allience[x])
	return allience[x]
}

// x와 y가 동맹을 맺는다
func union(x, y int) {
	x, y = find(x), find(y)
	if x != y {
		// 누가 수장인지는 중요하지 않으므로 x를 우선적으로 수장으로 본다
		allience[y] = x
		total := army[x] + army[y] // x와 y의 병력을 합친다
		army[x], army[y] = total, total
	}
}

// x와 y가 전쟁을 한다
func war(x, y int) {
	x, y = find(x), find(y)

	// x가 속한 동맹의 병력이 y가 속한 동맹의 병력보다 많은 경우
	if army[x] > army[y] {
		army[x] -= army[y]
		allience[y], army[y] = x, army[x] // y는 x의 속국이 된다
		return
	}

	// y가 속한 동맹의 병력이 x가 속한 동맹의 병력보다 많은 경우
	if army[x] < army[y] {
		army[y] -= army[x]
		allience[x], army[x] = y, army[y] // x는 y의 속국이 된다
		return
	}

	// x와 y의 병력 수가 같은 경우
	// x와 y 둘다 파격적인 멸망의 길을 걷게 된다
	allience[x], army[x] = 0, 0
	allience[y], army[y] = 0, 0
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
