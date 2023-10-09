use std::cmp::min;
use std::io::{stdin, stdout, BufWriter, Write};

// 14465번: 소가 길을 건너간 이유 5
// https://www.acmicpc.net/problem/14465
// 난이도: Silver 2
// 메모리: 13156KB
// 시간: 4ms
// 분류: 누적 합, 슬라이딩 윈도우
fn main() {
    let mut writer = BufWriter::new(stdout());
    let mut lines = stdin().lines();

    let (n, k, b) = {
        let line = split_line_to_numbers(&lines.next().unwrap().unwrap());
        (line[0], line[1], line[2])
    };

    let mut broken = vec![false; n as usize + 1]; // 부서진 신호등

    for _ in 0..b {
        let b = line_to_number(&lines.next().unwrap().unwrap());
        broken[b as usize] = true;
    }

    let mut sum = 0; // 부서진 신호등의 개수

    // 1 ~ k번째 신호등 중 부서진 신호등의 개수를 센다.
    for i in 1..=k {
        if broken[i as usize] {
            sum += 1;
        }
    }

    let mut min_sum = sum;

    // i ~ i + k - 1번째 신호등 중 부서진 신호등의 개수를 세어 최솟값을 갱신한다.
    for i in 2..=n - k + 1 {
        if broken[i as usize - 1] {
            sum -= 1;
        }
        if broken[i as usize + k as usize - 1] {
            sum += 1;
        }
        min_sum = min(min_sum, sum);
    }

    writeln!(writer, "{}", min_sum).unwrap(); // 최솟값 출력
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}

fn line_to_number(s: &str) -> i32 {
    s.parse().unwrap()
}
