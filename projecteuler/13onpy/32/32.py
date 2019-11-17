def is_pandigital(numbers):
	all_digits = ''.join(str(elem) for elem in numbers)

	if len(all_digits) != 9:
		return False

	sorted_digits = ''.join(sorted(all_digits))
	if sorted_digits == '123456789':
		return True
	else:
		return False

products = []
for i in xrange(2, 200):
	for j in xrange(2, 5000):
		if is_pandigital([i, j, i*j]) and not i*j in products:
			print i, j, i*j
			products.append(i*j)

total = 0
for product in products:
	total += product
print total