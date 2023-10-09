use std::collections::VecDeque;
use std::io::{stdin, Write};

// 12761번: 돌다리
// https://www.acmicpc.net/problem/12761
// 난이도: Silver I
// 메모리: 13712KB
// 시간: 8ms
// 분류: 그래프 이론, 그래프 탐색, 너비 우선 탐색
fn main() {
    let mut writer = std::io::BufWriter::new(std::io::stdout());
    let mut lines = stdin().lines();

    let line = split_line_to_numbers(&lines.next().unwrap().unwrap());

    let (a, b, n, m) = (line[0], line[1], line[2], line[3]);

    const MAX: usize = 100000;

    let mut graph = vec![0; MAX + 1];
    let mut visited = vec![false; MAX + 1];

    let mut queue: VecDeque<i32> = VecDeque::new();

    queue.push_back(n);
    visited[n as usize] = true;

    while !queue.is_empty() {
        let here = queue.pop_front().unwrap();

        if here == m {
            writeln!(writer, "{}", graph[here as usize]).unwrap();
            break;
        }

        let nexts = vec![
            here + 1,
            here - 1,
            here + a,
            here - a,
            here + b,
            here - b,
            here * a,
            here * b,
        ];

        for next in nexts {
            if next < 0 || next > MAX as i32 {
                continue;
            }

            if !visited[next as usize] {
                queue.push_back(next);
                visited[next as usize] = true;
                graph[next as usize] = graph[here as usize] + 1;
            }
        }
    }
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}
