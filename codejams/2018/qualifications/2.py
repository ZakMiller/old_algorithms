import sys


def problem_two(problem_pieces):
    odd_indices = problem_pieces[::2]
    even_indices = problem_pieces[1::2]

    odd_indices.sort(key=int)
    even_indices.sort(key=int)

    result = [None]*(len(odd_indices)+len(even_indices))
    result[::2] = odd_indices
    result[1::2] = even_indices

    for i in xrange(0, len(result)-1):
        first = result[i]
        second = result[i + 1]
        if int(first) > int(second):
            return str(i)
    return "OK"


problem_number = 1
infile = sys.stdin
next(infile)
for count in infile:
    line = next(infile)

    problem_pieces = line.strip().split(" ")
    print "Case #" + str(problem_number) + ": " + problem_two(problem_pieces)
    problem_number += 1
