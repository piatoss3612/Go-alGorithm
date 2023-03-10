use std::io;

// 난이도: Silver 5
// 메모리: 13156KB???
// 시간: 4ms
// 분류: 그리디 알고리즘, 정렬
// 러스트 왤케 어려워요;;
fn main() {
    let (p, _) = two_numbers();
    let mut fatigue = scan_numbers();

    fatigue.sort();

    let mut total_fatigue = p;
    let mut cnt = 0;

    for f in fatigue.iter() {
        if total_fatigue < 200 {
            total_fatigue += f;
            cnt += 1;
        } else {
            break;
        }
    }

    println!("{}", cnt)
}

fn scan_line() -> String {
    let mut buf = String::new();
    io::stdin()
        .read_line(&mut buf)
        .expect("unable to read line");

    buf
}

fn scan_numbers() -> Vec<i32> {
    let numbers = scan_line()
        .trim()
        .split(" ")
        .map(|x| x.parse().unwrap())
        .collect();
    numbers
}

fn two_numbers() -> (i32, i32) {
    match scan_numbers()[..] {
        [a, b] => return (a, b),
        _ => panic!("wrong input"),
    }
}
