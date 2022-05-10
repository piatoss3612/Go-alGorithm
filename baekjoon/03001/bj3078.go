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
)

// 메모리: 6348KB
// 시간: 72ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, k := scanInt(), scanInt()
	var counts [21][]int // 2~20글자의 이름 길이에 해당하는 학생의 성적 순위를 저장하는 큐

	ans := 0

	/*
		앞에서부터 등수의 차이가 k 이내인 학생을 찾으면 n*k 시간이 걸리므로 시간 초과 발생!
		큐를 사용해 앞에서부터 찾지 말고, 새로운 값을 입력받을 때마다 뒤에서부터!

		예제 입력:
		4 2
		IVA
		IVO
		ANA
		TOM

		1. IVA 입력
		counts[3] = [1]
		ans = 0

		2. IVO 입력
		counts[3] = [1, 2]
		2등 학생과 좋은 친구가 될 수 있는 학생: 1등
		ans = 1

		3. ANA 입력
		counts[3] = [1, 2, 3]
		3등 학생과 좋은 친구가 될 수 있는 학생: 1등, 2등
		ans = 3

		4. TOM 입력
		counts[3] = [2, 3, 4]
		1등 학생은 4등 학생과 순위 차이가 2이상이므로 큐에서 Pop
		4등 학생과 좋은 친구가 될 수 있는 학생: 2등, 3등
		ans = 5

		이런 식으로 코드를 실행하면 n*k만큼 연산을 수행할 필요 없이 더 빠른 시간 내에 연산을 마칠 수 있다
	*/

	var x int

	// 현재 학생의 성적 순위는 입력되는 순서 i와 같다
	for i := 1; i <= n; i++ {
		x = scanLen() // 이름의 길이

		// 이름 길이 x에 해당하는 큐의 길이가 0이 아니면서
		// (i: 현재 학생의 순위 - 가장 앞에 있는 학생의 순위)가 k보다 큰 경우
		// 즉, 가장 앞에 있는 학생이 i번째 순위의 학생과 좋은 친구가 될 수 없는 경우
		for len(counts[x]) > 0 && i-counts[x][0] > k {
			counts[x] = counts[x][1:] // 가장 앞에 있는 학생을 큐에서 Pop
		}
		ans += len(counts[x])            // i번째 순위의 학생과 좋은 친구가 될 수 있는 학생의 수는 len(counts[x])
		counts[x] = append(counts[x], i) // 마지막에 i번째 학생을 큐에 추가
	}
	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanLen() int {
	scanner.Scan()
	return len(scanner.Bytes())
}
