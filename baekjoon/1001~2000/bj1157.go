package bj1157

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()

	letters, _ := ioutil.ReadAll(reader)
	letters = letters[:len(letters)-1]
	letters = bytes.ToUpper(letters)

	counts := make(map[byte]int)
	for _, v := range letters {
		_, ok := counts[v]
		if ok {
			counts[v]++
		} else {
			counts[v] = 1
		}
	}

	var key byte
	max := 0

	for k, v := range counts {
		if v > max {
			key = k
			max = v
		} else if v == max {
			key = '?'
		}
	}

	writer.WriteString(string(key))
}
