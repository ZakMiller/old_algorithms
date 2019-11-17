def is_palindrome(string):
	for i in xrange(0, len(string)/2+1):
		if string[i] != string[-(i+1)]:
			return False
	return True



