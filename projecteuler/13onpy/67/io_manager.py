def get_lines():
	with open('input.txt') as input_file:
		lines = input_file.readlines()
	return lines

def get_line():
	with open('input.txt') as input_file:
		line = input_file.readlines()[0]
	return line

def output(string):
	output_filename = 'output.txt'
	if os.path.exists(output_filename):
		os.remove(output_filename)
	with open(output_filename, 'a') as output_file:
		output_file.write(string)