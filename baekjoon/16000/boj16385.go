package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner             = bufio.NewScanner(os.Stdin)
	writer              = bufio.NewWriter(os.Stdout)
	N                   int
	uniquePokemon       map[string]int   // 포켓몬의 이름과 아이디를 맵핑
	TheNumberOfPokemons int              // 유니크 포켓몬의 수
	pokeStops           [21]PokeStop     // 포케스탑 정보
	dp                  [21][1 << 16]int // 유니크 포켓몬을 모두 잡는데 필요한 이동거리의 최솟값을 메모이제이션
)

type PokeStop struct {
	r, c    int
	pokemon string
}

const INF = 987654321

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	uniquePokemon = make(map[string]int)

	for i := 1; i <= N; i++ {
		r, c, pokemon := scanInt(), scanInt(), scanString()
		_, exist := uniquePokemon[pokemon]
		// 입력받은 포켓몬 이름이 맵에 등록되지 않은 경우
		if !exist {
			TheNumberOfPokemons++                        // 포켓몬의 수 1증가
			uniquePokemon[pokemon] = TheNumberOfPokemons // 포켓몬에게 아이디 부여
		}
		pokeStops[i] = PokeStop{r, c, pokemon}
	}

	for i := 0; i < 21; i++ {
		for j := 0; j < 1<<16; j++ {
			dp[i][j] = -1
		}
	}
}

func Solve() {
	ans := rec(0, 1) // 항상 (0, 0)에서 시작하므로, 0번 포케스탑이 (0, 0)에 위치해 있다고 가정하고 아이디가 0번인 포켓몬을 잡은 상태에서 순회 시작
	fmt.Fprintln(writer, ans)
}

func rec(currPokeStop, gotcha int) int {
	// 기저 사례: 모든 포켓몬을 잡은 경우
	if gotcha == (1<<(TheNumberOfPokemons+1))-1 {
		return dist(currPokeStop, 0) // 시작 위치로 되돌아가는 거리 추가
	}

	// 기저 사례2: 이미 이동 거리의 최솟값을 구한 경우
	ret := &dp[currPokeStop][gotcha]
	if *ret != -1 {
		return *ret
	}

	*ret = INF

	for next := 1; next <= N; next++ {
		// next 포케스탑에 위치한 포켓몬을 아직 잡지 않은 경우
		if pokemonID := uniquePokemon[pokeStops[next].pokemon]; gotcha&(1<<pokemonID) == 0 {
			*ret = min(*ret, rec(next, gotcha+(1<<pokemonID))+dist(currPokeStop, next))
		}
	}

	return *ret
}

func dist(a, b int) int {
	return abs(pokeStops[a].r-pokeStops[b].r) + abs(pokeStops[a].c-pokeStops[b].c)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}
