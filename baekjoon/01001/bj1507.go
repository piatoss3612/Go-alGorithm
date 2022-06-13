package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	fw        [21][21]int  // 플로이드 와샬 결과 그래프
	notOrigin [21][21]bool // 원래 경로인지 아닌지 판별
)

// 메모리: 916KB
// 시간: 4ms
// 플로이드 와샬 알고리즘의 결과를 이용해 역으로 원래의 경로와 가중치의 합을 구하는 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fw[i][j] = scanInt()
		}
	}

	flag := true

Loop:
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if i == j || j == k || k == i {
					continue
				}

				// 플로이드 와샬 알고리즘이 성립되지 않는 경우
				if fw[i][j] > fw[i][k]+fw[k][j] {
					flag = false
					break Loop
				}

				// i에서 k를 거쳐 j로 가는 경로의 가중치가 i에서 j로 가는 경로의 가중치와 같다면
				// i->k->j는 최단 경로이므로 i->j로 가는 경로를 제외한다
				if fw[i][j] == fw[i][k]+fw[k][j] {
					notOrigin[i][j] = true
				}
			}
		}
	}

	if !flag {
		fmt.Fprintln(writer, -1)
		return
	}

	sum := 0

	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			if !notOrigin[i][j] {
				sum += fw[i][j]
			}
		}
	}

	fmt.Fprintln(writer, sum)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
