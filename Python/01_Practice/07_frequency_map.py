numbers = [1, 2, 2, 3, 1, 2, 4, 3]


frequecy_map = {}

for num in numbers:
    frequecy_map[num] = frequecy_map.setdefault(num, 0) + 1


print(frequecy_map)