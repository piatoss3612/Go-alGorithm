package bj1964

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)

	dots := getDots(n)
	fmt.Println(dots % 45678)
}

func getDots(n int64) int64 {
	var sum int64 = 5
	for i := int64(2); i <= n; i++ {
		sum += 3*i + 1
	}
	return sum
}
