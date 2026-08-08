user = {
    "name": "Harsh",
    "age": 24,
    "city": "Indore"
}


def print_dictenory(user_dict):
    for key, value in user_dict.items():
        print(f'{key} : {value}')


print_dictenory(user)