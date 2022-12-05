import os.path
import pytest

INPUT_TXT = os.path.join(os.path.dirname(__file__), 'input.txt')


def compute(s: str) -> str:
    # num of columns (change to 3 for example, 9 for my input)
    num_of_stacks = 9
    # num of rows (change to 3 for example, 8 for my input)
    drawing_lines = 8

    parts = s.split("\n\n")
    cargo = parts[0].split("\n")
    procedures = parts[1].split("\n")
    stacks = [[] for x in range(num_of_stacks)]

    for i in range(drawing_lines):
        line = cargo[i]
        crates = line[1::4]
        for s in range(len(crates)):
            if crates[s] != " ":
                stacks[s].append(crates[s])

    # note: we are reading from the top down
    # but the crates have been put in from the bottom up
    # so, we need to reverse to represent real order
    stacks = [stack[::-1] for stack in stacks]

    # move crates in the cargo
    for procedure in procedures:
        tokens = procedure.split(" ")
        n, from_col, to_col = map(int, [tokens[1], tokens[3], tokens[5]])
        # make things 0-indexed
        from_col -= 1
        to_col -= 1
        # add/remove crates
        for x in range(n):
            crate = stacks[from_col].pop()
            stacks[to_col].append(crate)

    return "".join([stack[-1] for stack in stacks])


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
