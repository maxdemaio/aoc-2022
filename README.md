# advent-of-code-template

all my homies like the advent of code

---

## dev setup

- create venv

```
python3 -m venv venv
```

- activate venv

windows

```
venv\Scripts\activate.bat
```

unix/macos

```
source venv/bin/activate
```

- install requirements

```
pip install -r requirements.txt
```

To deactivate the virtual environment, type `deactivate` in the respective terminal window.

## creating a new day's folder

```
bash makeday.sh day01
```

## testing against the small sample

```
pytest day00/part1.py
```
