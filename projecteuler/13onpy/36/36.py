import palindrome

total = 0
for i in xrange(1, 1000000):
	if palindrome.is_palindrome(str(i)) and palindrome.is_palindrome(str(bin(i)[2:])):
		total += i

print total