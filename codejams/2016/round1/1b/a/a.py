import io_helper
def remove_each_char_from_dict(count, str_chars, dict):
	if count == 0:
		return dict
	for char in str_chars:
		dict[char] -= count
	return dict

def a(str):
	
	chars = {}
	for char in str:
		if char in chars:
			chars[char] += 1
		else:
			chars[char] = 1

	if 'Z' in chars:
		num_zeroes = chars['Z']
	else:
		num_zeroes = 0

	chars = remove_each_char_from_dict(num_zeroes, 'ZERO', chars)

	if 'W' in chars:
		num_twoes = chars['W']
	else:
		num_twoes = 0

	chars = remove_each_char_from_dict(num_twoes, 'TWO', chars)

	if 'X' in chars:
		num_sixes = chars['X']
	else:
		num_sixes = 0

	chars = remove_each_char_from_dict(num_sixes, 'SIX', chars)		

	if 'U' in chars:
		num_fours = chars['U']
	else:
		num_fours = 0

	chars = remove_each_char_from_dict(num_fours, 'FOUR', chars)	

	if 'R' in chars:
		num_threes = chars['R']
	else:
		num_threes = 0

	chars = remove_each_char_from_dict(num_threes, 'THREE', chars)	

	if 'O' in chars:
		num_ones = chars['O']
	else:
		num_ones = 0

	chars = remove_each_char_from_dict(num_ones, 'ONE', chars)	

	if 'F' in chars:
		num_fives = chars['F']
	else:
		num_fives = 0

	chars = remove_each_char_from_dict(num_fives, 'FIVE', chars)	

	if 'V' in chars:
		num_sevens = chars['V']
	else:
		num_sevens = 0

	chars = remove_each_char_from_dict(num_sevens, 'SEVEN', chars)

	if 'G' in chars:
		num_eights = chars['G']
	else:
		num_eights = 0

	chars = remove_each_char_from_dict(num_eights, 'EIGHT', chars)

	if 'I' in chars:
		num_nines = chars['I']
	else:
		num_nines = 0

	number = '0' * num_zeroes + '1' * num_ones + '2' * num_twoes + '3' * num_threes + '4' * num_fours + '5' * num_fives + '6' * num_sixes + '7' * num_sevens + '8' * num_eights + '9' * num_nines
	return number



	

testCases = io_helper.get_input()
numCases = testCases[0]
testCases.pop(0)

io_helper.prep_output()

for i in xrange(0, len(testCases)):
	io_helper.write_output(i, a(testCases[i]))