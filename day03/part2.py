import os.path
import pytest

INPUT_TXT = os.path.join(os.path.dirname(__file__), 'input.txt')


def set_intersection(set1: set, set2: set, set3: set) -> list[str]:
    return [x for x in set1 if (x in set2 and x in set3)]


def compute(s: str) -> int:
    lines = s.splitlines()
    point_total = 0
    chunk_size = 3

    for i in range(0, len(lines), chunk_size):
        # grab 3 rucksacks at a time
        three_rucksacks = lines[i:i+chunk_size]
        set1 = set(three_rucksacks[0])
        set2 = set(three_rucksacks[1])
        set3 = set(three_rucksacks[2])
        intersection: list(str) = set_intersection(set1, set2, set3)

        # if there is an intersection, do the stuff
        if len(intersection) == 1:
            letter: str = intersection[0]
            print(intersection)
            print(ord(letter))
            if letter.islower() and letter.isalpha():
                point_total += ord(letter) - 96
            elif letter.isupper() and letter.isalpha():
                point_total += ord(letter) - 38
        else:
            raise Exception("no intersection in rucksacks!")
            return 1

        print(three_rucksacks)
    return point_total

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
