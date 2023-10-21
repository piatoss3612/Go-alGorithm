use std::io::{stdin, stdout, BufWriter, Write};

// 3474번: 교수가 된 현우
// https://www.acmicpc.net/problem/3474
// 메모리: 13212 KB
// 시간: 20 ms
// 분류: 수학, 정수론
fn main() {
    let mut writer = BufWriter::new(stdout());
    let mut lines = stdin().lines();

    let n = line_to_number(&lines.next().unwrap().unwrap());

    // x!의 0의 개수는 x를 소인수분해 했을 때 5의 개수와 같다. (2의 개수는 항상 5의 개수보다 많으므로)
    for i in 0..n {
        let x = line_to_number(&lines.next().unwrap().unwrap());

        let mut y = 5;
        let mut cnt = 0;

        while y <= x {
            cnt += x / y;
            y *= 5;
        }

        writeln!(writer, "{}", cnt).unwrap();
    }
}

fn line_to_number(s: &str) -> i32 {
    s.parse().unwrap()
}
