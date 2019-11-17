val = 2**1000
val_string = str(val)
total = 0
for char in val_string:
	total += int(char)
print total