# 00 — Python Syntax

Foundations of the Python language: variables, data types, strings, operators,
control flow, and functions.

Runnable companion: [main.py](main.py) — every concept below appears there as
working, commented code.

```bash
python main.py
```

---

## Why Python

| Trait | Meaning |
| --- | --- |
| High-level | Memory management and low-level details are handled for you |
| Multi-purpose | Web, automation, data science, ML, scripting, DevOps |
| Cross-platform | The same code runs on Windows, macOS, and Linux |
| Large ecosystem | ~500k+ packages on PyPI |
| Huge community | Answers, tutorials, and libraries for nearly any problem |
| Object-oriented | Everything is an object — including functions and classes |

---

## How Python Runs

```
your_code.py  →  CPython compiler  →  bytecode (.pyc)  →  PVM  →  machine code
```

1. **CPython** — the reference interpreter, written in C. It compiles source to bytecode.
2. **Bytecode** — a platform-independent intermediate form, cached in `__pycache__/`.
3. **PVM** (Python Virtual Machine) — executes bytecode line by line.

This two-step process is what makes Python cross-platform: bytecode is the
same everywhere, only the PVM differs per OS.

---

## Core Concepts

### Variables

A name bound to a value. No type declaration — Python infers it at runtime.

```python
student_count = 1000        # int
rating = 4.99               # float
is_published = False        # bool
course_name = "Python"      # str
```

Convention: `snake_case`, descriptive, lowercase.

### Data Types

**Primitive:** `int`, `float`, `bool`, `str`, `complex`
**Non-primitive:** `list`, `tuple`, `set`, `dict` *(covered in later modules)*

Check any value's type with `type(value)`.

### Strings

Immutable sequences of characters — methods return a **new** string, never
modify the original.

```python
title = "Python Programming"

len(title)          # 18
title[0]            # 'P'    first character
title[-1]           # 'g'    last character
title[0:6]          # 'Python'      [start:end] — end is exclusive
title[::-1]         # reversed
"Pro" in title      # True
```

Useful methods: `.upper()` `.lower()` `.strip()` `.find()` `.replace()` `.split()`

**f-strings** are the preferred way to build strings:

```python
full = f"{first} {last}"
info = f"{first} has {len(first)} letters"   # expressions allowed
```

**Escape characters** — `\` gives the next character a special meaning:

| Escape | Meaning |
| --- | --- |
| `\"` | double quote |
| `\'` | single quote |
| `\\` | backslash |
| `\n` | new line |
| `\t` | tab |

Prefix with `r` for a raw string (no escaping) — handy for Windows paths:
`r"C:\Users\new_folder"`

### Arithmetic Operators

| Operator | Meaning | `10 op 3` |
| --- | --- | --- |
| `+` | addition | `13` |
| `-` | subtraction | `7` |
| `*` | multiplication | `30` |
| `/` | true division → always `float` | `3.333...` |
| `//` | floor division → drops remainder | `3` |
| `%` | modulus → remainder | `1` |
| `**` | exponent | `1000` |

Augmented assignment: `x += 3` is shorthand for `x = x + 3` (also `-=`, `*=`, `/=`, `//=`, `%=`, `**=`).

Precedence: `()` → `**` → `* / // %` → `+ -`

### Math

Built-in (no import): `round()` `abs()` `min()` `max()`
From `math`: `math.ceil()` `math.floor()` `math.sqrt()` `math.pi`

### Type Conversion

`input()` **always returns a string** — cast before doing arithmetic.

```python
value = input("Enter a number: ")   # "5"  (str)
result = int(value) + 1             # 6    (int)
```

Converters: `int()` `float()` `bool()` `str()`
Note: `int(2.9)` → `2` — it truncates, it does not round.

### Truthy and Falsy

Every value has a boolean meaning in a condition.

- **Falsy:** `""` `0` `0.0` `None` `[]` `{}` `()` `False`
- **Truthy:** everything else

### Comparison and Logical Operators

Comparison: `==` `!=` `>` `<` `>=` `<=`
Logical: `and` (both true) · `or` (either true) · `not` (invert)

Python supports chained comparison:

```python
if 18 <= age < 65:              # cleaner
if age >= 18 and age < 65:      # equivalent, more verbose
```

### Conditionals

Indentation (4 spaces) defines the block — Python has no braces.

```python
if temperature > 35:
    print("Hot")
elif temperature > 20:
    print("Warm")
else:
    print("Cold")
```

Ternary expression — a one-line `if/else` that produces a value:

```python
message = "Eligible" if age >= 18 else "Not eligible"
```

### Loops

`range(start, stop, step)` — `stop` is exclusive.

```python
range(5)         # 0 1 2 3 4
range(1, 5)      # 1 2 3 4
range(1, 10, 2)  # 1 3 5 7 9
```

```python
for number in range(1, 4):      # iterate a sequence
    print(number)

while number > 0:               # repeat while a condition holds
    number //= 2
```

- `break` — exit the loop immediately
- `continue` — skip to the next iteration
- `for...else` — the `else` block runs **only** if the loop completed without `break`

### Functions

A reusable block of code that carries out one task.

```python
def increment(number, by=1):    # 'by' is a default parameter
    return number + by

increment(5)                    # 6
increment(5, by=3)              # 8
```

- **Parameters** are names in the definition; **arguments** are the values passed in.
- Positional args go by order; keyword args (`by=3`) go by name and are order-free.
- `*args` collects extra positional args into a **tuple**.
- `**kwargs` collects extra keyword args into a **dict**.
- A function with no `return` returns `None`.

### Scope

Variables created inside a function are **local** and vanish when it returns.
File-level variables are **global** and readable inside functions.

### The `__main__` Guard

```python
if __name__ == "__main__":
    main()
```

Python sets `__name__` to `"__main__"` when a file is run directly, and to the
module name when it is imported. This guard keeps entry-point code from
executing on import.

---

## VS Code Tips

| Shortcut | Action |
| --- | --- |
| `Ctrl + Shift + M` | Open the Problems panel |
| `Ctrl + `` ` `` | Toggle the integrated terminal |
| `Shift + Alt + F` | Format the document |
| `F5` | Run and debug |

---

## Quick Reference

```python
# Output & input
print(value)                      # end=" " to change the line ending
value = input("prompt: ")         # always returns a str

# Inspect
type(x)      len(x)      bool(x)

# Cast
int(x)       float(x)    str(x)

# String
x.upper()    x.lower()   x.strip()
x.find(s)    x.replace(a, b)       x.split(sep)
f"{a} {b}"   "text" * 3            s in x

# Math
round(x, n)  abs(x)      min(a, b)  max(a, b)
math.ceil(x) math.floor(x)          math.sqrt(x)
```

---

## Next

`01_Data_Structures` — lists, tuples, sets, and dictionaries.
