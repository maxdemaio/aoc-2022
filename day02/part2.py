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


def get_choice(encrypted: str, elf_answer: str) -> str:
    if encrypted == "X":
        print("we need to lose")
        if elf_answer == "A":
            return "C"
        elif elf_answer == "B":
            return "A"
        else:
            return "B"
    elif encrypted == "Y":
        print("we need to draw")
        return elf_answer
    else:
        print("we need to win")
        if elf_answer == "A":
            return "B"
        elif elf_answer == "B":
            return "C"
        else:
            return "A"


def compute(s: str) -> int:
    lines = s.splitlines()
    score = 0
    for line in lines:
        choices = line.split()

        # TODO: calculate what we'd choose based on 2nd col
        elf_answer = choices[0]
        encrypted_answer = choices[1]
        unencrypted_answer = get_choice(encrypted_answer, elf_answer)

        # get base/unencrypted value of our choice
        base_val = ord(unencrypted_answer) - 64
        score += base_val

        # check if won, lost, or draw

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
EXPECTED = 12


def test() -> None:
    assert compute(INPUT_S) == EXPECTED


def main() -> int:
    with open(INPUT_TXT) as f:
        print(compute(f.read()))

    return 0


if __name__ == '__main__':
    raise SystemExit(main())
