def is_palindrome(string):
	for i in xrange(0, len(string)/2+1):
		if string[i] != string[-(i+1)]:
			return False
	return True

def is_pandigital(numbers):
	all_digits = ''.join(str(elem) for elem in numbers)

	if len(all_digits) != 9:
		return False

	sorted_digits = ''.join(sorted(all_digits))
	if sorted_digits == '123456789':
		return True
	else:
		return False