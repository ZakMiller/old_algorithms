
import prime

def get_index_of_closest_lower(digits, target):
	closest_value = -10
	closest_index = 0
	for i in xrange(0, len(digits)):
		difference = int(target) - int(digits[i])
		if difference > 0 and difference < int(target) - int(closest_value):
			closest_value = digits[i]
			closest_index = i
	return closest_index


prime_finder = prime.prime('bil')
def next_highest_permutation(digits):
	for i in xrange(len(digits)-1, 0, -1):
		if digits[i] < digits[i-1]:
			lower = digits[i]
			higher = digits[i-1]
			remainder = digits[i:]
			index_of_closest_lower = get_index_of_closest_lower(remainder, higher)
			remainder = remainder[:index_of_closest_lower] + remainder[index_of_closest_lower+1:]
			remainder = higher + remainder
			digits = digits[0:i-1] + digits[i+index_of_closest_lower] + ''.join(sorted(remainder, reverse=True))
			return digits

original_digits = '987654321'
digits = original_digits
while not prime_finder.is_prime(digits):
	digits = next_highest_permutation(digits)
	if digits == None:
		digits = original_digits[1:]
		original_digits = digits

print digits




