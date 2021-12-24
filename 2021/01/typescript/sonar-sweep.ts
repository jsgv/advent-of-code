import fs from 'fs';

function getItems(): number[] {
    return fs
        .readFileSync('../input.txt', 'utf8')
        .split('\n')
        .map((s: string) => parseInt(s));
}

function part1(): number {
    return getItems().reduce((acc, curr, idx, arr) => {
        if (curr > arr[idx - 1]) {
            return acc + 1;
        }

        return acc;
    }, 0);
}

function part2(): number {
    return getItems().reduce((acc, _, idx, arr) => {
        if (
            arr
                .slice(idx, idx + 3)
                .filter((i) => i)
                .reduce((a, i) => a + i, 0) <
            arr
                .slice(idx + 1, idx + 4)
                .filter((i) => i)
                .reduce((a, i) => a + i, 0)
        ) {
            return acc + 1;
        }

        return acc;
    }, 0);
}

console.log('Part 1:', part1());
console.log('Part 2:', part2());
