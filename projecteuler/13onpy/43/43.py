import permutations

def has_substring_divisibility(digits):
	for i in xrange(1, len(digits)-2):
		n = int(digits[i:i+3])
		if i == 1:
			divis = 2
		elif i == 2:
			divis = 3
		elif i == 3:
			divis = 5
		elif i == 4:
			divis = 7
		elif i == 5:
			divis = 11
		elif i == 6:
			divis = 13
		elif i == 7:
			divis = 17
		if n % divis != 0:
			return False
	return True

pandigital = '9876543210'
total = 0
while pandigital != None:
	if has_substring_divisibility(pandigital):
		total += int(pandigital)
		print pandigital
	pandigital = permutations.get_next_highest_permutation(pandigital)

print total


