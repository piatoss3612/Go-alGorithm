package bj1110

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n < 10 {
		n *= 10
	}

	temp := n

	for cnt := 0; true; cnt++ {

		n = (n%10)*10 + (((n / 10) + (n % 10)) % 10)

		if temp == n {
			cnt++
			fmt.Println(cnt)
			break
		}
	}
}
