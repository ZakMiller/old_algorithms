import io_helper

def findSpotsToCheck(numInitialTiles, complexityLevel, numStudents):
	spotsToCheck = []
	totalNumTiles = numInitialTiles ** complexityLevel
	finalTilesPerInitialTile = totalNumTiles / numInitialTiles
	if finalTilesPerInitialTile == 1:
		finalTilesPerInitialTile = 0
	for i in xrange(0, numInitialTiles):
		spotsToCheck.append(str((i + (finalTilesPerInitialTile*i)+1)))
	if len(spotsToCheck) > numStudents:
		return "IMPOSSIBLE"
	else:	
		return ' '.join(spotsToCheck)
def fractiles():
	testCases = io_helper.get_input()
	numTestCases = testCases[0]
	testCases.remove(numTestCases)
	output_filename = 'output.txt'

	io_helper.prep_output()

	for i in xrange(0, len(testCases)):
		testCase = testCases[i].split(' ')
		spotsToCheck = findSpotsToCheck(int(testCase[0]), int(testCase[1]), int(testCase[2]))

		io_helper.write_output(i, spotsToCheck)

fractiles()







