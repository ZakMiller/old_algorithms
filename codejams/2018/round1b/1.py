import sys


def get_combinations_helper(completed_permutations, list_so_far, max_size):
    if max_size == 0:
        completed_permutations.append(tuple(list_so_far))
        return completed_permutations
    else:
        if len(list_so_far) > 0:
            most_recent_value = list_so_far[-1]
        else:
            most_recent_value = 10000
        highest_value_going_forward = min(max_size, most_recent_value)
        for i in xrange(highest_value_going_forward, 0, -1):
            completed_permutations = get_combinations_helper(
                completed_permutations, list_so_far + [i], max_size - i)
    return completed_permutations


def get_combinations(total_people):
    lists = get_combinations_helper([], [], total_people)
    return list(set([tuple(sorted(t)) for t in lists]))


def is_possible(combination, so_far, total_people):
    people_left = total_people
    end_goal = sorted(list(combination))
    end_goal.reverse()
    starting_point = sorted(so_far)
    starting_point.reverse()

    shorter_end = min(len(end_goal), len(so_far))
    for i in xrange(0, shorter_end):
        people_left -= max(int(end_goal[i]) - int(starting_point[i]), 0)

    for i in xrange(shorter_end + 1, len(combination)):
        people_left -= int(combination[i])
    return people_left >= 0


def reduce_combinations(combinations, so_far, total_people):
    possible_combinations = []
    for combination in combinations:
        if is_possible(combination, so_far, total_people):
            possible_combinations.append(combination)

    combo_set = list(set([tuple(sorted(t)) for t in possible_combinations]))
    return combo_set


def get_total(combination, total_people):
    total = 0
    for entry in combination:
        percentage = entry * 100 / total_people
        remainder = (float(entry * 100) / float(total_people)) - percentage
        if remainder >= .5:
            percentage += 1
        total += percentage
    return total


def get_max_total(combinations, total_people):
    max_total = 0
    for combination in combinations:
        max_total = max(get_total(combination, total_people), max_total)
    return max_total


def already_answered(language_choices):
    total = 0
    for language in language_choices:
        total += int(language)
    return total


def problem_one(total_people, language_choices):
    combinations = get_combinations(total_people)
    reduced_combinations = reduce_combinations(
        combinations, language_choices, total_people - already_answered(language_choices))
    max_total = get_max_total(reduced_combinations, total_people)
    return max_total


def problem_1():
    problem_number = 1
    infile = sys.stdin
    next(infile)
    for line in infile:
        pieces_part_one = line.split(" ")
        total_people = int(pieces_part_one[0])
        total_languages = int(pieces_part_one[1])
        next_line = next(infile)
        language_choices = next_line.split(" ")
        print "Case #" + str(problem_number) + ": " + \
            str(problem_one(total_people, language_choices))
        problem_number += 1


#print problem_one(7, ["5"])
problem_1()
