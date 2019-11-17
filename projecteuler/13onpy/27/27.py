def is_prime(n):
	for i in xrange(2, abs(n)):
		if n % i == 0:
			return False
	return True

def get_num_consecutive_primes_for_formula(a, b):
	num_primes = 0
	n = 0
	while True:
		if not is_prime(n ** 2 + a * n + b):
			return num_primes
		else:
			num_primes += 1
			n += 1
def problem_27():
	mostPrimes = 0
	a_times_b = 0
	for i in xrange(-1000, 1001):
		for j in xrange(-1000, 1001):
			consecutive_primes = get_num_consecutive_primes_for_formula(i, j)
			if consecutive_primes > mostPrimes:
				mostPrimes = consecutive_primes
				a_times_b = i * j
				print i,j,consecutive_primes
	print a_times_b

problem_27()