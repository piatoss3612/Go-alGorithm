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
	digits, _ := ioutil.ReadAll(reader)
	if digits[len(digits)-1] == '\n' {
		digits = digits[:len(digits)-1]
	}

	var sum int
	cnt := 1

	if len(digits) == 1 {
		sum = int(digits[0] - '0')
	} else {
		for _, v := range digits {
			sum += int(v - '0')
		}
	}

	s := strconv.Itoa(sum)
	for len(s) > 1 {
		sum = 0
		newDigits := strings.Split(s, "")
		for _, v := range newDigits {
			d, _ := strconv.Atoi(v)
			sum += d
		}
		s = strconv.Itoa(sum)
		cnt++
	}

	fmt.Fprintf(writer, "%d\n", cnt)

	if s == "3" || s == "6" || s == "9" {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
	}
}
