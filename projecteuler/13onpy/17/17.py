ones = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]
teens = ["eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"]
tens = ["ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"]

def get_num_letters(words):
	num_letters = 0
	for char in words:
		if char != ' ' and char != '-':
			num_letters+=1
	return num_letters

def get_ones(digit):
	if int(digit) == 0:
		return ""
	else:
		return ones[int(digit)- 1] 

def get_teens(two_digits):
	return teens[int(two_digits)- 11]

def get_tens(digit):
	if int(digit) == 0:
		return ""
	else:
		return tens[int(digit)- 1]

def get_hundreds(digit):
	return ones[int(digit)- 1] + " hundred"

def get_ones_and_tens(two_digits):
	if int(two_digits) == 0:
		return ""
	elif int(two_digits) > 10 and int(two_digits) < 20:
		return get_teens(two_digits)

	tens = get_tens(str(two_digits)[:1])
	ones = get_ones(str(two_digits)[1:])
	if not ones:
		return tens
	elif not tens:
		return ones
	else:
		return tens + "-" + ones

def get_ones_and_tens_and_hundreds(three_digits):
	hundreds = get_hundreds(str(three_digits)[:1])
	ones_and_tens = get_ones_and_tens(str(three_digits)[1:])
	if not ones_and_tens:
		return hundreds
	else:
		return hundreds + " and " + ones_and_tens
	
def get_words(number):
	if number < 10:
		return get_ones(number)
	elif number == 10:
		return "ten"
	elif number < 20:
		return get_teens(number)
	elif number < 100:
		return get_ones_and_tens(number)
	elif number < 1000:
		return get_ones_and_tens_and_hundreds(number)
	else:
		return "one thousand"

def problem_17():
	num_letters = 0
	for i in xrange(1, 1001):
		words = get_words(i)
		num_letters += get_num_letters(words)

	print num_letters

problem_17()
