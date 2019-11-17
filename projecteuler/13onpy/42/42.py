import io_manager

def get_int(letter):
	return ord(letter) - 64

def get_value_word(word):
	sum = 0
	for char in word:
		sum += get_int(char)
	return sum

triangle_numbers = {}
for i in xrange(1, 10000):
	triangle_numbers[.5 * i * (i+1)] = True
words = io_manager.get_list_from_project_euler('words.txt')

triangle_numbers_in_list = 0
for word in words:
	if get_value_word(word) in triangle_numbers:
		triangle_numbers_in_list += 1
print triangle_numbers_in_list