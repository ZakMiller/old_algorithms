def n_is_sum_of_power_of_digits(n, pow):
	total = 0
	for digit in str(n):
		total += int(digit) ** pow
	return total == n


total = 0
for i in xrange(10, 1000000):
	if n_is_sum_of_power_of_digits(i, 5):
		total += i
print total
