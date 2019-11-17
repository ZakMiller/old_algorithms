def parse(string, count):
	subStrs = []
	lenSubstr = len(string) / count
	for i in xrange(0, count):
		subStrs.append(string[i*lenSubstr:(i+1)*lenSubstr])
	return subStrs