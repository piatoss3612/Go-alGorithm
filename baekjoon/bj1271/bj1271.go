package bj1271

import (
	"fmt"
	"math/big"
)

func main() {
	n := new(big.Int)
	m := new(big.Int)

	fmt.Scan(n, m)
	fmt.Println(new(big.Int).Div(n, m))
	fmt.Println(new(big.Int).Mod(n, m))
}
