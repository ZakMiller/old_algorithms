from decimal import *
getcontext().prec = 10000

def recurs(sequence, length_subsequence):
	num_instances = len(sequence) / length_subsequence
	for i in xrange(0, num_instances-2):
		if sequence[i*length_subsequence:(i+1)*length_subsequence] != sequence[(i+1)*length_subsequence:(i+2)*length_subsequence]:
			return False
	return True
def get_len_longest_recurring_sequence(n):
	for i in xrange(1, len(n) / 4):
		if recurs(n, i):
			return i
	return -1

max_length = 0
best_d = 0
d = Decimal(2)
while d < 1000:
	n = Decimal(1) / d
	length = get_len_longest_recurring_sequence(str(n)[2:])
	if length > max_length:
		max_length = length
		best_d = d
		string = str(n)[2:]
	d += 1
print best_d, max_length

