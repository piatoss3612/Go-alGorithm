use std::io::{stdin, stdout, BufWriter, Write};

// 17427번: 약수의 합 2
// https://www.acmicpc.net/problem/17427
// 난이도: 실버 2
// 메모리: 13208 KB
// 시간: 8 ms
// 분류: 수학, 정수론
fn main() {
    let mut writer = BufWriter::new(stdout());
    let mut lines = stdin().lines();

    let n = line_to_number(&lines.next().unwrap().unwrap());

    /*
       f(1) = 1
       f(2) = 1 + 2
       f(3) = 1 + 3
       f(4) = 1 + 2 + 4
       f(5) = 1 + 5
       f(6) = 1 + 2 + 3 + 6

       g(1) = 1
       g(2) = 1*2 + 2
       g(3) = 1*3 + 2 + 3
       g(4) = 1*4 + 2*2 + 3 + 4
       g(5) = 1*5 + 2*2 + 3 + 4 + 5
       g(6) = 1*6 + 2*3 + 3*2 + 4 + 6
       ...
       g(n) = 1*n + 2*(n/2) + 3*(n/3) + ... + n
    */

    let mut sum: i64 = 0;

    for i in 1..=n {
        sum += i * (n / i);
    }

    writeln!(writer, "{}", sum).unwrap();
}

fn line_to_number(s: &str) -> i64 {
    s.parse().unwrap()
}
