import io_manager
def is_prime(n, primes):
	if int(n) in primes:
		return True
	else:
		return False

circular_primes = []
primes = io_manager.get_list('primes_to_over_million')
prime_dictionary = {}
for prime in primes:
	prime_dictionary[int(prime)] = True

truncatable_primes = []
for i in xrange(10, 1000000):
	is_truncatable_prime = True
	for j in xrange(len(str(i))):
		if not is_prime(str(i)[j:], prime_dictionary) or not is_prime(str(i)[0:len(str(i))-j], prime_dictionary):
			is_truncatable_prime = False
	if is_truncatable_prime:
		truncatable_primes.append(i)

print truncatable_primes
print sum(truncatable_primes)

