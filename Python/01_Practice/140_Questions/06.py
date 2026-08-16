kilometers = float(input("Enter Distance in Kilometer: "))

#Converstion factor 1 km = 0.621371 miles
conversion_factor = 0.621371

miles = kilometers*conversion_factor

print(f'{kilometers} Kilometers is equal to {miles} miles')