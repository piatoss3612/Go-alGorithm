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
	M, N, K int
	Jungle  [1001][1001]int // Jungle[a][b]: (0,0)에서 (a,b)까지 정글의 개수
	Ocean   [1001][1001]int // Ocean[a][b]: (0,0)에서 (a,b)까지 바다의 개수
	Ice     [1001][1001]int // Ice[a][b]: (0,0)에서 (a,b)까지 얼음의 개수
)

// 난이도: Gold 5
// 메모리: 28628KB
// 시간: 148ms
// 분류: 누적 합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	M, N, K = scanInt(), scanInt(), scanInt()

	for i := 1; i <= M; i++ {
		line := scanBytes()

		for j := 1; j <= N; j++ {
			here := line[0]
			line = line[1:]

			switch here {
			case 'J':
				Jungle[i][j] += 1
			case 'O':
				Ocean[i][j] += 1
			case 'I':
				Ice[i][j] += 1
			}

			// (0,0)을 왼쪽 위 모서리, (i,j) 오른쪽 아래 모서리로 갖는 직사각형에 포함된 정글, 바다, 얼음의 누적 개수의 합 구하기
			Jungle[i][j] += Jungle[i-1][j] + Jungle[i][j-1] - Jungle[i-1][j-1]
			Ocean[i][j] += Ocean[i-1][j] + Ocean[i][j-1] - Ocean[i-1][j-1]
			Ice[i][j] += Ice[i-1][j] + Ice[i][j-1] - Ice[i-1][j-1]
		}
	}
}

func Solve() {
	var a, b, c, d int

	for i := 1; i <= K; i++ {
		a, b, c, d = scanInt(), scanInt(), scanInt(), scanInt()
		j := junglePartialSum(a, b, c, d)
		o := oceanPartialSum(a, b, c, d)
		i := icePartialSum(a, b, c, d)
		fmt.Fprintln(writer, j, o, i)
	}
}

// (a,b)에서 (c,d)에 포함된 정글의 개수를 반환한다
func junglePartialSum(a, b, c, d int) int {
	return Jungle[c][d] - Jungle[c][b-1] - Jungle[a-1][d] + Jungle[a-1][b-1]
}

// (a,b)에서 (c,d)에 포함된 바다의 개수를 반환한다
func oceanPartialSum(a, b, c, d int) int {
	return Ocean[c][d] - Ocean[c][b-1] - Ocean[a-1][d] + Ocean[a-1][b-1]
}

// (a,b)에서 (c,d)에 포함된 얼음의 개수를 반환한다
func icePartialSum(a, b, c, d int) int {
	return Ice[c][d] - Ice[c][b-1] - Ice[a-1][d] + Ice[a-1][b-1]
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
