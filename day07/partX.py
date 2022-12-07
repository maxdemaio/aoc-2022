import os.path
import pytest

INPUT_TXT = os.path.join(os.path.dirname(__file__), 'input.txt')


def compute(s: str) -> int:
    lines = s.splitlines()
    dirs: list[int] = []
    res = 0
    for i in range(0, len(lines)):
        if not lines[i].endswith("(dir)"):
            continue
        print("checking!! " + lines[i])
        # get spacing
        dir_spacing = 0
        for char in lines[i]:
            if char == " ":
                dir_spacing += 1
            else:
                break
        print("spacing!! " + str(dir_spacing))
        dirs.append(0)
        for j in range(i + 1, len(lines)):
            # skip dirs because they're size 0
            if lines[j].endswith("(dir)"):
                continue
            # check spacing
            file_spacing = 0
            for char in lines[j]:
                if char == " ":
                    file_spacing += 1
                else:
                    break
            # skip files not in dir/subdirs
            if dir_spacing >= file_spacing:
                break
            print("\t" + lines[j])
            # grab size and append to list
            found_pos = lines[j].find("=")
            if found_pos != -1:
                dirs[-1] += int(lines[j][found_pos+1:-1])
    for folder in dirs:
        if folder <= 100000:
            res += folder
    return res

# TODO: change for the small example given
INPUT_S = '''\
- / (dir)
  - a (dir)
    - e (dir)
      - i (file, size=584)
    - f (file, size=29116)
    - g (file, size=2557)
    - h.lst (file, size=62596)
  - b.txt (file, size=14848514)
  - c.dat (file, size=8504156)
  - d (dir)
    - j (file, size=4060174)
    - d.log (file, size=8033020)
    - d.ext (file, size=5626152)
    - k (file, size=7214296)'''
EXPECTED = 95437


def test() -> None:
    assert compute(INPUT_S) == EXPECTED


def main() -> int:
    with open(INPUT_TXT) as f:
        print(compute(f.read()))

    return 0


if __name__ == '__main__':
    raise SystemExit(main())
