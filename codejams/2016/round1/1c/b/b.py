import io_helper
def permuts(str):
	if not '?' in str:
		return [str]
	options = ['0','1','2','3','4','5','6','7','8','9']
	permutations = []
	for i in xrange(0, len(str)):
		if str[i] == '?':
			if permutations == []:
				for option in options:
					permutations.append(str[0:i] + option)
			else:
				new_permutations = []
				for j in xrange(0, len(permutations)):
					for option in options:
						new_permutations.append(permutations[j] + option)
				permutations = new_permutations
		else:
			for k in xrange(0, len(permutations)):
				permutations[k] += str[i]

	return permutations


def find_closest_match(coders, jammers):
	min_difference = float("inf")
	min_difference_scores = []

	coders.sort()
	jammers.sort()
	for coder_score in coders:
		for jammer_score in jammers:
			if abs(int(coder_score) - int(jammer_score)) < min_difference:
				min_difference = abs(int(coder_score) - int(jammer_score))
				min_difference_scores = [(coder_score, jammer_score)]
			elif abs(int(coder_score) - int(jammer_score)) == min_difference:
				min_difference_scores.append((coder_score, jammer_score))

	return sorted(min_difference_scores, key=lambda tup: (tup[0],tup[1]) )[0]





	
def b(scores):
	scores_list = scores.split()

	coders_permutations = permuts(scores_list[0])
	jammers_permutations = permuts(scores_list[1])

	match = find_closest_match(coders_permutations, jammers_permutations)
	return " ".join(match)
	
def everything():
	testCases = io_helper.get_input()
	numCases = testCases[0]
	testCases.pop(0)

	io_helper.prep_output()

	for i in xrange(0, len(testCases)):
		io_helper.write_output(i, b(testCases[i]))
everything()
