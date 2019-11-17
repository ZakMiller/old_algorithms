import io_helper

def composeSquare(list):
	listLen = len(list[0].split())
	for str in list:
		subList = str.split()



def b():
	testCases = io_helper.get_input()
	numTestCases = testCases[0]
	testCases.pop(0)

	squads = []
	squadNumber = -1
	for str in testCases:
		if len(str.split()) == 1:
			squadNumber+=1
			squads.append([])
		else:
			squads[squadNumber].append(str)
	return squads




def o():
	io_helper.prep_output()

	for i in xrange(0, len(testCases)):
		io_helper.write_output(i, b(testCases[i]))

print b()