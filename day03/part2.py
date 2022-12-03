import os.path
import pytest

INPUT_TXT = os.path.join(os.path.dirname(__file__), 'input.txt')


def set_intersection(set1: set, set2: set) -> list[str]:
    return [x for x in set1 if x in set2]


def compute(s: str) -> int:
    lines = s.splitlines()
    pointTotal = 0

    for line in lines:
        ruck1 = line[:len(line)//2]
        ruck2 = line[len(line)//2:]

        # convert to sets
        ruck1set = set(ruck1)
        ruck2set = set(ruck2)

        intersection: list(str) = set_intersection(ruck1set, ruck2set)

        # if there is an intersection, do the stuff
        if len(intersection) == 1:
            letter: str = intersection[0]
            print(intersection)
            print(ord(letter))
            if letter.islower() and letter.isalpha():
                pointTotal += ord(letter) - 96
            elif letter.isupper() and letter.isalpha():
                pointTotal += ord(letter) - 38
        else:
            raise Exception("no intersection in rucksacks!")
            return 1
    return pointTotal

# TODO: change for the small example given
INPUT_S = '''\
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw'''
EXPECTED = 70


def test() -> None:
    assert compute(INPUT_S) == EXPECTED


def main() -> int:
    with open(INPUT_TXT) as f:
        print(compute(f.read()))

    return 0


if __name__ == '__main__':
    raise SystemExit(main())
