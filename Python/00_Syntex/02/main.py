# Tuple

number = (1,2,3)
print(number[0])


coordinates = (1,2,3)
x = coordinates[0]
y = coordinates[1]
z = coordinates[2]

x, y , z = coordinates # from line no 6 to 8 exact same as line number 10


customer = {
    "name" : "Harsh Jha",
    "age" : 34,
    "is_verified": True 
}

print(customer["age"])


# phone = input("Phone: ")

# digits_mapping = {
#     "1":"One",
#     "2":"Two",
#     "3":"Three",
#     "4":"Four",
# }


# output = ""
# for ch in phone:
#     output += digits_mapping.get(ch, "!") + " "

# print(output)


message = input(">")
word = message.split(' ')
emojis = {
    ":)":"😊",
    "()": "❤️"
}

output = ""
for wordFind in word:
    output += emojis.get(wordFind, wordFind) + " "

print(output)