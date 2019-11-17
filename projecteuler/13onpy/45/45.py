t_nums = {}
p_nums = {}
h_nums = {}

for n in xrange(144, 500000):
	t = n * (n + 1) / 2
	t_nums[t] = True
	p = n * (3 * n - 1) / 2
	p_nums[p] = True
	h = n * (2 * n - 1)
	h_nums[h] = True

for key in t_nums:
	if key in p_nums and key in h_nums:
		print key
		break