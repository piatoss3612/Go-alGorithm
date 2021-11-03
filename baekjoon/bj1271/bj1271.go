package bj1271

import (
	"fmt"
	"math/big"
)

func main() {
	//n := new(big.Int)
	//m := new(big.Int)

	n := big.NewInt(0)
	m := big.NewInt(0)

	fmt.Scan(n, m)
	//fmt.Println(new(big.Int).Div(n, m))
	//fmt.Println(new(big.Int).Mod(n, m))

	fmt.Println(big.NewInt(0).Div(n, m))
	fmt.Print(big.NewInt(0).Mod(n, m))
}
