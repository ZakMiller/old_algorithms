import os

def get_lines():
	with open('input.txt') as input_file:
		lines = input_file.readlines()
	return lines

def get_lines(file_name):
	with open(file_name) as input_file:
		lines = input_file.readlines()
	return lines

def get_line():
	with open('input.txt') as input_file:
		line = input_file.readlines()[0]
	return line

def get_line(file_name):
	with open(file_name) as input_file:
		line = input_file.readlines()[0]
	return line

def get_list(file_name):
	entries = []
	lines = get_lines(file_name)
	for line in lines:
		for entry in line.split():
			entries.append(entry)
	return entries

def output(string):
	output(string, 'output.txt')

def output(string, filename):
	output_filename = filename
	if os.path.exists(output_filename):
		os.remove(output_filename)
	with open(output_filename, 'a') as output_file:
		output_file.write(string)
def output_list(list, filename):
	output_filename = filename
	if os.path.exists(output_filename):
		os.remove(output_filename)
	with open(output_filename, 'a') as output_file:
		for string in list:
			output_file.write(str(string) + '\n')