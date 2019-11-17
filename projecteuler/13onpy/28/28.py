side_length = 1
sum_diagonals = 1
n = 1
while side_length < 1001:
	side_length += 2
	for _ in xrange(0, 4):
		n += side_length - 1
		sum_diagonals += n
print sum_diagonals

