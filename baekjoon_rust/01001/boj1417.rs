use std::{
    collections::BinaryHeap,
    io::{stdin, Write},
};

// 1417번: 국회의원 선거
// https://www.acmicpc.net/problem/1417
// 난이도: Silver 5
// 메모리: 13156KB
// 시간: 4ms
// 분류: 그리디 알고리즘, 우선순위 큐
fn main() {
    let mut lines = stdin().lines();

    let n = split_line_to_numbers(&lines.next().unwrap().unwrap())[0];

    let mut heap = BinaryHeap::with_capacity(n as usize); // 최대 힙 초기화

    let mut dasom = split_line_to_numbers(&lines.next().unwrap().unwrap())[0]; // 다솜이의 득표수

    // 나머지 후보들의 득표수를 힙에 넣는다.
    for _ in 0..n - 1 {
        let input = split_line_to_numbers(&lines.next().unwrap().unwrap())[0];
        heap.push(input);
    }

    let mut ans = 0;

    let mut writer = std::io::BufWriter::new(std::io::stdout());

    // 다솜이의 득표수가 힙의 최대값보다 커질 때까지 반복한다.
    while let Some(max) = heap.pop() {
        if dasom > max {
            break;
        }

        dasom += 1;
        ans += 1;

        heap.push(max - 1);
    }

    writeln!(writer, "{}", ans).unwrap();
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}
