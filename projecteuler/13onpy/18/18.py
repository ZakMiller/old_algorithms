import io_manager

lines = io_manager.get_lines()

list = []
for line in lines:
	list.append(map(int,line.split()))

for i in xrange(1, len(list)):
	for j in xrange(0, len(list[i])):
		if j > 0 and j < len(list[i])-1:
			list[i][j] += max( list[i-1][j-1], list[i-1][j] )
		elif j > 0:
			list[i][j] += list[i-1][j-1]
		else:
			list[i][j] += list[i-1][j]

print max(list[len(list)-1])