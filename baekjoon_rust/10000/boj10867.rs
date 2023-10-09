use std::io::{stdin, Write};

// 10867번: 중복 빼고 정렬하기
// https://www.acmicpc.net/problem/10867
// 메모리: 14196KB
// 시간: 8ms
// 분류: 정렬
fn main() {
    let mut lines = stdin().lines();

    let line = split_line_to_numbers(&lines.next().unwrap().unwrap());
    let n = line[0];
    let mut arr = split_line_to_numbers(&lines.next().unwrap().unwrap());

    arr.sort_unstable(); // 중복 제거를 위해 정렬, 순서는 중요하지 않음

    let mut writer = std::io::BufWriter::new(std::io::stdout());

    for i in 0..n {
        // 중복 제거
        if i != 0 && arr[i as usize] == arr[(i - 1) as usize] {
            continue;
        }
        write!(writer, "{} ", arr[i as usize]).unwrap();
    }
    write!(writer, "\n").unwrap();
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}
