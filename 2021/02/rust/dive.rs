use std::str::FromStr;
use std::num::ParseIntError;

struct Movement {
    direction: String,
    value: u64,
}

impl FromStr for Movement {
    type Err = ParseIntError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let parts: Vec<&str> = s.split(" ").collect();

        let direction = parts[0].to_string();
        let value = parts[1].parse::<u64>().unwrap();

        Ok(Movement{ direction, value })
    }
}

#[derive(Copy, Clone)]
struct Location {
    horizontal: u64,
    depth: u64,
    aim: u64,
}

impl Location {
    fn answer(self) -> u64 {
        self.horizontal * self.depth
    }
}

fn main() {
    let part_1 = include_str!("../input.txt")
        .split_terminator("\n")
        .filter_map(|s| s.parse::<Movement>().ok())
        .collect::<Vec<Movement>>()
        .iter()
        .fold(Location{ horizontal: 0, depth: 0, aim: 0 }, |acc, i| match i {
            Movement { ref direction, .. }  if direction == "up" => Location {
                depth: acc.depth - i.value,
                ..acc
            },
            Movement { ref direction, .. } if direction == "down" => Location {
                depth: acc.depth + i.value,
                ..acc
            },
            Movement { ref direction, .. } if direction == "forward" => Location {
                horizontal: acc.horizontal + i.value,
                ..acc
            },
            _ => panic!("no case"),
        });

    let part_2 = include_str!("../input.txt")
        .split_terminator("\n")
        .filter_map(|s| s.parse::<Movement>().ok())
        .collect::<Vec<Movement>>()
        .iter()
        .fold(Location{ horizontal: 0, depth: 0, aim: 0 }, |acc, i| match i {
            Movement { ref direction, .. }  if direction == "up" => Location {
                aim: acc.aim - i.value,
                ..acc
            },
            Movement { ref direction, .. } if direction == "down" => Location {
                aim: acc.aim + i.value,
                ..acc
            },
            Movement { ref direction, .. } if direction == "forward" => Location {
                horizontal: acc.horizontal + i.value,
                depth: acc.depth + (i.value * acc.aim),
                ..acc
            },
            _ => panic!("no case"),
        });

    assert_eq!(part_1.answer(), 1868935);
    assert_eq!(part_2.answer(), 1965970888);

    println!("Part 1: {:?}", part_1.answer());
    println!("Part 2: {:?}", part_2.answer());
}

