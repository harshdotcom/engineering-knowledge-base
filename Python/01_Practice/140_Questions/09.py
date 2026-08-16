# Write a program to find the roots of quardatic eqn : ax2  + bx + c = 0
# Roots -b +- (b2-4ac)/2a


import math
a = int(input("Enter coffecient of a: "))
b = int(input("Enter coffecient of b: "))
c = int(input("Enter coffecient of c: "))


# Calculate the Descriminant
discriminent = b*b - 4*a*c

# Check if discriminent is +ive -ive or zero 
if discriminent > 0:
    root1 = (-b + math.sqrt(discriminent))/(2*a)
    root2 = (-b - math.sqrt(discriminent))/(2*a)
    print(f"Root 1: {root1}")
    print(f"Root 1: {root2}")
elif discriminent == 0:
    root = -b/(2*a)
else:
    real_part = -b/(2*a)
    imaginary_part = math.sqrt(abs(discriminent)) / (2*a)
    print(f"Root 1: {real_part} + {imaginary_part}i")
    print(f"Root 2: {real_part} - {imaginary_part}i")
