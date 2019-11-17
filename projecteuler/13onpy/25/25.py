before_previous = 1
previous = 1
current = 2
n = 3
while len(str(current)) < 1000:
	before_previous = previous
	previous = current
	current = previous + before_previous
	n += 1

print n

