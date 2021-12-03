fn main() {
    let part_1 = include_str!("../input.txt")
        .split("\n")
        .filter_map(|s| s.parse::<usize>().ok())
        .collect::<Vec<usize>>()
        .windows(2)
        .fold(0, |acc, win| if win[0] < win[1] { acc + 1 } else { acc } );

    println!("Part 1: {:?}", part_1);

    let part_2 = include_str!("../input.txt")
        .split("\n")
        .filter_map(|s| s.parse::<usize>().ok())
        .collect::<Vec<usize>>()
        .windows(4)
        .fold(0, |acc, win| {
            if win[0..3].iter().sum::<usize>() < win[1..4].iter().sum::<usize>() {
                acc + 1
            } else {
                acc
            }
        });

    println!("Part 2: {:?}", part_2);
}
