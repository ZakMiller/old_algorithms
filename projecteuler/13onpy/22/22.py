import io_manager

def get_word_value(word):
	sum = 0
	for char in word:
		sum += ord(char) - 64
	return sum

unsorted_names = io_manager.get_line('names.txt').split(',')
names = sorted(unsorted_names)

total = 0

for i in xrange(0, len(names)):
	total += (i + 1) * get_word_value(names[i][1:-1])
print total
