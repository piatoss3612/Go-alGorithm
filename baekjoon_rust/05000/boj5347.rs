use std::io::{stdin, Write};

// 5347번: LCM
// https://www.acmicpc.net/problem/5347
// 메모리: 13156KB
// 시간: 4ms
// 분류: 수학, 정수론, 유클리드 호제법
fn main() {
    let mut lines = stdin().lines();

    let n = split_line_to_numbers(&lines.next().unwrap().unwrap())[0] as i64;

    let mut writer = std::io::BufWriter::new(std::io::stdout());

    for _ in 0..n {
        let line = split_line_to_numbers(&lines.next().unwrap().unwrap());
        let (a, b) = (line[0] as i64, line[1] as i64);
        writeln!(writer, "{}", lcm(a, b)).unwrap();
    }
}

// 최소공배수를 구하는 함수: a * b / gcd(a, b)
fn lcm(a: i64, b: i64) -> i64 {
    a * b / gcd(a, b)
}

// 최대공약수를 구하는 함수
fn gcd(a: i64, b: i64) -> i64 {
    if b == 0 {
        return a;
    }
    gcd(b, a % b)
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}
