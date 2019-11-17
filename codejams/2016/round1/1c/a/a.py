import io_helper

def a(str):
	parties_initial = str.split()
	parties = {}
	leaving_str = ''
	for i in xrange(0, len(parties_initial)):
		parties[chr(i + ord('A'))] = int(parties_initial[i])
	leaving_str = ''
	while parties != {}:
		total = 0
		for key in parties:
			total += parties[key]
		for key in parties:
			if parties[key] - 2 < total / 2










	

testCases = io_helper.get_input()
numCases = testCases[0]
testCases.pop(0)
testCases = testCases[1::2]

io_helper.prep_output()

for i in xrange(0, len(testCases)):
	io_helper.write_output(i, a(testCases[i]))