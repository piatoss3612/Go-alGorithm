package bj11720

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanBytes)
	var bs1 []byte
	for scanner.Scan() {
		if scanner.Bytes()[0] == '\n' {
			break
		}
		bs1 = append(bs1, scanner.Bytes()[0])
	}
	bs2 := bytes.NewBuffer(bs1)
	s := bs2.String()
	n, _ := strconv.Atoi(s)

	sum := 0

	for i := 0; i < n; i++ {
		scanner.Scan()
		b := scanner.Bytes()[0]
		if b == '0' {
			continue
		}
		sum += int(b - '0')
	}
	fmt.Println(sum)
}
