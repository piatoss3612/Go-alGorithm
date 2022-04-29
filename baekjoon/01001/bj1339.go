package main

import (
	"bufio"
	_ "bytes"
	_ "container/heap"
	"fmt"
	_ "io/ioutil"
	"math"
	_ "math/big"
	"os"
	"sort"
	"strconv"
	_ "strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

// 메모리: 916KB
// 시간: 4ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	apb := make([]int, 26) // 알파벳의 갯수만큼 인덱스를 가지는 슬라이스
	words := make([][]byte, n)

	for i := 0; i < n; i++ {
		words[i] = scanBytes()
		length := len(words[i])

		// A-Z, 0-9에 해당하는 모든 값들을 일일이 대입하면 시간 초과!

		// apb의 각각의 알파벳이 해당하는 인덱스에
		// 알파벳의 자릿수에 해당하는 10의 제곱값을 누적해서 더해준다
		for j, c := range words[i] {
			apb[int(c-'A')] += int(math.Pow(10, float64(length-j-1)))
		}
		/*
			예제 입력:
			2
			GCF
			ACDEB

			apb 슬라이스:
			[10000 1 1010 100 10 1 100 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

			apb 정렬:
			[10000 1010 100 100 10 1 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

			ans:
			10000*9 + 1010*8 + 100*7 + 100*6 + 10*5 + 4 + 3

			출력:
			99437
		*/
	}

	// apb 슬라이스에 더해진 값들은 알파벳 순서와 상관없이 연산의 편의를 위해 오름차순으로 정렬
	sort.Slice(apb, func(i, j int) bool {
		return apb[i] > apb[j]
	})

	ans := 0
	// 가장 큰 수부터 가장 작은 수까지 9~1을 곱해준 값을 결과에 더해준다
	for i := 0; i < 9; i++ {
		ans += apb[i] * (9 - i)
	}

	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
