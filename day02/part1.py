import os.path
import pytest

INPUT_TXT = os.path.join(os.path.dirname(__file__), 'input.txt')


def compare(user: str, computer: str) -> int:
    if user == "A" and computer == "C":
        print('You won! Rock beats scissors.')
        return 6
    elif user == "C" and computer == "B":
        print('You won! Scissors beats paper.')
        return 6
    elif user == "B" and computer == "A":
        print('You won! Paper beats rock')
        return 6
    elif user == "C" and computer == "A":
        print('You lost! Rock beats scissors.')
        return 0
    elif user == "B" and computer == "C":
        print('You lost! Scissors beats paper.')
        return 0
    elif user == "A" and computer == "B":
        print('You lost! Paper beats rock.')
        return 0
    else:
        print('Tie game.')
        return 3


def compute(s: str) -> int:
    lines = s.splitlines()
    score = 0
    for line in lines:
        choices = line.split()
        # get base/unencrypted value of our choice
        base_val = ord(choices[1]) - ord('W')
        score += base_val
        unencrypted_answer = chr(base_val + 64)

        # check if won, lost, or draw
        elf_answer = choices[0]
        score += compare(unencrypted_answer, elf_answer)

        print(base_val)
        print(unencrypted_answer)
        print(choices)
    # TODO: implement solution here!
    return score


# TODO: change for the small example given
INPUT_S = '''\
A Y
B X
C Z'''
EXPECTED = 15


def test() -> None:
    assert compute(INPUT_S) == EXPECTED


def main() -> int:
    with open(INPUT_TXT) as f:
        print(compute(f.read()))

    return 0


if __name__ == '__main__':
    raise SystemExit(main())
