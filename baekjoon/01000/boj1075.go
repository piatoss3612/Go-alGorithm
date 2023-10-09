package bj1075

import (
	"fmt"
)

func main() {
	var n int64
	var f int
	fmt.Scan(&n)
	fmt.Scan(&f)

	d := n % int64(f)
	c := n % 100

	c -= d
	if c < 0 {
		c += int64(f)
	}
	c %= int64(f)

	fmt.Printf("%02d\n", c)

}
