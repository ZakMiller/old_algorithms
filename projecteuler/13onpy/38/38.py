import number_analysis

pandigitals = []
for integer in xrange(1, 10000):
	for n in xrange(2, 10000-integer):
		string = ''
		for i in xrange(1, n):
			string = string + str(integer * i)
			if len(string) > 9:
				break
		if number_analysis.is_pandigital(string):
			pandigitals.append(string)
print pandigitals
