import os.path
import pytest
from collections import deque

INPUT_TXT = os.path.join(os.path.dirname(__file__), 'input.txt')


def compute(s: str) -> str:
    stack_1 = deque()
    stack_2 = deque()
    stack_3 = deque()
    directions = deque()

    # NOTE: we are reading from the top down
    # but the crates have been put in from the bottom up
    lines = s.splitlines()
    for line in lines:
        split_line = line.split(" ")
        line_len: int = len(split_line)

        # populate directions
        if line.startswith("m"):
            directions.append(int(split_line[1]))
            directions.append(int(split_line[3]))
            directions.append(int(split_line[5]))
        # populate stacks with characters
        else:
            # 9 if one col populated, 6 if two, 3 if all three
            if line_len == 9:
                if split_line[0].startswith("["):
                    stack_1.append(split_line[0][1])
                elif split_line[4].startswith("["):
                    stack_2.append(split_line[4][1])
                elif split_line[8].startswith("["):
                    stack_3.append(split_line[8][1])
            elif line_len == 6:
                if split_line[0].startswith("[") and split_line[5].startswith("["):
                    stack_1.append(split_line[0][1])
                    stack_3.append(split_line[5][1])
                elif split_line[0].startswith("[") and split_line[1].startswith("["):
                    stack_1.append(split_line[0][1])
                    stack_2.append(split_line[1][1])
                elif split_line[4].startswith("[") and split_line[5].startswith("["):
                    stack_2.append(split_line[4][1])
                    stack_3.append(split_line[5][1])
            elif line_len == 3:
                stack_1.append(split_line[0][1])
                stack_2.append(split_line[1][1])
                stack_3.append(split_line[2][1])

    stack_1.reverse()
    stack_2.reverse()
    stack_3.reverse()

    while len(directions) != 0:
        amount = directions.popleft()
        from_col = directions.popleft()
        to_col = directions.popleft()
        # move the crates
        for i in range(0, amount):
            if from_col == 1 and to_col == 2:
                crate: chr = stack_1.pop()
                stack_2.append(crate)
            elif from_col == 1 and to_col == 3:
                crate: chr = stack_1.pop()
                stack_3.append(crate)
            elif from_col == 1 and to_col == 1:
                pass
            elif from_col == 2 and to_col == 1:
                crate: chr = stack_2.pop()
                stack_1.append(crate)
            elif from_col == 2 and to_col == 3:
                crate: chr = stack_2.pop()
                stack_3.append(crate)
            elif from_col == 2 and to_col == 2:
                pass
            elif from_col == 3 and to_col == 1:
                crate: chr = stack_3.pop()
                stack_1.append(crate)
            elif from_col == 3 and to_col == 2:
                crate: chr = stack_3.pop()
                stack_2.append(crate)
            elif from_col == 3 and to_col == 3:
                pass

    return stack_1.pop() + stack_2.pop() + stack_3.pop()


# TODO: change for the small example given
INPUT_S = '''\
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2'''
EXPECTED = "CMZ"


def test() -> None:
    assert compute(INPUT_S) == EXPECTED


def main() -> int:
    with open(INPUT_TXT) as f:
        print(compute(f.read()))

    return 0


if __name__ == '__main__':
    raise SystemExit(main())
