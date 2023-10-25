use std::io::{stdin, stdout, BufWriter, Write};

// 15965번. K번째 소수
// https://www.acmicpc.net/problem/15965
// 난이도: 실버 2
// 메모리: 21024 KB
// 시간: 48 ms
// 분류: 수학, 정수론, 소수 판정, 에라토스테네스의 체
fn main() {
    let mut writer = BufWriter::new(stdout());
    let mut lines = stdin().lines();

    let k = line_to_number(&lines.next().unwrap().unwrap());

    let mut is_prime = vec![true; 8000001];
    is_prime[0] = false;
    is_prime[1] = false;
    let mut i = 0;
    while i * i <= 8000000 {
        if is_prime[i] {
            let mut j = i * i;
            while j <= 8000000 {
                is_prime[j] = false;
                j += i;
            }
        }
        i += 1;
    }

    let mut ans = 0;
    let mut cnt = 0;

    for i in 2..=8000000 {
        if is_prime[i] {
            cnt += 1;
            if cnt == k {
                ans = i;
                break;
            }
        }
    }

    writeln!(writer, "{}", ans).unwrap();
}

fn line_to_number(s: &str) -> i32 {
    s.parse().unwrap()
}
