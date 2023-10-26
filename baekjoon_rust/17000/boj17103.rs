use std::io::{stdin, stdout, BufWriter, Write};

// 17103번: 골드바흐 파티션
// https://www.acmicpc.net/problem/17103
// 난이도: 실버 2
// 메모리: 14192 KB
// 시간: 72 ms
// 분류: 수학, 정수론, 소수 판정, 에라토스테네스의 체
fn main() {
    let mut writer = BufWriter::new(stdout());
    let mut lines = stdin().lines();

    let mut is_prime = vec![true; 1000001];
    is_prime[0] = false;
    is_prime[1] = false;
    let mut i = 0;
    while i * i <= 1000000 {
        if is_prime[i] {
            let mut j = i * i;
            while j <= 1000000 {
                is_prime[j] = false;
                j += i;
            }
        }
        i += 1;
    }

    let t = line_to_number(&lines.next().unwrap().unwrap());

    for _ in 0..t {
        let n = line_to_number(&lines.next().unwrap().unwrap());

        let mut cnt = 0;
        for j in 2..=n / 2 {
            if is_prime[j as usize] && is_prime[(n - j) as usize] {
                cnt += 1;
            }
        }

        writeln!(writer, "{}", cnt).unwrap();
    }
}

fn line_to_number(s: &str) -> i32 {
    s.parse().unwrap()
}
