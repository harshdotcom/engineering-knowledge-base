"""
Python Syntax Reference — 00_Syntex
===================================

A single runnable file covering Python fundamentals: variables, strings,
numbers, type conversion, conditionals, loops, and functions.

Every section is self-contained. Run the whole file to see all output:

    python main.py

Author: Harsh Jha
"""

import math


# ---------------------------------------------------------------------------
# 1. VARIABLES
# ---------------------------------------------------------------------------
# A variable is a name bound to a value. No type declaration needed —
# Python infers the type at runtime (dynamic typing).
# Naming convention: snake_case, lowercase, descriptive.

student_count = 1000          # int
rating = 4.99                 # float
is_published = False          # bool
course_name = "Python"        # str

print("--- 1. Variables ---")
print(student_count, rating, is_published, course_name)


# ---------------------------------------------------------------------------
# 2. DATA TYPES
# ---------------------------------------------------------------------------
# Primitive (built-in scalar types):
#   int    -> whole numbers        1, -42
#   float  -> decimal numbers      1.1, -0.5
#   bool   -> True / False
#   str    -> text                 "hello"
#   complex-> a + bj               1 + 2j
#
# Non-primitive (containers, covered in later modules):
#   list, tuple, set, dict

whole = 1
decimal = 1.1
complex_number = 1 + 2j       # 'j' is the imaginary unit

print("\n--- 2. Data Types ---")
print(type(whole), type(decimal), type(complex_number))
print(type(is_published), type(course_name))


# ---------------------------------------------------------------------------
# 3. STRINGS
# ---------------------------------------------------------------------------
# Strings are immutable sequences of characters. Any method that "changes"
# a string actually returns a NEW string.

title = "Python Programming"

print("\n--- 3. Strings ---")

# 3.1 Repetition and length
print("*" * 10)               # ********** (repeat operator)
print(len(title))             # 18 — number of characters

# 3.2 Indexing — 0-based from the start, negative from the end
#     P  y  t  h  o  n     P  r  o  g ...
#     0  1  2  3  4  5  6  7  8  9 10 ...
#   -18 ...                            -1
print(title[0])               # 'P'  — first character
print(title[-1])              # 'g'  — last character

# 3.3 Slicing — [start:end:step], 'end' is exclusive
print(title[0:6])             # 'Python'
print(title[7:])              # 'Programming'   (start to end)
print(title[:6])              # 'Python'        (beginning to index 6)
print(title[:])               # full copy
print(title[::-1])            # reversed string

# 3.4 f-strings — the preferred way to build strings (Python 3.6+)
first = "Harsh"
last = "Jha"
full_name = f"{first} {last}"
print(full_name)
print(f"{first} has {len(first)} letters")   # expressions allowed inside {}

# 3.5 Common string methods (all return new values)
print(title.upper())          # 'PYTHON PROGRAMMING'
print(title.lower())          # 'python programming'
print(title.strip())          # removes leading/trailing whitespace
print(title.find("Pro"))      # 7  — index of first match, -1 if not found
print(title.replace("P", "J"))  # case-sensitive replacement
print(title.split(" "))       # ['Python', 'Programming']

# 3.6 Membership — the 'in' operator
print("Pro" in title)         # True
print("Java" not in title)    # True

# 3.7 Escape characters — backslash gives the next character a special meaning
#   \"  double quote     \'  single quote
#   \\  backslash        \n  new line        \t  tab
print("She said \"hello\"\nOn a new line\tafter a tab")

# 3.8 Raw strings — prefix r to disable escaping (useful for Windows paths)
print(r"C:\Users\new_folder")


