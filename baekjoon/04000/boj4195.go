package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	friends map[string]friend
)

type friend struct {
	parent string
	count  int
}

// memory: 68444KB -> 48812KB
// time: 224ms -> 212ms
// solution using disjoint set and map data structure
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	t := scanInt()
	for i := 1; i <= t; i++ {
		TestCase()
	}
}

// test case
func TestCase() {
	n := scanInt()
	friends = make(map[string]friend)

	var x, y string
	for i := 1; i <= n; i++ {
		x, y = scanName(), scanName()

		_, ok := friends[x]
		// if x is not included in friends relationship yet
		if !ok {
			friends[x] = friend{x, 1}
		}

		_, ok = friends[y]
		// if y is not included in friends relationship yet
		if !ok {
			friends[y] = friend{y, 1}
		}

		union(x, y)                                  // union set x and set y
		fmt.Fprintln(writer, friends[find(x)].count) // print parent of x's friends count
	}
}

// find parent of x
func find(x string) string {
	if friends[x].parent == x {
		return x
	}
	friends[x] = friend{find(friends[x].parent), friends[x].count}
	return friends[x].parent
}

// union set x and set y
func union(x, y string) {
	x = find(x) // get parent of x
	y = find(y) // get parent of y
	if x != y { // if parent of x is not equal to parent of y
		friends[y] = friend{x, friends[y].count}                                    // set parent of y to x
		friends[x] = friend{friends[x].parent, friends[x].count + friends[y].count} // get total number of friends
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanName() string {
	scanner.Scan()
	return scanner.Text()
}
