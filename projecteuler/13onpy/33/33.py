from fractions import Fraction

def is_digit_cancelling(numerator, denominator):
	numerator = str(numerator)
	denominator = str(denominator)
	for digit in str(numerator):
		if digit in str(denominator):
			new_numerator = numerator.replace(digit, "")
			new_denominator = denominator.replace(digit, "")
			if new_denominator == '0' or new_denominator == '' or new_numerator == '':
				continue
			if float(new_numerator) / float(new_denominator) == float(numerator) / float(denominator) and digit != '0':
				return True
	return False

digit_canceling_values_multiplication = 1
for i in xrange(10, 100):
	for j in xrange(10, 100):
		if is_digit_cancelling(i,j) and i < j:
			digit_canceling_values_multiplication *= (float(i)/float(j))

print Fraction(digit_canceling_values_multiplication).limit_denominator()