# ---------------------------------------------------------------------------
# 4. ARITHMETIC OPERATORS
# ---------------------------------------------------------------------------
print("\n--- 4. Arithmetic ---")
print(10 + 3)                 # 13   addition
print(10 - 3)                 # 7    subtraction
print(10 * 3)                 # 30   multiplication
print(10 / 3)                 # 3.333...  true division -> always float
print(10 // 3)                # 3    floor division -> int, drops remainder
print(10 % 3)                 # 1    modulus -> remainder
print(10 ** 3)                # 1000 exponent

# Augmented assignment: x = x + 3 can be written as x += 3
counter = 10
counter += 3                  # also -= *= /= //= %= **=
print(counter)                # 13

# Operator precedence: ()  ->  **  ->  * / // %  ->  + -
print(10 + 3 * 2)             # 16 — multiplication binds tighter
print((10 + 3) * 2)           # 26 — parentheses force order


# ---------------------------------------------------------------------------
# 5. MATH FUNCTIONS
# ---------------------------------------------------------------------------
print("\n--- 5. Math ---")

# Built-in — always available, no import required
print(round(2.9))             # 3   nearest integer
print(round(2.567, 2))        # 2.57  rounded to 2 decimal places
print(abs(-2.9))              # 2.9 absolute value
print(min(3, 7), max(3, 7))   # 3 7

# math module — richer set of functions
print(math.ceil(2.2))         # 3   round up
print(math.floor(2.9))        # 2   round down
print(math.sqrt(16))          # 4.0
print(math.pi)                # 3.14159...


# ---------------------------------------------------------------------------
# 6. TYPE CONVERSION (CASTING)
# ---------------------------------------------------------------------------
# input() ALWAYS returns a string — cast it before doing arithmetic.
#
#   value = input("Enter a number: ")   # "5" (a string)
#   result = int(value) + 1             # 6   (an int)
#
# Conversion functions: int(), float(), bool(), str()

print("\n--- 6. Type Conversion ---")
print(int("5") + 1)           # 6    string -> int
print(float("2.5") + 1)       # 3.5  string -> float
print(str(5) + "1")           # '51' int -> string (concatenation)
print(int(2.9))               # 2    float -> int (truncates, no rounding)


# ---------------------------------------------------------------------------
# 7. TRUTHY AND FALSY
# ---------------------------------------------------------------------------
# Every value has a boolean meaning when used in a condition.
#   Falsy:  ""   0   0.0   None   []   {}   ()   False
#   Truthy: everything else

print("\n--- 7. Truthy / Falsy ---")
print(bool(""), bool(0), bool(None), bool([]))      # all False
print(bool("a"), bool(1), bool([0]))                # all True


# ---------------------------------------------------------------------------
# 8. COMPARISON AND LOGICAL OPERATORS
# ---------------------------------------------------------------------------
# Comparison: ==  !=  >  <  >=  <=      (== compares values)
# Logical:    and (both true)  or (either true)  not (invert)

print("\n--- 8. Operators ---")

high_income = True
good_credit = True
print(high_income and good_credit)   # True
print(high_income or good_credit)    # True
print(not high_income)               # False

# Chained comparison — Python allows the mathematical form
age = 22
print(18 <= age < 65)         # True — cleaner than: age >= 18 and age < 65


# ---------------------------------------------------------------------------
# 9. CONDITIONAL STATEMENTS
# ---------------------------------------------------------------------------
# if / elif / else. Indentation (4 spaces) defines the block — not braces.

print("\n--- 9. Conditionals ---")

temperature = 35

if temperature > 35:
    print("Hot")
elif temperature > 20:
    print("Warm")
else:
    print("Cold")

# Ternary (conditional) expression — a one-line if/else that returns a value
#   <value_if_true> if <condition> else <value_if_false>
message = "Eligible" if age >= 18 else "Not eligible"
print(message)


# ---------------------------------------------------------------------------
# 10. LOOPS
# ---------------------------------------------------------------------------
print("\n--- 10. Loops ---")

# 10.1 range(start, stop, step) — stop is exclusive, returns a lazy sequence
print(list(range(5)))         # [0, 1, 2, 3, 4]
print(list(range(1, 5)))      # [1, 2, 3, 4]
print(list(range(1, 10, 2)))  # [1, 3, 5, 7, 9]

# 10.2 for loop — iterate over any sequence
for number in range(1, 4):
    print("Attempt", number)

for letter in "abc":
    print(letter)

# 10.3 break / continue
#   break    -> exit the loop immediately
#   continue -> skip to the next iteration
for number in range(1, 10):
    if number == 4:
        break                 # stops at 4
    if number % 2 == 0:
        continue              # skips even numbers
    print("Odd:", number)     # 1, 3

# 10.4 for...else — the else block runs ONLY if the loop finished
#      without hitting a break. Useful for search loops.
successful = False
for attempt in range(3):
    print("Trying...", attempt)
    if successful:
        print("Succeeded")
        break
else:
    print("All attempts failed")

# 10.5 Nested loops — the inner loop runs fully for each outer iteration
for x in range(2):
    for y in range(3):
        print(f"({x}, {y})")

# 10.6 while loop — repeats while a condition stays true
number = 8
while number > 0:
    print(number, end=" ")
    number //= 2              # 8 4 2 1
print()

# Interactive while loop (commented out so the file runs unattended):
#
#   command = ""
#   while command.lower() != "quit":
#       command = input("> ")
#       print("ECHO", command)


# ---------------------------------------------------------------------------
# 11. FUNCTIONS
# ---------------------------------------------------------------------------
# A function is a reusable block of code that performs one task.
#   def <name>(<parameters>):
#       <body>
#       return <value>        # optional; defaults to None
#
# Parameters = names in the definition.  Arguments = values passed in.

print("\n--- 11. Functions ---")


def greet(first_name, last_name):
    """Print a greeting. Returns None (no return statement)."""
    print(f"Hi {first_name} {last_name}")
    print("Welcome aboard")


def square(number):
    """Return the square of a number."""
    return number * number


def increment(number, by=1):
    """Default parameter: 'by' is optional and defaults to 1."""
    return number + by


greet("Harsh", "Jha")                 # positional arguments
greet(last_name="Jha", first_name="Harsh")   # keyword arguments — order free

print(square(4))                      # 16
print(increment(5))                   # 6  — uses the default
print(increment(5, by=3))             # 8  — overrides the default


# 11.1 *args — collects extra positional arguments into a tuple
def total(*numbers):
    return sum(numbers)


print(total(1, 2, 3, 4))              # 10


# 11.2 **kwargs — collects extra keyword arguments into a dict
def describe(**details):
    return details


print(describe(name="Harsh", role="Engineer"))


# ---------------------------------------------------------------------------
# 12. SCOPE
# ---------------------------------------------------------------------------
# Variables created inside a function are LOCAL — they do not exist outside.
# Variables at file level are GLOBAL and readable inside functions.

print("\n--- 12. Scope ---")

app_name = "KB"               # global


def show_scope():
    local_only = "inside"     # local — destroyed when the function returns
    print(app_name, local_only)


show_scope()
# print(local_only)           # NameError — not visible out here


# ---------------------------------------------------------------------------
# 13. THE __main__ GUARD
# ---------------------------------------------------------------------------
# When a file is run directly, Python sets __name__ to "__main__".
# When it is imported as a module, __name__ is the module's name.
# This guard keeps demo/entry-point code from running on import.

def main():
    print("\n--- 13. Entry Point ---")
    print("Done")


if __name__ == "__main__":
    main()
