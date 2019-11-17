import io_manager

abundant_numbers_str = io_manager.get_lines('abundant_numbers.txt')

value_is_sum_of_two_abundant_numbers = [False] * 28123

for i in xrange(0, len(abundant_numbers_str)):
	for j in xrange(0, len(abundant_numbers_str)):
		x = int(abundant_numbers_str[i])
		y = int(abundant_numbers_str[j])
		if x + y < 28123:
			value_is_sum_of_two_abundant_numbers[x + y] = True

values_are_not_sum_of_two_abundant_numbers = []
total = 0
for i in xrange(0, len(value_is_sum_of_two_abundant_numbers)):
	if value_is_sum_of_two_abundant_numbers[i] == False:
		values_are_not_sum_of_two_abundant_numbers.append(i)
		total += i

print values_are_not_sum_of_two_abundant_numbers
print total



