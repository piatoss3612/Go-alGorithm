package bj10814

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

type Member struct {
	Age  int
	Name string
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	var members []Member

	for i := 0; i < n; i++ {
		members = append(members, scanMember())
	}
	sort.SliceStable(members, func(i, j int) bool {
		return members[i].Age < members[j].Age
	})

	for _, m := range members {
		fmt.Fprintf(writer, "%d %s\n", m.Age, m.Name)
	}
}

func scanMember() Member {
	scanner.Scan()
	age, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	name := scanner.Text()
	return Member{age, name}
}
