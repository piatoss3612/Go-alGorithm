use std::io::{stdin, Write};

// 13301번: 타일 장식물
// https://www.acmicpc.net/problem/13301
// 난이도: Silver 5
// 메모리: 13152KB
// 시간: 4ms
// 분류: 수학, 다이나믹 프로그래밍
fn main() {
    let mut lines = stdin().lines();

    let line = split_line_to_numbers(&lines.next().unwrap().unwrap());
    let n = line[0];

    let mut fibo: Vec<i64> = vec![0; 82]; // 피보나치 수열 메모이제이션
    fibo[1] = 1; // 피보나치 수열 첫번째 원소는 1
    for i in 2..=n + 1 {
        fibo[i as usize] = fibo[(i - 1) as usize] + fibo[(i - 2) as usize]; // 피보나치 수열의 점화식
    }

    println!("{}", fibo[n as usize] * 2 + fibo[(n + 1) as usize] * 2); // n번째 사각형의 둘레
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}
