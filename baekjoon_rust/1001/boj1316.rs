use std::io;

// 난이도: Silver 5
// 메모리: 13152KB
// 시간: 4ms
// 분류: 구현, 문자열
fn main() {
    let n = scan_numbers()[0];

    let mut cnt = 0; // 그룹 단어의 개수

    for _ in 0..n {
        let s = scan_line(); // 단어 스캔

        let mut check = [false; 26]; // 알파벳 소문자가 이전에 등장했는지 체크

        let mut prev = s.chars().nth(0).unwrap(); // 인덱스가 0인 문자를 이전 문자로 설정
        check[prev as usize - 'a' as usize] = true; // 인덱스가 0인 문자 등장 여부 체크

        let mut is_group = true; // 그룹 단어인지 여부

        // 인덱스가 1인 문자부터 검사
        for c in s.chars().skip(1) {
            // 이전 문자와 다른 문자가 등장
            if prev != c {
                // 이전에 등장한 문자라면 그룹 단어가 아님
                if check[c as usize - 'a' as usize] {
                    is_group = false;
                    break;
                } else {
                    // 이전에 등장하지 않은 문자라면 등장 여부 체크
                    check[c as usize - 'a' as usize] = true;
                    prev = c; // 이전 문자를 현재 문자로 설정
                }
            }
        }

        // 그룹 단어라면 카운트 증가
        if is_group {
            cnt += 1;
        }
    }

    println!("{}", cnt); // 그룹 단어의 개수 출력
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
