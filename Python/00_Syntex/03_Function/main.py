# def greet_user(name):
#     print(f"Hi {name}")
#     print("Welcome abroad")

# greet_user("Harsh")


# def square(number):
#     return number*number


# result = square(3)
# print(result)
# print(square(3))

try:
    age = int(input('Age: '))
    print(age)
except ZeroDivisionError:
    print('Age cannot be 0')
except ValueError:
    print("Invalid Error")

