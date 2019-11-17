import io_manager

class prime():
	prime_dictionary = {}
	def __init__(self, number):
		if number == 'mil':
			file = 'mil_primes.txt'
		elif number == 'bil':
			file = 'bil_primes.txt'
		primes = io_manager.get_list(file)
		
		for prime in primes:
			self.prime_dictionary[int(prime)] = True

	def is_prime(self, n):
		if int(n) in self.prime_dictionary:
			return True
		else:
			return False