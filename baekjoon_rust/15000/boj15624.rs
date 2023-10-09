use std::cmp::min;
use std::io::{stdin, stdout, BufWriter, Write};

// 15624번: 피보나치 수 7
// https://www.acmicpc.net/problem/15624
// 난이도: Silver 4
// 메모리: 13156KB
// 시간: 4ms
// 분류: 수학, 다이나믹 프로그래밍
fn main() {
    let mut writer = BufWriter::new(stdout());
    let mut lines = stdin().lines();

    let n = line_to_number(&lines.next().unwrap().unwrap());

    let modulo = 1000000007;
    let (mut a, mut b) = (0, 1);

    if n == 0 {
        writeln!(writer, "{}", a).unwrap();
        return;
    }

    if n == 1 {
        writeln!(writer, "{}", b).unwrap();
        return;
    }

    for _ in 2..=n {
        let temp = b;
        b = (a + b) % modulo;
        a = temp;
    }

    writeln!(writer, "{}", b).unwrap();
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}

fn line_to_number(s: &str) -> i32 {
    s.parse().unwrap()
}
