use std::io;

// 11004번: K번째 수
// https://www.acmicpc.net/problem/11004
fn main() {
    let (_, k) = {
        let v = scan_numbers();
        (v[0], v[1])
    };

    let mut arr = scan_numbers();

    /*
    // stable sort
    // 메모리: 129392KB
    // 시간: 972ms

    arr.sort();

    println!("{}", arr[k as usize - 1]);
    */

    // unstable sort
    // 메모리: 129392KB
    // 시간: 520ms
    println!("{}", arr.select_nth_unstable(k as usize - 1).1);
}

fn scan_line() -> String {
    let mut buf = String::new();
    io::stdin().read_line(&mut buf).unwrap();
    buf.trim().to_string()
}

fn scan_numbers() -> Vec<i32> {
    scan_line()
        .split_whitespace()
        .map(|s| s.parse().unwrap())
        .collect()
}
