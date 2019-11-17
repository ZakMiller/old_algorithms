def canFinish(circle, list):
	for i in xrange(0, len(list)):
		if circle[len(circle)-1][1] == i and list[i] == circle[0][0]:
			return True
	return False

def getHighestCircle(circle, list):
	highest = 0
	circ2 = []
	for i in xrange(0, len(list)):
		if circle[len(circle)-1][1] == i:

			circ2 = circle + [[i, list[i]]]
		if circ2 != []:	
			if canFinish(circ2, list):
				highest = max(highest,len(circ2))
	return highest
def getCircles(list):
	highest = 0
	for i in xrange(0, len(list)):
		highest = max(getHighestCircle([[i, list[i]]], list), highest)
	return highest




print getCircles([2,3,4,1])