use std::io::{stdin, Write};

// 15810번: 풍선 공장
// https://www.acmicpc.net/problem/15810
// 난이도: Silver 2
// 메모리: 25460KB
// 시간: 288ms
// 분류: 이분 탐색, 매개 변수 탐색
fn main() {
    let mut writer = std::io::BufWriter::new(std::io::stdout());
    let mut lines = stdin().lines();

    let (_, m) = {
        let line = split_line_to_numbers(&lines.next().unwrap().unwrap());
        let n: i32 = line[0];
        let k: i32 = line[1];
        (n, k)
    };

    let arr = split_line_to_numbers(&lines.next().unwrap().unwrap());

    let (mut l, mut r) = (0 as i64, 1e12 as i64);

    while l <= r {
        let mid = (l + r) / 2; // 풍선을 만드는 시간
        let mut sum = 0;
        for &a in &arr {
            sum += mid / a as i64;
        }

        // 풍선을 만드는 시간이 mid일 때 만들 수 있는 풍선의 수가 m보다 크거나 같으면
        if sum >= m as i64 {
            r = mid - 1; // 시간을 줄여서 다시 탐색
        } else {
            // 풍선을 만드는 시간이 mid일 때 만들 수 있는 풍선의 수가 m보다 작으면
            l = mid + 1; // 시간을 늘려서 다시 탐색
        }
    }

    writeln!(writer, "{}", l).unwrap(); // lower bound 출력
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}
