import os
import prime
import itertools.product


def solve(case):
    list_of_possibility_lists = []
    num_ingredients = case[0]
    num_packages = case[1]
    required_amounts = case[2]
    all_amounts = case[3]

    for i in xrange(0, len(required_amounts)):
        possiblities_list = []
        required = int(required_amounts[i])
        amounts = all_amounts[i]
        for amount in amounts:
            possibilities = []
            for i in xrange(1, 100):
                if amount >= .9 * required * i and amount <= 1.1 * required * i:
                    possibilities.append(i)
            possiblities_list.append(possibilities)
        list_of_possibility_lists.append(
            list(itertools.product(*possibilities)))
    all_combinations = itertools.product(*list_of_possibility_lists)

    max_shared = 0

    # if amount >= .9 * required * i and amount <= 1.1 * required * i:
    #     if i in map:
    #         map[i] += 1
    #     else:
    #         map[i] = 1
    #     break
        list_of_maps.append(map)

    combination
    set(combinations)
    max(possibilities)

    total_orders = 0

    common_keys = set(list_of_maps[0].keys())
    for d in list_of_maps[1:]:
        common_keys.intersection_update(set(d.keys()))
    for key in common_keys:
        seq = [x[key] for x in list_of_maps]
        total_orders += min(seq)

    return total_orders


def do():
    with open('input.txt') as input_file:
        testCases = input_file.readlines()
    numTestCases = testCases[0]
    testCases.remove(numTestCases)
    output_filename = 'output.txt'
    if os.path.exists(output_filename):
        os.remove(output_filename)

    cases = []
    while len(testCases) > 0:
        start = testCases.pop(0)
        num_ingredients, num_packages = start.split(" ")
        required_amounts = testCases.pop(0).split(" ")
        required_amounts = [x.strip() for x in required_amounts]
        all_amounts = []
        for _ in xrange(0, int(num_ingredients)):
            amounts = testCases.pop(0).split(" ")
            amounts = [int(x.strip()) for x in amounts]
            all_amounts.append(amounts)
        case = (num_ingredients.strip(), num_packages.strip(),
                required_amounts, all_amounts)
        cases.append(case)

    solutions = []
    for case in cases:
        solutions.append(solve(case))
    print solutions

    # with open(output_filename, 'a') as output_file:
    #     for i in xrange(0, len(solutions)):
    #         output_file.write('Case #' + str(i + 1) + ': ' + solutions[i])


do()
