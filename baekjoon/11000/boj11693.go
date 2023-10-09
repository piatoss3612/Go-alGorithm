package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner      = bufio.NewScanner(os.Stdin)
	writer       = bufio.NewWriter(os.Stdout)
	n, m         int
	primeFactors map[int]int // n을 구성하는 소인수 각각의 개수
)

const MOD = 1000000007

// 난이도: Gold 2
// 메모리: 924KB
// 시간: 8ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	n, m = scanInt(), scanInt()
	primeFactors = make(map[int]int)
}

func Solve() {
	/*
		#1. n이 소수인 경우 n**m의 모든 약수의 합 S

		1 + (1 * n) + (1 * (n**2)) + (1 * (n**3)) + ... + (1 * (n ** (m-1))) + (1 * (n**m))

		즉, 첫 항이 1이고 등비가 n인 등비수열의 첫번째 항부터 m+1번째 항까지의 합으로 나타낼 수 있다

		S = ((n**(m+1) - 1) / (n - 1)) mod 1000000007

		여기서 (n-1)의 모듈로 곱셈 역원을 구하여 S를 구할 수 있다
	*/
	/*
		#2. n이 소수가 아닌 경우 -> 소인수 분해

		ex. 6**2의 모든 약수의 합 S를 구하는 경우

		S = (2**0 * 3**0) + (2**0 * 3**1) + (2**0 * 3**2) +
			(2**1 * 3**0) + (2**1 * 3**1) + (2**1 * 3**2) +
			(2**2 * 3**0) + (2**2 * 3**1) + (2**2 * 3**2)
		  = (2**0)(3**0 + 3**1 + 3**2) + (2**1)(3**0 + 3**1 + 3**2) + (2**2)(3**0 + 3**1 + 3**2)
		  = (2**0 + 2**1 + 2**2)(3**0 + 3**1 + 3**2)
		  = (2**2의 약수 합)*(3**2의 약수의 합)
		  = ((2**3-1) *1) * ((3**3-1)*500000004) mod 1000000007
	*/
	/*
		#3. n이 소수가 아니고 동일한 소인수가 여러 번 나오는 경우

		ex. 4**2의 모든 약수의 합 S를 구하는 경우

		4**2 = (2**2)**2 = 2**4이므로
		S = (2**5-1)*1 mod 1000000007

		(2**2의 약수의 합) * (2**2의 약수의 합)으로 구하는 경우는
		(2**0 * 2**1), (2**0 * 2**2), (2**1 * 2**1), (2**1 * 2**2)가 중복되어서 더해지므로 답이 완전히 달라지므로 주의
	*/

	// n이 1인 경우, 답은 항상 1
	if n == 1 {
		fmt.Fprintln(writer, 1)
		return
	}

	// 소인수 분해
	PrimeFactorization()

	ans := 1
	var a, b int

	for prime, cnt := range primeFactors {
		a = Pow(prime, (cnt*m)+1) - 1 // (소인수 ** (소인수의 등장 횟수 * m) + 1) - 1
		b = ExEuclidean(prime - 1)    // (소인수 - 1)의 모듈로 곱셈 역원
		ans *= (a * b) % MOD
		ans %= MOD
	}

	fmt.Fprintln(writer, ans)
}

// n을 소인수 분해하여 n을 구성하는 소인수 각각의 개수를 구한다
func PrimeFactorization() {
	target := n
	for i := 2; i*i <= n; i++ {
		if target == 1 {
			break
		}

		for target%i == 0 {
			primeFactors[i]++
			target /= i
		}
	}

	if target != 1 {
		primeFactors[target]++
	}
}

// 분할 정복을 통해 (x**y) mod 1000000007를 구한다
func Pow(x, y int) int {
	if y == 1 {
		return x
	}

	if y%2 == 0 {
		half := Pow(x, y/2)
		return (half * half) % MOD
	}

	return (Pow(x, y-1) * Pow(x, 1)) % MOD
}

// 확장 유클리드 알고리즘을 사용해 b의 모듈로 곱셈 역원을 구한다
// MOD가 큰 소수(1000000007)이고 b가 MOD보다 작으므로 항상 b의 모듈로 곱셈 역원을 구할 수 있다
func ExEuclidean(b int) int {
	a := MOD
	q, r := 0, 0
	t1, t2, T := 0, 1, 0

	for b > 0 {
		q = a / b
		r = a % b
		a, b = b, r
		T = t1 - t2*q
		t1, t2 = t2, T
	}
	return (t1 + MOD) % MOD // 음수인 경우, 양수로 변환
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
