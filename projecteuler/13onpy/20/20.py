import math

factorial_str = str(math.factorial(100))
total = 0
for digit in factorial_str:
	total += int(digit)

print total