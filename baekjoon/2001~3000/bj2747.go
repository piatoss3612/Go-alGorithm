package bj2747

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

func main() {
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	result := fibonacci2(n)
	fmt.Println(result)
}

func fibonacci2(n int) int64 {
	if n <= 2 {
		return 1
	}
	var a, b int64 = 1, 1
	for i := 3; i <= n; i++ {
		temp := b
		b = a + b
		a = temp
	}
	return b
}
