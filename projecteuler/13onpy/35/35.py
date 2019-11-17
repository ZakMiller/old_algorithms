import io_manager
def get_cyclic_permutations(n):
	nums = [digit for digit in str(n)]
	split_nums_permutations = [[nums[i - j] for i in range(len(nums))] for j in range(len(nums))]
	complete_nums_permutations = []
	for list_chars in split_nums_permutations:
		complete_nums_permutations.append(int(''.join(list_chars)))
	return complete_nums_permutations

circular_primes = []
primes = io_manager.get_list('primes_to_over_million')
prime_dictionary = {}
for prime in primes:
	prime_dictionary[int(prime)] = True


num_circular_primes = 0
for prime in primes:
	cyclic_permutations = get_cyclic_permutations(prime)
	for cyclic_option in cyclic_permutations:
		all_prime = True
		if not cyclic_option in prime_dictionary:
			all_prime = False
			break
	if all_prime:
		num_circular_primes += 1
		

print num_circular_primes