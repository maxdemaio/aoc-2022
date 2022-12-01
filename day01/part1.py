import os.path
import pytest

INPUT_TXT = os.path.join(os.path.dirname(__file__), 'input.txt')


def compute(s: str) -> int:
    lines = s.splitlines()
    print(lines)

    elves = {}
    i = 0

    for line in lines:
        # new line, next elf
        if line == "":
            i += 1
            continue
        if i in elves:
            elves[i] += int(line)
        else:
            elves[i] = 0
            elves[i] += int(line)

    return max(elves.values())

INPUT_S = '''\
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
'''
EXPECTED = 24000


def test() -> None:
    assert compute(INPUT_S) == EXPECTED


def main() -> int:
    with open(INPUT_TXT) as f:
        print(compute(f.read()))

    return 0


if __name__ == '__main__':
    raise SystemExit(main())
