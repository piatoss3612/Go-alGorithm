package main

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
	n, k := scanInt(), scanInt()
	input := make([]int, n+1)
	for i := 1; i <= n; i++ {
		input[i] = scanInt()
	}

	var diff []int
	for i := 2; i <= n; i++ {
		diff = append(diff, input[i]-input[i-1])
	}

	/*
		입력:
		7 3

		1 3 5 6 10 11 16

		프로세스:
		(1 3 5 6), (10 11), 16

		출력:
		6

		해설:
		최소 비용을 구하기 위해서는 먼저 인접한 유치원생들의 키 차이를 구해야한다
		인접한 원생들의 키 차이를 diff 슬라이스에 저장하고 오름차순으로 정렬
		여기서 diff 값을 때 사용한 좌표들은 중복되지 않음
		1번째 & 4번째 원생의 차이 =
			1번째 & 2번째 원생의 차이 + 2번째 & 3번째 원생의 차이 + 3번째 & 4번째 원생의 차이

		오름차순으로 정렬한 diff의 앞에서부터 n-k개의 값을 가져와 더한 값이 결과

		n-k:
			n명의 사람이 각각 조를 이루는 경우: 총 n개의 조
			n명의 사람 중 인접한 2개의 값을 1개의 조로 묶는 경우: 총 n - 1개의 조
			...
			n명의 사람 중 인접한 2개의 값을 k개의 조로 묶는 경우: 총 n - k개의 조
	*/

	sort.Ints(diff)
	ans := 0

	for i := 0; i < n-k; i++ {
		ans += diff[i]
	}

	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
