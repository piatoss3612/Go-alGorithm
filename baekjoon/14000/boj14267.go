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
	inferior [][]int
	dp       []int
	n, m     int
)

// 메모리: 17432KB
// 시간: 76ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()
	inferior = make([][]int, n+1)
	_ = scanInt()
	for i := 2; i <= n; i++ {
		superior := scanInt()
		inferior[superior] = append(inferior[superior], i) // 직속 부하 직원 추가
	}

	dp = make([]int, n+1)

	// 각 직원이 받은 칭찬 수치를 저장
	for i := 1; i <= m; i++ {
		employee, praise := scanInt(), scanInt()
		dp[employee] += praise
	}

	DFS(1) // 내리 갈굼 문화? 이므로 보스에서부터 말단까지 한 번만 훑어주면서 누적된 칭찬을 더해준다

	for i := 1; i <= n; i++ {
		fmt.Fprintf(writer, "%d ", dp[i])
	}
	fmt.Fprintln(writer)
}

func DFS(sup int) {
	for i := 0; i < len(inferior[sup]); i++ {
		dp[inferior[sup][i]] += dp[sup]
		DFS(inferior[sup][i])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
