package bj2845

import "fmt"

func main() {
	var l, p int
	var a, b, c, d, e int

	fmt.Scan(&l, &p)
	fmt.Scan(&a, &b, &c, &d, &e)

	lp := l * p
	fmt.Printf("%d %d %d %d %d", a-lp, b-lp, c-lp, d-lp, e-lp)
}
