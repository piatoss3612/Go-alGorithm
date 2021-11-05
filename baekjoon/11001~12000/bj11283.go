package bj11283

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var hangul []byte
	fmt.Scan(&hangul)
	r, _ := utf8.DecodeRune(hangul)
	fmt.Println(r - 44031)
}
