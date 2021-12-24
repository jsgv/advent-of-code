import fs from 'fs';

function getItems(): { direction: string; value: number }[] {
    return fs
        .readFileSync('../input.txt', 'utf8')
        .split('\n')
        .filter((s: string) => s)
        .map((s: string) => {
            const parts = s.split(' ');
            return {
                direction: parts[0],
                value: parseInt(parts[1]),
            };
        });
}

class Location {
    horizontal: number;
    depth: number;
    aim: number;

    constructor() {
        this.horizontal = 0;
        this.depth = 0;
        this.aim = 0;
    }

    answer(): number {
        return this.horizontal * this.depth;
    }
}

function part1(): number {
    return getItems().reduce((acc, i) => {
        switch (i.direction) {
            case "up":
                acc.depth -= i.value;
                break;
            case "down":
                acc.depth += i.value;
                break;
            case "forward":
                acc.horizontal += i.value;
                break;
        }

        return acc;
    }, new Location()).answer();
}

function part2(): number {
    return getItems().reduce((acc, i) => {
        switch (i.direction) {
            case "up":
                acc.aim -= i.value;
                break;
            case "down":
                acc.aim += i.value;
                break;
            case "forward":
                acc.horizontal += i.value;
                acc.depth += i.value * acc.aim;
                break;
        }

        return acc;
    }, new Location()).answer();
}

console.log('Part 1:', part1());
console.log('Part 2:', part2());
