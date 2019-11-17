size = 21
matrix = [[0 for x in xrange(0,size)] for x in xrange(0, size)]

matrix[0][0] = 1

for i in xrange(size):
	for j in xrange(size):
		if i > 0:
			matrix[i][j] += matrix[i-1][j]
		if j > 0:
			matrix[i][j] += matrix[i][j-1]

print matrix[size-1][size-1]
