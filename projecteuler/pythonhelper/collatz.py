def getCollatzTerms(start):
	numTerms = 1;
	value = start
	while(value > 1):
		value = getNextTerm(value)
		numTerms+=1
	return numTerms

def getNextTerm(term):
	if term % 2 == 0:
		return term / 2
	else:
		return (3 * term) + 1

def getStartingValueWithHighestNumTerms(startingValMin, startingValMax):
	highestNumTerms = 0
	startingValueWithHighestNumTerms = 0
	for i in xrange(startingValMin, startingValMax):
		numTerms = getCollatzTerms(i)
		if numTerms > highestNumTerms:
			highestNumTerms = numTerms
			startingValueWithHighestNumTerms = i
	return startingValueWithHighestNumTerms

