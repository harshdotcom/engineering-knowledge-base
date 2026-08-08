data = {}

students = [
    ("Harsh", "Indore"),
    ("Rahul", "Delhi"),
    ("Amit", "Indore"),
    ("Priya", "Delhi")
]


for current_student in students:
    data.setdefault(current_student[1], []).append(current_student[0])


print(data)
