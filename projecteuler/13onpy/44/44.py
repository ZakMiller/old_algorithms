pentagonal_numbers = {}
pentagonal_numbers_list = []
for i in xrange(1, 10000):
	pentagonal_number = i * (3 * i - 1) / 2
	pentagonal_numbers[pentagonal_number] = True
	pentagonal_numbers_list.append(pentagonal_number)

diff = 10000000
for i in pentagonal_numbers_list:
	for j in pentagonal_numbers_list:
		if abs(i-j) in pentagonal_numbers and i+j in pentagonal_numbers and abs(i-j) < diff and i != j:
			diff = abs(i-j)
				
print diff


