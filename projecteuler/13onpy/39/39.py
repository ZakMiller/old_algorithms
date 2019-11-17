import time
def is_right_triangle(a, b, c):
	if a**2 + b**2 == c**2:
		return True
	else:
		return False
start = time.time()
max_num_right_triangles = 0
highest_p = 0
for p in xrange(120, 1001):
	num_right_triangles = 0
	for c in xrange(p/3, p/2):
		for b in xrange(1, c):
			a = p - (c+b)
			if is_right_triangle(a,b,c):
				num_right_triangles+=1
	if num_right_triangles > max_num_right_triangles:
		max_num_right_triangles = num_right_triangles
		highest_p = p
end = time.time()
print highest_p, max_num_right_triangles, end - start