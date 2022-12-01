import os.path
import pytest

INPUT_TXT = os.path.join(os.path.dirname(__file__), 'input.txt')


def compute(s: str) -> int:
    groups = s.split("\n\n")

    elfCalCount = [0, 0, 0]

    for group in groups:
        lines = group.split("\n")
        sum = 0
        for line in lines:
            sum += int(line)
        if sum > elfCalCount[0]:
            elfCalCount[2] = elfCalCount[1]
            elfCalCount[1] = elfCalCount[0]
            elfCalCount[0] = sum
        elif sum > elfCalCount[1]:
            elfCalCount[2] = elfCalCount[1]
            elfCalCount[1] = sum
        elif sum > elfCalCount[2]:
            elfCalCount[2] = sum

    return elfCalCount[0]

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

10000\
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
