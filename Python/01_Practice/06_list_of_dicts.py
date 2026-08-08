users = [
    {"name": "Harsh", "age": 24},
    {"name": "Rahul", "age": 25},
    {"name": "Amit", "age": 23}
]


for user in users:
    if user["age"] > 23:
        print(user["name"])