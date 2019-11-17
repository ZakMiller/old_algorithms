amicable_numbers = [[] for x in xrange(10000)]
sum = 0
def sum_proper_divisors(n):
	total = 0
	for i in xrange(1, (n/2)+1):
		if n % i == 0:
			total += i
	return total
for i in xrange(0, 10000):
	friend = sum_proper_divisors(i)
	if friend < 10000:
		amicable_numbers[i].append(friend)
		if i in amicable_numbers[friend] and i != friend:
			sum+= friend + i

print sum