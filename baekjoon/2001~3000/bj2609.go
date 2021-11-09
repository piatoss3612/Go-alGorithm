package bj2609

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	gcd := getGcd(a, b)
	lcm := getLcm(a, b, gcd)
	fmt.Println(gcd)
	fmt.Println(lcm)
}

func getGcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func getLcm(a, b, gcd int) int {
	return a * b / gcd
}
