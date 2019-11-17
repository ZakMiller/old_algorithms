def get_distinct_terms_a_to_the_b(a_start, a_end, b_start, b_end):
	terms = {}
	num_terms = 0
	for a in xrange(a_start, a_end+1):
		for b in xrange(b_start, b_end+1):
			val = str(a ** b)
			if not val in terms:
				terms[val] = True
				num_terms += 1
	return num_terms

print get_distinct_terms_a_to_the_b(2, 100, 2, 100)
