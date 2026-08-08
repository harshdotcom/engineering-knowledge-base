users = {
    "user1": {"name": "Harsh", "age": 24},
    "user2": {"name": "Rahul", "age": 25},
    "user3": {"name": "Amit", "age": 23}
}


for key, value in users.items():
    print(value["name"])