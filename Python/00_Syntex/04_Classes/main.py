class Point:
    def __init__(self, x, y):
        self.x = x
        self.y = y
    def move(self):
        print("move")

    def draw(self):
        print("draw")


point1 = Point(19,28)
print(point1.x)
# point1.x  = 10
# point1.y = 20



class Person:
    def __init__(self, name):
        self.name = name
    def talk(self):
        print(f"Hi, I am {self.name}")


john = Person("Harsh Jha")
john.talk()


# Inheritence 

