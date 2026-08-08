users = {
    "u1": {"name": "Harsh", "skills": ["Python", "Go"]},
    "u2": {"name": "Rahul", "skills": ["Java", "Python"]},
    "u3": {"name": "Amit", "skills": ["Go", "Rust"]}
}


for key, value in users.items():
    skills = ", ".join(value["skills"])
    print(f'{value["name"]}: {skills}')