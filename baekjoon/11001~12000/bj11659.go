package bj11659

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

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	n, m := scanInt(), scanInt()
	inputs := make([]int, n+1)
	for i := 1; i <= n; i++ { // 0에서 시작하면 인덱스 오류 발생 가능
		inputs[i] = scanInt() + inputs[i-1] // 누적합을 구하는 것이므로 입력을 받음과 동시에 이전 항과 더해준다
	}

	for j := 0; j < m; j++ {
		s, e := scanInt(), scanInt()
		result := inputs[e] - inputs[s-1] // s~e까지의 항의 누적합은 e항까지의 누적합에서 s-1항까지의 누적합을 뺀 값
		fmt.Fprintln(writer, result)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
