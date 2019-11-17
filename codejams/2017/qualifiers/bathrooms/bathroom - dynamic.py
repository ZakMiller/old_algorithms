def get_ls(stalls, possible_spot):
    number_of_empty_stalls = 0
    for i in xrange(possible_spot - 1, 0, -1):
        if stalls[i] == False:
            number_of_empty_stalls += 1
        else:
            return number_of_empty_stalls
    return number_of_empty_stalls


def get_rs(stalls, possible_spot):
    number_of_empty_stalls = 0
    for i in xrange(possible_spot + 1, len(stalls)):
        if stalls[i] == False:
            number_of_empty_stalls += 1
        else:
            return number_of_empty_stalls
    return number_of_empty_stalls


def solve(sub_stalls):
    possible_stalls_to_sit_in = []
    for i in xrange(0, len(sub_stalls)):
        if sub_stalls[i] == False:
            Ls = get_ls(sub_stalls, i)
            Rs = get_rs(sub_stalls, i)
            possible_stalls_to_sit_in.append((i, min(Ls, Rs), max(Ls, Rs)))
    possible_stalls_to_sit_in = sorted(
        possible_stalls_to_sit_in, key=lambda element: (-element[1], -element[2]))
    return possible_stalls_to_sit_in[0]


def get_subproblems(stalls):
    filled_stalls = [i for i, x in enumerate(stalls) if x == True]
    problems = []
    for i in xrange(0, len(filled_stalls) - 1):
        first_index = filled_stalls[i]
        second_index = filled_stalls[i + 1]
        if second_index > first_index + 1:
            problems.append(stalls[first_index:second_index + 1])
    return problems

stalls = [True, True, False, True, False, False, False, True]
solutions = {}
for problem in get_subproblems(stalls):
    problem_key = ''.join(problem)
    if problem_key in solutions:
        solution = solutions[problem_key]
    else:
        solution = solve(problem)
        solutions[problem_key] = solution
