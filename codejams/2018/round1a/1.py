import sys


def reduce(instructions):
    to_be_flipped = instructions.rfind("CS")
    return instructions[0:to_be_flipped] + "SC" + instructions[to_be_flipped + 2:]


def problem_one(row_count, col_count, horizontal_cuts, vertical_cuts, rows):
    return str(swaps)


problem_number = 1
infile = sys.stdin
next(infile)
for line in infile:
    problem_pieces = line.split(" ")
    row_count = problem_pieces[0]
    col_count = problem_pieces[1]
    horizontal_cuts = problem_pieces[2]
    vertical_cuts = problem_pieces[3]
    rows = []
    for i in xrange(0, row_count):
        rows.append(news(infile))
    print "Case #" + str(problem_number) + ": " + problem_one(row_count,
                                                              col_count, horizontal_cuts, vertical_cuts, rows)
    problem_number += 1
