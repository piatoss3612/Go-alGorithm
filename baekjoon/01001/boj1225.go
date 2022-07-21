package bj1225

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b []byte
	fmt.Scan(&a, &b)
	var sum1, sum2 int64
	for _, v := range a {
		n, _ := strconv.Atoi(string(v))
		sum1 += int64(n)
	}
	for _, v := range b {
		n, _ := strconv.Atoi(string(v))
		sum2 += int64(n)
	}

	fmt.Println(sum1 * sum2)
}
