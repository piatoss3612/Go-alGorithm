package main

import (
	"bufio"
	_ "bytes"
	"fmt"
	_ "io/ioutil"
	_ "math"
	_ "math/big"
	"os"
	_ "sort"
	"strconv"
	_ "strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	visited []bool
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m := scanInt(), scanInt()
	dict1 := make(map[int]string) // 포켓몬 이름: 도감번호
	dict2 := make(map[string]int) // 도감번호: 포켓몬 이름
	for i := 1; i <= n; i++ {
		pokemon := scanString()
		dict1[i] = pokemon
		dict2[pokemon] = i
	}
	for i := 1; i <= m; i++ {
		scanner.Scan()
		tmp := scanner.Text()
		n, err := strconv.Atoi(scanner.Text())
		if err != nil { // 정수로 변환할 때 에러가 발생했다면 포켓몬 이름이 입력된 것
			// 따라서 이름: 도감 번호 쌍이 저장되어 있는 맵에서 데이터를 읽어 온다
			fmt.Fprintln(writer, dict2[tmp])
		} else {
			// 정수로 오류없이 변환된다면 도감번호: 이름 쌍이 저장되어 있는 맵에서 데이터를 읽어 온다
			fmt.Fprintln(writer, dict1[n])
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}
