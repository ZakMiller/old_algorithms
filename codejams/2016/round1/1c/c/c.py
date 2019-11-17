import io_helper
def c(list_of_word_pairs):
	first = {}
	second = {}
	for word_pair in list_of_word_pairs:
		word_pair_list = word_pair.split()
		if word_pair_list[0] in first:
			first[word_pair_list[0]].append(word_pair_list[1])
		else:
			first[word_pair_list[0]] = [word_pair_list[1]]

		if word_pair_list[1] in second:
			second[word_pair_list[1]].append(word_pair_list[0])
		else:
			second[word_pair_list[1]] = [word_pair_list[0]]

	for first_word in first:
		if len(first[first_word]) > 1:
			for second_word in first[first_word]:
				if len(second[second_word]) > 1

	
def everything():
	lines = io_helper.get_input()
	lines.pop(0)
	testCases = []
	while len(lines) > 0:
		num_word_pairs = int(lines[0])
		lines.pop(0)
		case = []
		for i in xrange(0, num_word_pairs):
			case.append(lines[i].strip())
		for _ in xrange(0, num_word_pairs):
			lines.pop(0)
		testCases.append(case)


	io_helper.prep_output()

	for i in xrange(0, len(testCases)):
		io_helper.write_output(i, c(testCases[i]))
everything()
