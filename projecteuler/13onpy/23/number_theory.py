import io_manager

def sum_proper_divisors(n):
	total = 0
	for i in xrange(1, (n/2)+1):
		if n % i == 0:
			total += i
	return total

def save_all_abundant_numbers_up_to(n):
	abundant_numbers = []
	for i in xrange(12, n):
		if sum_proper_divisors(i) > i:
			abundant_numbers.append(i)
	io_manager.output_list(abundant_numbers, 'abundant_numbers.txt')

save_all_abundant_numbers_up_to(28123)