number = [10, 7, 4, 9, 12, 15, 20, 20, 40, 59]



def findEven(nums : list):
    result = []
    for i in nums:
        if i%2 == 0:
            result.append(i)
    return result


print(findEven(number))


