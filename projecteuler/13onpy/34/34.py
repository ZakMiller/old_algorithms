import math

def is_factorial_sum(n):
	total = 0
	for digit in str(n):
		total += math.factorial(int(digit))
	if total == n:
		return True
	else:
		return False
total = 0
for i in xrange(3, 1000000):
	if is_factorial_sum(i):
		total += i
print total