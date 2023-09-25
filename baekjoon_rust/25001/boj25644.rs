use std::io::{stdin, Write};

// 25644번: 최대 상승
// https://www.acmicpc.net/problem/25644
// 난이도: Silver 5
// 메모리: 16240KB
// 시간: 24ms
// 분류: 다이나믹 프로그래밍
fn main() {
    let mut lines = stdin().lines();

    let n = split_line_to_numbers(&lines.next().unwrap().unwrap())[0];

    let arr = split_line_to_numbers(&lines.next().unwrap().unwrap());

    let mut dp = vec![0; n as usize];

    // dp[i] = arr[0..i] 중 arr[i]보다 작은 값들 중 가장 작은 값의 인덱스
    for i in 1..n {
        if arr[i as usize] < arr[dp[i as usize - 1] as usize] {
            dp[i as usize] = i;
        } else {
            dp[i as usize] = dp[i as usize - 1];
        }
    }

    let mut ans = 0;

    // arr[i] - arr[dp[i]] 중 가장 큰 값이 정답
    for i in 0..n {
        let diff = arr[i as usize] - arr[dp[i as usize] as usize];
        if diff > ans {
            ans = diff;
        }
    }

    let mut writer = std::io::BufWriter::new(std::io::stdout());
    writeln!(writer, "{}", ans).unwrap();
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}
