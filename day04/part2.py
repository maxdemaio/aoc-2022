import os.path
import pytest

INPUT_TXT = os.path.join(os.path.dirname(__file__), 'input.txt')


def get_intersection(ids1: list[int], ids2: list[int]) -> list[int]:
    # init 2 pointers
    i = 0
    j = 0
    result: list[int] = []
    while i < len(ids1) and j < len(ids2):
        if ids1[i] < ids2[j]:
            i += 1
        elif ids1[i] > ids2[j]:
            j += 1
        else:
            result.append(ids1[i])
            i += 1
            j += 1
    return result


def compute(s: str) -> int:
    lines = s.splitlines()
    overlap_counter = 0
    for line in lines:
        group: list[str] = line.split(",")
        elf_1_ids = [x for x in range(int(group[0].split("-")[0]), int(group[0].split("-")[1]) + 1)]
        elf_2_ids = [x for x in range(int(group[1].split("-")[0]), int(group[1].split("-")[1]) + 1)]
        intersection: list[int] = get_intersection(elf_1_ids, elf_2_ids)
        if len(intersection) > 0:
            overlap_counter += 1
    # TODO: implement solution here!
    return overlap_counter


# TODO: change for the small example given
INPUT_S = '''\
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8'''
EXPECTED = 4


def test() -> None:
    assert compute(INPUT_S) == EXPECTED


def main() -> int:
    with open(INPUT_TXT) as f:
        print(compute(f.read()))

    return 0


if __name__ == '__main__':
    raise SystemExit(main())
