package bj1946

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	t := scanInt()
	for i := 0; i < t; i++ {
		n := scanInt()
		ranks := make([][]int, n)
		for j := 0; j < n; j++ {
			ranks[j] = append(ranks[j], scanInt())
			ranks[j] = append(ranks[j], scanInt())
		}
		// 두 개의 성적 중 하나를 오름차순으로 정렬
		sort.Slice(ranks, func(i, j int) bool {
			return ranks[i][0] < ranks[j][0]
		})
		cnt := n
		flag := ranks[0][1]
		// 정렬하지 않은 성적의 순위는 앞에 있는 모든 성적보다 반드시 높아야 신입사원으로 채용된다
		for k := 1; k < n; k++ {
			if ranks[k][1] < flag {
				// 모든 성적을 조회하면 시간이 초과되기 때문에 더 높은 순위가 나왔을 때만 기준을 바꿔준다
				flag = ranks[k][1]
			} else {
				cnt -= 1
			}
		}
		fmt.Fprintln(writer, cnt)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
