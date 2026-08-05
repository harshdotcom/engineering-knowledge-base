import math

# fArray = [5, 2, 5, 2,3]

# for i in fArray:
#     print('X'*i)



# names = ['John', 'Bob', 'Mosh', 'Sarah', 'Marry']
# print(names[2:3])

# numArray = [3,4,1,2,34,5,6]
# maxNumm = numArray[0]

# for num in numArray:
#     if num > maxNumm:
#         maxNumm = num

# print(maxNumm)


matrix = [
    [1,2,3],
    [4,5,6],
    [7,8,9]    
]

# for row in matrix:
#     for ele in row:
#         print(ele)

numbers = [5,3,1,4,5]
numbers.append(20)
print(numbers)
numbers.insert(0, 10)
numbers.remove(5)
# numbers.clear()
numbers.pop() # removing last item
numbers.count(12) # showing the count
print(numbers.index(1))
numbers.sort()

numbers = [5,3,1,4,5]
noDupliate = []

for i in numbers:
    if i not in noDupliate:
        noDupliate.append(i)

print(noDupliate)







