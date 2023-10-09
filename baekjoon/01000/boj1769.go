package bj1769

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	bytes, _ := ioutil.ReadAll(reader)
	if bytes[len(bytes)-1] == '\n' {
		bytes = bytes[:len(bytes)-1]
	}

	sum := 0
	cnt := 0

	if len(bytes) == 1 {
		sum = int(bytes[0] - '0')
	} else {
		for _, v := range bytes {
			sum += int(v - '0')
		}
		cnt++
	}
	for sum > 9 {
		tmp := 0
		s := strconv.Itoa(sum)
		digits := strings.Split(s, "")
		for _, d := range digits {
			n, _ := strconv.Atoi(d)
			tmp += n
		}
		sum = tmp
		cnt++
	}

	fmt.Fprintln(writer, cnt)
	if sum == 3 || sum == 6 || sum == 9 {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
	}
}
