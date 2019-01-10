use std::collections::HashSet;
use std::time::Instant;

pub fn main() {
    let start = Instant::now();
    let n: u64 = 16777215;

    let mut prev_r1;
    let mut r1;
    let mut r5;
    let mut seen = HashSet::new();

    loop {
        r1 = 123;
        if (r1 & 456) == 72 {
            break
        }
    }

    r1 = 0;

'done:
    loop {
        prev_r1 = r1;

        r5 = prev_r1 | 65536;
        r1 = 8586263;

        loop {
            r1 = (((r1 + (r5 & 255))& n) * 65899) & n;

            if 256 > r5 {
                if seen.is_empty() {
                    println!("Part 1 in {}s: {}", elapsed_str(start.elapsed()), r1);
                }
                if !seen.insert(r1) {
                    println!("Part 2 in {}s: {}", elapsed_str(start.elapsed()), prev_r1);
                    break 'done;
                }
                break;
            } else {
                r5 = r5 / 256;
            }
        }
    }
}

fn elapsed_str(elapsed: std::time::Duration) -> f64 {
    return (elapsed.as_secs() as f64) + (elapsed.subsec_nanos() as f64 / 1000_000_000.0);
}
